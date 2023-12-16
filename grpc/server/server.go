package server

import (
	"context"
	"errors"
	"log"
	"net"

	authservice "github.com/BerryTracer/auth-service/grpc/proto"
	gen "github.com/BerryTracer/device-service/grpc/proto"
	"github.com/BerryTracer/device-service/model"
	"github.com/BerryTracer/device-service/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type DeviceGrpcServer struct {
	DeviceService service.DeviceService
	AuthService   authservice.AuthServiceClient
	gen.UnimplementedDeviceServiceServer
}

func NewDeviceGrpcServer(deviceService service.DeviceService, authService authservice.AuthServiceClient) *DeviceGrpcServer {
	return &DeviceGrpcServer{
		DeviceService: deviceService,
		AuthService:   authService,
	}
}

func (s *DeviceGrpcServer) Run(port string) error {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("DeviceGrpcServer failed to listen: %v\n", err)
		return err
	}

	server := grpc.NewServer()                 // Consider adding any gRPC options here, if needed
	gen.RegisterDeviceServiceServer(server, s) // Register your Device service with the gRPC server

	log.Printf("DeviceGrpcServer listening on port %s\n", port)
	if err := server.Serve(lis); err != nil {
		log.Fatalf("DeviceGrpcServer failed to serve: %v\n", err)
		return err
	}
	return nil
}

func (s *DeviceGrpcServer) CreateDevice(ctx context.Context, req *gen.CreateDeviceRequest) (*gen.DeviceResponse, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, errors.New("missing metadata from context")
	}

	token := md["authorization"][0]
	tokenResults, err := s.AuthService.VerifyToken(ctx, &authservice.VerifyTokenRequest{
		Token: token,
	})

	if err != nil {
		return nil, err
	}

	if !tokenResults.Valid {
		return nil, errors.New("invalid token")
	}

	device := &model.Device{
		ID:               req.Device.Id,
		UserID:           req.Device.UserId,
		SerialNumber:     req.Device.SerialNumber,
		Name:             req.Device.Name,
		Status:           req.Device.Status,
		DeviceType:       req.Device.DeviceType,
		RegistrationDate: req.Device.RegistrationDate,
		BatteryLevel:     int(req.Device.BatteryLevel),
	}

	err = s.DeviceService.CreateDevice(ctx, device)
	if err != nil {
		return nil, err
	}

	return &gen.DeviceResponse{
		Id:      device.ID,
		Success: true,
	}, nil
}

func (s *DeviceGrpcServer) GetDeviceById(ctx context.Context, req *gen.DeviceRequest) (*gen.Device, error) {
	device, err := s.DeviceService.GetDeviceById(ctx, req.Id)

	if err != nil {
		return nil, err
	}

	return &gen.Device{
		Id:               device.ID,
		UserId:           device.UserID,
		SerialNumber:     device.SerialNumber,
		Name:             device.Name,
		Status:           device.Status,
		DeviceType:       device.DeviceType,
		RegistrationDate: device.RegistrationDate,
		BatteryLevel:     int32(device.BatteryLevel),
	}, nil
}

func (s *DeviceGrpcServer) GetDeviceBySerialNumber(ctx context.Context, req *gen.DeviceRequest) (*gen.Device, error) {
	device, err := s.DeviceService.GetDeviceBySerialNumber(ctx, req.Id)

	if err != nil {
		return nil, err
	}

	return &gen.Device{
		Id:               device.ID,
		UserId:           device.UserID,
		SerialNumber:     device.SerialNumber,
		Name:             device.Name,
		Status:           device.Status,
		DeviceType:       device.DeviceType,
		RegistrationDate: device.RegistrationDate,
		BatteryLevel:     int32(device.BatteryLevel),
	}, nil
}

func (s *DeviceGrpcServer) GetDevicesByUserId(ctx context.Context, req *gen.DeviceRequest) (*gen.DeviceList, error) {
	devices, err := s.DeviceService.GetDevicesByUserId(ctx, req.Id)

	if err != nil {
		return nil, err
	}

	var deviceList []*gen.Device

	for _, device := range devices {
		deviceList = append(deviceList, &gen.Device{
			Id:               device.ID,
			UserId:           device.UserID,
			SerialNumber:     device.SerialNumber,
			Name:             device.Name,
			Status:           device.Status,
			DeviceType:       device.DeviceType,
			RegistrationDate: device.RegistrationDate,
			BatteryLevel:     int32(device.BatteryLevel),
		})
	}

	return &gen.DeviceList{
		Devices: deviceList,
	}, nil
}
