// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.24.3
// source: grpc/proto/device.proto

package gen

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Represents a Device
type Device struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id               string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"` // Use "_id" for BSON in Go, but just "id" in proto
	UserId           string `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	DeviceType       string `protobuf:"bytes,3,opt,name=device_type,json=deviceType,proto3" json:"device_type,omitempty"`
	Name             string `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	Status           string `protobuf:"bytes,5,opt,name=status,proto3" json:"status,omitempty"`
	SerialNumber     string `protobuf:"bytes,6,opt,name=serial_number,json=serialNumber,proto3" json:"serial_number,omitempty"`
	RegistrationDate int64  `protobuf:"varint,7,opt,name=registration_date,json=registrationDate,proto3" json:"registration_date,omitempty"` // Unix timestamp (seconds since epoch)
	BatteryLevel     int32  `protobuf:"varint,8,opt,name=battery_level,json=batteryLevel,proto3" json:"battery_level,omitempty"`
}

func (x *Device) Reset() {
	*x = Device{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_proto_device_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Device) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Device) ProtoMessage() {}

func (x *Device) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_proto_device_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Device.ProtoReflect.Descriptor instead.
func (*Device) Descriptor() ([]byte, []int) {
	return file_grpc_proto_device_proto_rawDescGZIP(), []int{0}
}

func (x *Device) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Device) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *Device) GetDeviceType() string {
	if x != nil {
		return x.DeviceType
	}
	return ""
}

func (x *Device) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Device) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *Device) GetSerialNumber() string {
	if x != nil {
		return x.SerialNumber
	}
	return ""
}

func (x *Device) GetRegistrationDate() int64 {
	if x != nil {
		return x.RegistrationDate
	}
	return 0
}

func (x *Device) GetBatteryLevel() int32 {
	if x != nil {
		return x.BatteryLevel
	}
	return 0
}

// Request format for creating a device
type CreateDeviceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Device *Device `protobuf:"bytes,1,opt,name=device,proto3" json:"device,omitempty"`
}

func (x *CreateDeviceRequest) Reset() {
	*x = CreateDeviceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_proto_device_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateDeviceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateDeviceRequest) ProtoMessage() {}

func (x *CreateDeviceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_proto_device_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateDeviceRequest.ProtoReflect.Descriptor instead.
func (*CreateDeviceRequest) Descriptor() ([]byte, []int) {
	return file_grpc_proto_device_proto_rawDescGZIP(), []int{1}
}

func (x *CreateDeviceRequest) GetDevice() *Device {
	if x != nil {
		return x.Device
	}
	return nil
}

// Request format for a single device
type DeviceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeviceRequest) Reset() {
	*x = DeviceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_proto_device_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeviceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeviceRequest) ProtoMessage() {}

