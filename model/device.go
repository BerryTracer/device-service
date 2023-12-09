package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Device struct {
	ID               string `bson:"_id,omitempty" json:"id,omitempty"`
	UserID           string `bson:"user_id" json:"user_id"`
	SerialNumber     string `bson:"serial_number" json:"serial_number"`
	DeviceType       string `bson:"device_type" json:"device_type"`
	Name             string `bson:"name" json:"name"`
	Status           string `bson:"status" json:"status"`
	RegistrationDate int64  `bson:"registration_date" json:"registration_date"`
	BatteryLevel     int    `bson:"battery_level" json:"battery_level"`
}

type DeviceDB struct {
	ID               primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	UserID           string             `bson:"user_id" json:"user_id"`
	SerialNumber     string             `bson:"serial_number" json:"serial_number"`
	DeviceType       string             `bson:"device_type" json:"device_type"`
	Name             string             `bson:"name" json:"name"`
	Status           string             `bson:"status" json:"status"`
	RegistrationDate int64              `bson:"registration_date" json:"registration_date"`
	BatteryLevel     int                `bson:"battery_level" json:"battery_level"`
}

func (d *Device) ToDeviceDB() (*DeviceDB, error) {
	objectID, err := primitive.ObjectIDFromHex(d.ID)
	if err != nil {
		return nil, err
	}

	return &DeviceDB{
		ID:               objectID,
		UserID:           d.UserID,
		SerialNumber:     d.SerialNumber,
		DeviceType:       d.DeviceType,
		Name:             d.Name,
		Status:           d.Status,
		RegistrationDate: d.RegistrationDate,
		BatteryLevel:     d.BatteryLevel,
	}, nil
}

func (d *DeviceDB) ToDevice() *Device {
	return &Device{
		ID:               d.ID.Hex(),
		UserID:           d.UserID,
		SerialNumber:     d.SerialNumber,
		DeviceType:       d.DeviceType,
		Name:             d.Name,
		Status:           d.Status,
		RegistrationDate: d.RegistrationDate,
		BatteryLevel:     d.BatteryLevel,
	}
}
