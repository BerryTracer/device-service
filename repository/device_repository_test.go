package repository_test

import (
	"context"
	"errors"
	"testing"
	"time"

	mock "github.com/BerryTracer/common-service/adapter/database/mongodb/mock"
	"github.com/BerryTracer/device-service/model"
	"github.com/BerryTracer/device-service/repository"
	"github.com/golang/mock/gomock"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func TestDeviceMongoRepository_CreateDevice(t *testing.T) {
	// Create a new mock controller instance.
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Create a mock adapter instance.
	mockAdapter := mock.NewMockMongoAdapter(ctrl)
	repo := repository.NewDeviceMongoRepository(mockAdapter)

	// Create a new context.
	ctx := context.Background()
	device := &model.Device{
		ID:               primitive.NewObjectID().Hex(),
		UserID:           primitive.NewObjectID().Hex(),
		SerialNumber:     "1234567890",
		Name:             "Test Device",
		DeviceType:       "Test Type",
		Status:           "Test Status",
		RegistrationDate: time.Now().Unix(),
		BatteryLevel:     100,
	}

	// Create a new deviceDB instance.
	deviceDB, err := device.ToDeviceDB()
	if err != nil {
		t.Fatal(err)
	}

	// Create a new mock adapter instance.
	mockAdapter.EXPECT().
		InsertOne(ctx, deviceDB).
		Return(&mongo.InsertOneResult{InsertedID: deviceDB.ID}, nil).
		Times(1)

	// Call the CreateDevice method.
	err = repo.CreateDevice(ctx, device)

	// Check if the error is nil.
	if err != nil {
		t.Fatal(err)
	}
}

func TestDeviceMongoRepository_CreateDevice_Error_ToDeviceDB(t *testing.T) {
	// Create a new mock controller instance.
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Create a mock adapter instance.
	mockAdapter := mock.NewMockMongoAdapter(ctrl)
	repo := repository.NewDeviceMongoRepository(mockAdapter)

	// Create a new context.
	ctx := context.Background()
	device := &model.Device{
		ID:               "Error",
		UserID:           primitive.NewObjectID().Hex(),
		SerialNumber:     "1234567890",
		Name:             "Test Device",
		DeviceType:       "Test Type",
		Status:           "Test Status",
		RegistrationDate: time.Now().Unix(),
		BatteryLevel:     100,
	}

	// Call the CreateDevice method.
	err := repo.CreateDevice(ctx, device)

	// Check if the error is nil.
	if err == nil {
		t.Fatal(err)
	}

}

func TestDeviceMongoRepository_CreateDevice_Error_InsertOne(t *testing.T) {
	// Create a new mock controller instance.
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Create a mock adapter instance.
	mockAdapter := mock.NewMockMongoAdapter(ctrl)
	repo := repository.NewDeviceMongoRepository(mockAdapter)

	// Create a new context.
	ctx := context.Background()
	device := &model.Device{
		ID:               primitive.NewObjectID().Hex(),
		UserID:           primitive.NewObjectID().Hex(),
		SerialNumber:     "1234567890",
		Name:             "Test Device",
		DeviceType:       "Test Type",
		Status:           "Test Status",
		RegistrationDate: time.Now().Unix(),
		BatteryLevel:     100,
	}

	// Create a new deviceDB instance.
	deviceDB, err := device.ToDeviceDB()
	if err != nil {
		t.Fatal(err)
	}

	// Create a new mock adapter instance.
	mockAdapter.EXPECT().
		InsertOne(ctx, deviceDB).
		Return(nil, errors.New("insert failed")).
		Times(1)

	// Call the CreateDevice method.
	err = repo.CreateDevice(ctx, device)

	// Check if the error is nil.
	if err == nil {
		t.Fatal(err)
	}

	// Check if the error is equal to the expected error.
	if err.Error() != "insert failed" {
		t.Fatal(err)
	}

}

func TestDeviceMongoRepository_GetDeviceById(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAdapter := mock.NewMockMongoAdapter(ctrl)
	mockSingleResult := mock.NewMockSingleResult(ctrl)
	repo := repository.NewDeviceMongoRepository(mockAdapter)

	ctx := context.Background()
	testID := primitive.NewObjectID().Hex()
	objectID, err := primitive.ObjectIDFromHex(testID)
	if err != nil {
		t.Fatalf("failed to create objectID from hex: %v", err)
	}

	deviceDB := &model.DeviceDB{ID: objectID}

	// Set up the expectation for FindOne, using the correct type for the ID.
	mockAdapter.EXPECT().
		FindOne(ctx, primitive.M{"_id": testID}). // Use objectID directly
		Return(mockSingleResult).
		Times(1)

	// Set up the expectation for Decode to populate the model.DeviceDB.
	mockSingleResult.EXPECT().
		Decode(gomock.Any()).
		SetArg(0, *deviceDB).
		Return(nil).
		Times(1)

	// Call the GetDeviceById method.
	device, err := repo.GetDeviceById(ctx, testID)

	// Check if the error is nil and the device ID matches the test ID.
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if device.ID != testID {
		t.Errorf("expected device ID to be %v, got %v", testID, device.ID)
	}
}

func TestDeviceMongoRepository_GetDeviceById_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAdapter := mock.NewMockMongoAdapter(ctrl)
	mockSingleResult := mock.NewMockSingleResult(ctrl)
	repo := repository.NewDeviceMongoRepository(mockAdapter)

	ctx := context.Background()
	testID := primitive.NewObjectID().Hex()

	// Set up the expectation for FindOne to return an error.
	mockAdapter.EXPECT().
		FindOne(ctx, primitive.M{"_id": testID}).
		Return(mockSingleResult).
		Times(1)

	// Set up the Decode method to return an error.
	mockSingleResult.EXPECT().
		Decode(gomock.Any()).
		Return(mongo.ErrNoDocuments). // Simulating a 'not found' error
		Times(1)

	// Call the GetDeviceById method.
	_, err := repo.GetDeviceById(ctx, testID)

	// Check if the error is not nil and is the expected error.
	if err != mongo.ErrNoDocuments {
		t.Errorf("expected mongo.ErrNoDocuments, got %v", err)
	}
}