func (x *DeviceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_proto_device_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeviceRequest.ProtoReflect.Descriptor instead.
func (*DeviceRequest) Descriptor() ([]byte, []int) {
	return file_grpc_proto_device_proto_rawDescGZIP(), []int{2}
}

func (x *DeviceRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

// Response format for device creation and other actions
type DeviceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Success bool   `protobuf:"varint,2,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *DeviceResponse) Reset() {
	*x = DeviceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_proto_device_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeviceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeviceResponse) ProtoMessage() {}

func (x *DeviceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_proto_device_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeviceResponse.ProtoReflect.Descriptor instead.
func (*DeviceResponse) Descriptor() ([]byte, []int) {
	return file_grpc_proto_device_proto_rawDescGZIP(), []int{3}
}

func (x *DeviceResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *DeviceResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

// Response format for a list of devices
type DeviceList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Devices []*Device `protobuf:"bytes,1,rep,name=devices,proto3" json:"devices,omitempty"`
}

func (x *DeviceList) Reset() {
	*x = DeviceList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_proto_device_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeviceList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeviceList) ProtoMessage() {}

func (x *DeviceList) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_proto_device_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeviceList.ProtoReflect.Descriptor instead.
func (*DeviceList) Descriptor() ([]byte, []int) {
	return file_grpc_proto_device_proto_rawDescGZIP(), []int{4}
}

func (x *DeviceList) GetDevices() []*Device {
	if x != nil {
		return x.Devices
	}
	return nil
}

var File_grpc_proto_device_proto protoreflect.FileDescriptor

var file_grpc_proto_device_proto_rawDesc = []byte{
	0x0a, 0x17, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x64, 0x65, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x22, 0xf5, 0x01, 0x0a, 0x06, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x17, 0x0a,
	0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65,
	0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x64, 0x65, 0x76,
	0x69, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x23, 0x0a, 0x0d, 0x73, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x5f, 0x6e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x73, 0x65, 0x72, 0x69,
	0x61, 0x6c, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x2b, 0x0a, 0x11, 0x72, 0x65, 0x67, 0x69,
	0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x10, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x44, 0x61, 0x74, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x62, 0x61, 0x74, 0x74, 0x65, 0x72, 0x79,
	0x5f, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x08, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x62, 0x61,
	0x74, 0x74, 0x65, 0x72, 0x79, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x22, 0x3e, 0x0a, 0x13, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x27, 0x0a, 0x06, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x44, 0x65, 0x76, 0x69,
	0x63, 0x65, 0x52, 0x06, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x22, 0x1f, 0x0a, 0x0d, 0x44, 0x65,
	0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x3a, 0x0a, 0x0e, 0x44,
	0x65, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x18, 0x0a,
	0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07,
	0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x22, 0x37, 0x0a, 0x0a, 0x44, 0x65, 0x76, 0x69, 0x63,
	0x65, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x29, 0x0a, 0x07, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x52, 0x07, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x32, 0xd3, 0x01, 0x0a, 0x0d, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x45, 0x0a, 0x0c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x65, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x1c, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x17, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x44, 0x65, 0x76, 0x69, 0x63,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x38, 0x0a, 0x0d, 0x47, 0x65, 0x74,
	0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x42, 0x79, 0x49, 0x64, 0x12, 0x16, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x0f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x44, 0x65, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x41, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x42, 0x79, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x16, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x13, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x44, 0x65, 0x76, 0x69,
	0x63, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x2f, 0x5a, 0x2d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x42, 0x65, 0x72, 0x72, 0x79, 0x54, 0x72, 0x61, 0x63, 0x65, 0x72,
	0x2f, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f,
	0x67, 0x65, 0x6e, 0x3b, 0x67, 0x65, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_grpc_proto_device_proto_rawDescOnce sync.Once
	file_grpc_proto_device_proto_rawDescData = file_grpc_proto_device_proto_rawDesc
)

func file_grpc_proto_device_proto_rawDescGZIP() []byte {
	file_grpc_proto_device_proto_rawDescOnce.Do(func() {
		file_grpc_proto_device_proto_rawDescData = protoimpl.X.CompressGZIP(file_grpc_proto_device_proto_rawDescData)
	})
	return file_grpc_proto_device_proto_rawDescData
}

var file_grpc_proto_device_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_grpc_proto_device_proto_goTypes = []interface{}{
	(*Device)(nil),              // 0: service.Device
	(*CreateDeviceRequest)(nil), // 1: service.CreateDeviceRequest
	(*DeviceRequest)(nil),       // 2: service.DeviceRequest
	(*DeviceResponse)(nil),      // 3: service.DeviceResponse
	(*DeviceList)(nil),          // 4: service.DeviceList
}
var file_grpc_proto_device_proto_depIdxs = []int32{
	0, // 0: service.CreateDeviceRequest.device:type_name -> service.Device
	0, // 1: service.DeviceList.devices:type_name -> service.Device
	1, // 2: service.DeviceService.CreateDevice:input_type -> service.CreateDeviceRequest
	2, // 3: service.DeviceService.GetDeviceById:input_type -> service.DeviceRequest
	2, // 4: service.DeviceService.GetDevicesByUserId:input_type -> service.DeviceRequest
	3, // 5: service.DeviceService.CreateDevice:output_type -> service.DeviceResponse
	0, // 6: service.DeviceService.GetDeviceById:output_type -> service.Device
	4, // 7: service.DeviceService.GetDevicesByUserId:output_type -> service.DeviceList
	5, // [5:8] is the sub-list for method output_type
	2, // [2:5] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_grpc_proto_device_proto_init() }
func file_grpc_proto_device_proto_init() {
	if File_grpc_proto_device_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_grpc_proto_device_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Device); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_grpc_proto_device_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateDeviceRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_grpc_proto_device_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeviceRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_grpc_proto_device_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeviceResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_grpc_proto_device_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeviceList); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_grpc_proto_device_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_grpc_proto_device_proto_goTypes,
		DependencyIndexes: file_grpc_proto_device_proto_depIdxs,
		MessageInfos:      file_grpc_proto_device_proto_msgTypes,
	}.Build()
	File_grpc_proto_device_proto = out.File
	file_grpc_proto_device_proto_rawDesc = nil
	file_grpc_proto_device_proto_goTypes = nil
	file_grpc_proto_device_proto_depIdxs = nil
}
