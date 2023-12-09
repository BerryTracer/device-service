package repository

import (
	"context"

	"github.com/BerryTracer/common-service/adapter"
	"github.com/BerryTracer/device-service/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type DeviceRepository interface {
	CreateDevice(ctx context.Context, device *model.Device) error
	GetDeviceById(ctx context.Context, id string) (*model.Device, error)
	GetDeviceBySerialNumber(ctx context.Context, serialNumber string) (*model.Device, error)
	GetDevicesByUserId(ctx context.Context, userId string) ([]*model.Device, error)
}

type DeviceMongoRepository struct {
	Collection adapter.MongoAdapter
}

// NewDeviceMongoRepository returns a new DeviceMongoRepository.
func NewDeviceMongoRepository(collection adapter.MongoAdapter) *DeviceMongoRepository {
	return &DeviceMongoRepository{Collection: collection}
}

// CreateDevice implements DeviceRepository.
func (r *DeviceMongoRepository) CreateDevice(ctx context.Context, device *model.Device) error {
	deviceDB, err := device.ToDeviceDB()

	if err != nil {
		return err
	}

	_, err = r.Collection.InsertOne(ctx, deviceDB)

	if err != nil {
		return err
	}

	return nil
}

// GetDeviceById implements DeviceRepository.
func (r *DeviceMongoRepository) GetDeviceById(ctx context.Context, id string) (*model.Device, error) {
	var deviceDB model.DeviceDB
	err := r.Collection.FindOne(ctx, primitive.M{"_id": id}).Decode(&deviceDB)
	if err != nil {
		return nil, err
	}

	return deviceDB.ToDevice(), nil
}

// GetDevicesByUserId implements DeviceRepository.
func (r *DeviceMongoRepository) GetDevicesByUserId(ctx context.Context, userId string) ([]*model.Device, error) {
	var devicesDB []*model.DeviceDB
	cursor, err := r.Collection.Find(ctx, primitive.M{"user_id": userId})
	if err != nil {
		return nil, err
	}
	if err = cursor.All(ctx, &devicesDB); err != nil {
		return nil, err
	}

	var devices []*model.Device
	for _, deviceDB := range devicesDB {
		devices = append(devices, deviceDB.ToDevice())
	}

	return devices, nil
}

// GetDeviceBySerialNumber implements DeviceRepository.
func (r *DeviceMongoRepository) GetDeviceBySerialNumber(ctx context.Context, serialNumber string) (*model.Device, error) {
	var deviceDB model.DeviceDB
	err := r.Collection.FindOne(ctx, primitive.M{"serial_number": serialNumber}).Decode(&deviceDB)
	if err != nil {
		return nil, err
	}

	return deviceDB.ToDevice(), nil
}

// Ensure DeviceMongoRepository implements DeviceRepository interface
var _ DeviceRepository = &DeviceMongoRepository{}
