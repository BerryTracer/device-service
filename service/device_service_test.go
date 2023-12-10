package service_test

import (
	"context"
	"testing"

	"github.com/BerryTracer/device-service/model"
	"github.com/BerryTracer/device-service/service"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Mocking the repository
type DeviceRepositoryMock struct {
	mock.Mock
}

func (r *DeviceRepositoryMock) CreateDevice(ctx context.Context, device *model.Device) error {
	args := r.Called(ctx, device)
	return args.Error(0)
}

func (r *DeviceRepositoryMock) GetDeviceById(ctx context.Context, id string) (*model.Device, error) {
	args := r.Called(ctx, id)
	return args.Get(0).(*model.Device), args.Error(1)
}

func (r *DeviceRepositoryMock) GetDeviceBySerialNumber(ctx context.Context, serialNumber string) (*model.Device, error) {
	args := r.Called(ctx, serialNumber)
	return args.Get(0).(*model.Device), args.Error(1)
}

func (r *DeviceRepositoryMock) GetDevicesByUserId(ctx context.Context, userId string) ([]*model.Device, error) {
	args := r.Called(ctx, userId)
	return args.Get(0).([]*model.Device), args.Error(1)
}

func TestDeviceService_CreateDevice(t *testing.T) {
	// Arrange
	device := &model.Device{
		ID:           primitive.NewObjectID().Hex(),
		SerialNumber: "123456789",
		UserID:       "123456789",
	}

	repository := new(DeviceRepositoryMock)
	repository.On("CreateDevice", mock.Anything, device).Return(nil)

	service := service.NewDeviceService(repository)

	// Act
	err := service.CreateDevice(context.Background(), device)

	// Assert
	if err != nil {
		t.Errorf("Error was not expected while creating device: %s", err)
	}
}

func TestDeviceService_GetDeviceById(t *testing.T) {
	// Arrange
	device := &model.Device{
		ID:           primitive.NewObjectID().Hex(),
		SerialNumber: "123456789",
		UserID:       "123456789",
	}

	repository := new(DeviceRepositoryMock)
	repository.On("GetDeviceById", mock.Anything, device.ID).Return(device, nil)

	service := service.NewDeviceService(repository)

	// Act
	result, err := service.GetDeviceById(context.Background(), device.ID)

	// Assert
	if err != nil {
		t.Errorf("Error was not expected while getting device by id: %s", err)
	}

	if result == nil {
		t.Errorf("Result was expected while getting device by id")
	}
}

func TestDeviceService_GetDeviceBySerialNumber(t *testing.T) {
	// Arrange
	device := &model.Device{
		ID:           primitive.NewObjectID().Hex(),
		SerialNumber: "123456789",
		UserID:       "123456789",
	}

	repository := new(DeviceRepositoryMock)
	repository.On("GetDeviceBySerialNumber", mock.Anything, device.SerialNumber).Return(device, nil)

	service := service.NewDeviceService(repository)

	// Act
	result, err := service.GetDeviceBySerialNumber(context.Background(), device.SerialNumber)

	// Assert
	if err != nil {
		t.Errorf("Error was not expected while getting device by serial number: %s", err)
	}

	if result == nil {
		t.Errorf("Result was expected while getting device by serial number")
	}
}

func TestDeviceService_GetDevicesByUserId(t *testing.T) {
	// Arrange
	device := &model.Device{
		ID:           primitive.NewObjectID().Hex(),
		SerialNumber: "123456789",
		UserID:       "123456789",
	}

	repository := new(DeviceRepositoryMock)
	repository.On("GetDevicesByUserId", mock.Anything, device.UserID).Return([]*model.Device{device}, nil)

	service := service.NewDeviceService(repository)

	// Act
	result, err := service.GetDevicesByUserId(context.Background(), device.UserID)

	// Assert
	if err != nil {
		t.Errorf("Error was not expected while getting devices by user id: %s", err)
	}

	if result == nil {
		t.Errorf("Result was expected while getting devices by user id")
	}
}