func TestDeviceMongoRepository_GetDevicesByUserId(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAdapter := mock.NewMockMongoAdapter(ctrl)
	mockCursor := mock.NewMockCursor(ctrl)
	repo := repository.NewDeviceMongoRepository(mockAdapter)

	ctx := context.Background()
	userId := "user123"

	// Set up the expectation for Find to return a mock cursor.
	mockAdapter.EXPECT().
		Find(ctx, primitive.M{"user_id": userId}).
		Return(mockCursor, nil).
		Times(1)

	// Prepare mock data to simulate the devices returned by the database query.
	deviceDB1 := &model.DeviceDB{ID: primitive.NewObjectID(), UserID: userId}
	deviceDB2 := &model.DeviceDB{ID: primitive.NewObjectID(), UserID: userId}
	devicesDB := []*model.DeviceDB{deviceDB1, deviceDB2}

	// Set up the expectation for All to populate the slice with the mock data.
	mockCursor.EXPECT().
		All(ctx, gomock.Any()).
		DoAndReturn(func(_ context.Context, v interface{}) error {
			*v.(*[]*model.DeviceDB) = devicesDB
			return nil
		}).
		Times(1)

	// Call the GetDevicesByUserId method.
	devices, err := repo.GetDevicesByUserId(ctx, userId)

	// Check if the error is nil.
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	// Check the number of devices returned.
	if len(devices) != len(devicesDB) {
		t.Errorf("expected %d devices, got %d", len(devicesDB), len(devices))
	}

	// Additional checks can be made here to assert the fields of 'devices' are as expected.
}

