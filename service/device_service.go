package service

import (
	"context"

	"github.com/BerryTracer/device-service/model"
	"github.com/BerryTracer/device-service/repository"
)

type DeviceService interface {
	CreateDevice(ctx context.Context, device *model.Device) error
	GetDeviceById(ctx context.Context, id string) (*model.Device, error)
	GetDevicesByUserId(ctx context.Context, userId string) ([]*model.Device, error)
}

type DeviceServiceImpl struct {
	DeviceRepository repository.DeviceRepository
}

// NewDeviceService returns a new DeviceServiceImpl.
func NewDeviceService(deviceRepository repository.DeviceRepository) *DeviceServiceImpl {
	return &DeviceServiceImpl{DeviceRepository: deviceRepository}
}

// CreateDevice implements DeviceService.
func (s *DeviceServiceImpl) CreateDevice(ctx context.Context, device *model.Device) error {
	return s.DeviceRepository.CreateDevice(ctx, device)
}

// GetDeviceById implements DeviceService.
func (s *DeviceServiceImpl) GetDeviceById(ctx context.Context, id string) (*model.Device, error) {
	return s.DeviceRepository.GetDeviceById(ctx, id)
}

// GetDevicesByUserId implements DeviceService.
func (s *DeviceServiceImpl) GetDevicesByUserId(ctx context.Context, userId string) ([]*model.Device, error) {
	return s.DeviceRepository.GetDevicesByUserId(ctx, userId)
}