func TestDeviceMongoRepository_GetDevicesByUserId_FindError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAdapter := mock.NewMockMongoAdapter(ctrl)
	repo := repository.NewDeviceMongoRepository(mockAdapter)

	ctx := context.Background()
	userId := "user123"

	// Mock the Find method to return an error.
	mockAdapter.EXPECT().
		Find(ctx, primitive.M{"user_id": userId}).
		Return(nil, errors.New("find error")).
		Times(1)

	// Call the GetDevicesByUserId method.
	_, err := repo.GetDevicesByUserId(ctx, userId)

	// Check if an error is returned.
	if err == nil {
		t.Errorf("expected error, got nil")
	}
}

func TestDeviceMongoRepository_GetDevicesByUserId_AllError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAdapter := mock.NewMockMongoAdapter(ctrl)
	mockCursor := mock.NewMockCursor(ctrl)
	repo := repository.NewDeviceMongoRepository(mockAdapter)

	ctx := context.Background()
	userId := "user123"

	// Mock the Find method to return a mock cursor.
	mockAdapter.EXPECT().
		Find(ctx, primitive.M{"user_id": userId}).
		Return(mockCursor, nil).
		Times(1)

	// Mock the All method of the cursor to return an error.
	mockCursor.EXPECT().
		All(ctx, gomock.Any()).
		Return(errors.New("all error")).
		Times(1)

	// Call the GetDevicesByUserId method.
	_, err := repo.GetDevicesByUserId(ctx, userId)

	// Check if an error is returned.
	if err == nil {
		t.Errorf("expected error, got nil")
	}
}

func TestDeviceMongoRepository_GetDeviceBySerialNumber_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAdapter := mock.NewMockMongoAdapter(ctrl)
	mockSingleResult := mock.NewMockSingleResult(ctrl)
	repo := repository.NewDeviceMongoRepository(mockAdapter)

	ctx := context.Background()
	serialNumber := "serial123"

	deviceDB := &model.DeviceDB{SerialNumber: serialNumber}

	// Mock the FindOne method to return a mock SingleResult.
	mockAdapter.EXPECT().
		FindOne(ctx, primitive.M{"serial_number": serialNumber}).
		Return(mockSingleResult).
		Times(1)

	// Mock the Decode method to populate the deviceDB.
	mockSingleResult.EXPECT().
		Decode(gomock.Any()).
		SetArg(0, *deviceDB).
		Return(nil).
		Times(1)

	// Call the GetDeviceBySerialNumber method.
	device, err := repo.GetDeviceBySerialNumber(ctx, serialNumber)

	// Check for successful retrieval and correct data.
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if device.SerialNumber != serialNumber {
		t.Errorf("expected serial number to be %v, got %v", serialNumber, device.SerialNumber)
	}
}

func TestDeviceMongoRepository_GetDeviceBySerialNumber_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAdapter := mock.NewMockMongoAdapter(ctrl)
	mockSingleResult := mock.NewMockSingleResult(ctrl)
	repo := repository.NewDeviceMongoRepository(mockAdapter)

	ctx := context.Background()
	serialNumber := "serial123"

	// Mock the FindOne method to return a mock SingleResult.
	mockAdapter.EXPECT().
		FindOne(ctx, primitive.M{"serial_number": serialNumber}).
		Return(mockSingleResult).
		Times(1)

	// Mock the Decode method to return an error.
	mockSingleResult.EXPECT().
		Decode(gomock.Any()).
		Return(mongo.ErrNoDocuments). // Simulating a 'not found' error
		Times(1)

	// Call the GetDeviceBySerialNumber method.
	_, err := repo.GetDeviceBySerialNumber(ctx, serialNumber)

	// Check for error.
	if err != mongo.ErrNoDocuments {
		t.Errorf("expected mongo.ErrNoDocuments, got %v", err)
	}
}
