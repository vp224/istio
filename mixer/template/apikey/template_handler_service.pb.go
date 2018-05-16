// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: mixer/template/apikey/template_handler_service.proto

/*
	Package apikey is a generated protocol buffer package.

	The `apikey` template represents a single API key, which is used for authorization checks.

	The `apikey` template represents a single API key, used to authorize API calls.

	Sample config:

	```yaml
	apiVersion: "config.istio.io/v1alpha2"
	kind: apikey
	metadata:
	  name: validate-apikey
	  namespace: istio-system
	spec:
	  api: api.service | ""
	  api_version: api.version | ""
	  api_operation: api.operation | ""
	  api_key: api.key | ""
	  timestamp: request.time
	```

	It is generated from these files:
		mixer/template/apikey/template_handler_service.proto

	It has these top-level messages:
		HandleApiKeyRequest
		InstanceMsg
		Type
		InstanceParam
*/
package apikey

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"
import _ "istio.io/api/mixer/adapter/model/v1beta1"
import google_protobuf1 "github.com/gogo/protobuf/types"
import istio_mixer_adapter_model_v1beta11 "istio.io/api/mixer/adapter/model/v1beta1"
import istio_mixer_adapter_model_v1beta12 "istio.io/api/mixer/adapter/model/v1beta1"

import context "golang.org/x/net/context"
import grpc "google.golang.org/grpc"

import strings "strings"
import reflect "reflect"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// Request message for HandleApiKey method.
type HandleApiKeyRequest struct {
	// 'apikey' instance.
	Instance *InstanceMsg `protobuf:"bytes,1,opt,name=instance" json:"instance,omitempty"`
	// Adapter specific handler configuration.
	//
	// Note: Backends can also implement [InfrastructureBackend][https://istio.io/docs/reference/config/mixer/istio.mixer.adapter.model.v1beta1.html#InfrastructureBackend]
	// service and therefore opt to receive handler configuration during session creation through [InfrastructureBackend.CreateSession][TODO: Link to this fragment]
	// call. In that case, adapter_config will have type_url as 'google.protobuf.Any.type_url' and would contain string
	// value of session_id (returned from InfrastructureBackend.CreateSession).
	AdapterConfig *google_protobuf1.Any `protobuf:"bytes,2,opt,name=adapter_config,json=adapterConfig" json:"adapter_config,omitempty"`
	// Id to dedupe identical requests from Mixer.
	DedupId string `protobuf:"bytes,3,opt,name=dedup_id,json=dedupId,proto3" json:"dedup_id,omitempty"`
}

func (m *HandleApiKeyRequest) Reset()      { *m = HandleApiKeyRequest{} }
func (*HandleApiKeyRequest) ProtoMessage() {}
func (*HandleApiKeyRequest) Descriptor() ([]byte, []int) {
	return fileDescriptorTemplateHandlerService, []int{0}
}

// Contains instance payload for 'apikey' template. This is passed to infrastructure backends during request-time
// through HandleApiKeyService.HandleApiKey.
type InstanceMsg struct {
	// Name of the instance as specified in configuration.
	Name string `protobuf:"bytes,72295727,opt,name=name,proto3" json:"name,omitempty"`
	// The API being called (api.service).
	Api string `protobuf:"bytes,1,opt,name=api,proto3" json:"api,omitempty"`
	// The version of the API (api.version).
	ApiVersion string `protobuf:"bytes,2,opt,name=api_version,json=apiVersion,proto3" json:"api_version,omitempty"`
	// The API operation is being called.
	ApiOperation string `protobuf:"bytes,3,opt,name=api_operation,json=apiOperation,proto3" json:"api_operation,omitempty"`
	// API key used in API call.
	ApiKey string `protobuf:"bytes,4,opt,name=api_key,json=apiKey,proto3" json:"api_key,omitempty"`
	// Timestamp of API call.
	Timestamp *istio_mixer_adapter_model_v1beta12.TimeStamp `protobuf:"bytes,5,opt,name=timestamp" json:"timestamp,omitempty"`
}

func (m *InstanceMsg) Reset()      { *m = InstanceMsg{} }
func (*InstanceMsg) ProtoMessage() {}
func (*InstanceMsg) Descriptor() ([]byte, []int) {
	return fileDescriptorTemplateHandlerService, []int{1}
}

// Contains inferred type information about specific instance of 'apikey' template. This is passed to
// infrastructure backends during configuration-time through [InfrastructureBackend.CreateSession][TODO: Link to this fragment].
type Type struct {
}

func (m *Type) Reset()                    { *m = Type{} }
func (*Type) ProtoMessage()               {}
func (*Type) Descriptor() ([]byte, []int) { return fileDescriptorTemplateHandlerService, []int{2} }

// Represents instance configuration schema for 'apikey' template.
type InstanceParam struct {
	// The API being called (api.service).
	Api string `protobuf:"bytes,1,opt,name=api,proto3" json:"api,omitempty"`
	// The version of the API (api.version).
	ApiVersion string `protobuf:"bytes,2,opt,name=api_version,json=apiVersion,proto3" json:"api_version,omitempty"`
	// The API operation is being called.
	ApiOperation string `protobuf:"bytes,3,opt,name=api_operation,json=apiOperation,proto3" json:"api_operation,omitempty"`
	// API key used in API call.
	ApiKey string `protobuf:"bytes,4,opt,name=api_key,json=apiKey,proto3" json:"api_key,omitempty"`
	// Timestamp of API call.
	Timestamp string `protobuf:"bytes,5,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (m *InstanceParam) Reset()      { *m = InstanceParam{} }
func (*InstanceParam) ProtoMessage() {}
func (*InstanceParam) Descriptor() ([]byte, []int) {
	return fileDescriptorTemplateHandlerService, []int{3}
}

func init() {
	proto.RegisterType((*HandleApiKeyRequest)(nil), "apikey.HandleApiKeyRequest")
	proto.RegisterType((*InstanceMsg)(nil), "apikey.InstanceMsg")
	proto.RegisterType((*Type)(nil), "apikey.Type")
	proto.RegisterType((*InstanceParam)(nil), "apikey.InstanceParam")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for HandleApiKeyService service

type HandleApiKeyServiceClient interface {
	// HandleApiKey is called by Mixer at request-time to deliver 'apikey' instances to the backend.
	HandleApiKey(ctx context.Context, in *HandleApiKeyRequest, opts ...grpc.CallOption) (*istio_mixer_adapter_model_v1beta11.CheckResult, error)
}

type handleApiKeyServiceClient struct {
	cc *grpc.ClientConn
}

func NewHandleApiKeyServiceClient(cc *grpc.ClientConn) HandleApiKeyServiceClient {
	return &handleApiKeyServiceClient{cc}
}

func (c *handleApiKeyServiceClient) HandleApiKey(ctx context.Context, in *HandleApiKeyRequest, opts ...grpc.CallOption) (*istio_mixer_adapter_model_v1beta11.CheckResult, error) {
	out := new(istio_mixer_adapter_model_v1beta11.CheckResult)
	err := grpc.Invoke(ctx, "/apikey.HandleApiKeyService/HandleApiKey", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for HandleApiKeyService service

type HandleApiKeyServiceServer interface {
	// HandleApiKey is called by Mixer at request-time to deliver 'apikey' instances to the backend.
	HandleApiKey(context.Context, *HandleApiKeyRequest) (*istio_mixer_adapter_model_v1beta11.CheckResult, error)
}

func RegisterHandleApiKeyServiceServer(s *grpc.Server, srv HandleApiKeyServiceServer) {
	s.RegisterService(&_HandleApiKeyService_serviceDesc, srv)
}

func _HandleApiKeyService_HandleApiKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HandleApiKeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HandleApiKeyServiceServer).HandleApiKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/apikey.HandleApiKeyService/HandleApiKey",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HandleApiKeyServiceServer).HandleApiKey(ctx, req.(*HandleApiKeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _HandleApiKeyService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "apikey.HandleApiKeyService",
	HandlerType: (*HandleApiKeyServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "HandleApiKey",
			Handler:    _HandleApiKeyService_HandleApiKey_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "mixer/template/apikey/template_handler_service.proto",
}

func (m *HandleApiKeyRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *HandleApiKeyRequest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Instance != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintTemplateHandlerService(dAtA, i, uint64(m.Instance.Size()))
		n1, err := m.Instance.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if m.AdapterConfig != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintTemplateHandlerService(dAtA, i, uint64(m.AdapterConfig.Size()))
		n2, err := m.AdapterConfig.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	if len(m.DedupId) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintTemplateHandlerService(dAtA, i, uint64(len(m.DedupId)))
		i += copy(dAtA[i:], m.DedupId)
	}
	return i, nil
}

func (m *InstanceMsg) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *InstanceMsg) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Api) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintTemplateHandlerService(dAtA, i, uint64(len(m.Api)))
		i += copy(dAtA[i:], m.Api)
	}
	if len(m.ApiVersion) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintTemplateHandlerService(dAtA, i, uint64(len(m.ApiVersion)))
		i += copy(dAtA[i:], m.ApiVersion)
	}
	if len(m.ApiOperation) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintTemplateHandlerService(dAtA, i, uint64(len(m.ApiOperation)))
		i += copy(dAtA[i:], m.ApiOperation)
	}
	if len(m.ApiKey) > 0 {
		dAtA[i] = 0x22
		i++
		i = encodeVarintTemplateHandlerService(dAtA, i, uint64(len(m.ApiKey)))
		i += copy(dAtA[i:], m.ApiKey)
	}
	if m.Timestamp != nil {
		dAtA[i] = 0x2a
		i++
		i = encodeVarintTemplateHandlerService(dAtA, i, uint64(m.Timestamp.Size()))
		n3, err := m.Timestamp.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n3
	}
	if len(m.Name) > 0 {
		dAtA[i] = 0xfa
		i++
		dAtA[i] = 0xd2
		i++
		dAtA[i] = 0xe4
		i++
		dAtA[i] = 0x93
		i++
		dAtA[i] = 0x2
		i++
		i = encodeVarintTemplateHandlerService(dAtA, i, uint64(len(m.Name)))
		i += copy(dAtA[i:], m.Name)
	}
	return i, nil
}

func (m *Type) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Type) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	return i, nil
}

func (m *InstanceParam) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *InstanceParam) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Api) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintTemplateHandlerService(dAtA, i, uint64(len(m.Api)))
		i += copy(dAtA[i:], m.Api)
	}
	if len(m.ApiVersion) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintTemplateHandlerService(dAtA, i, uint64(len(m.ApiVersion)))
		i += copy(dAtA[i:], m.ApiVersion)
	}
	if len(m.ApiOperation) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintTemplateHandlerService(dAtA, i, uint64(len(m.ApiOperation)))
		i += copy(dAtA[i:], m.ApiOperation)
	}
	if len(m.ApiKey) > 0 {
		dAtA[i] = 0x22
		i++
		i = encodeVarintTemplateHandlerService(dAtA, i, uint64(len(m.ApiKey)))
		i += copy(dAtA[i:], m.ApiKey)
	}
	if len(m.Timestamp) > 0 {
		dAtA[i] = 0x2a
		i++
		i = encodeVarintTemplateHandlerService(dAtA, i, uint64(len(m.Timestamp)))
		i += copy(dAtA[i:], m.Timestamp)
	}
	return i, nil
}

func encodeVarintTemplateHandlerService(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *HandleApiKeyRequest) Size() (n int) {
	var l int
	_ = l
	if m.Instance != nil {
		l = m.Instance.Size()
		n += 1 + l + sovTemplateHandlerService(uint64(l))
	}
	if m.AdapterConfig != nil {
		l = m.AdapterConfig.Size()
		n += 1 + l + sovTemplateHandlerService(uint64(l))
	}
	l = len(m.DedupId)
	if l > 0 {
		n += 1 + l + sovTemplateHandlerService(uint64(l))
	}
	return n
}

func (m *InstanceMsg) Size() (n int) {
	var l int
	_ = l
	l = len(m.Api)
	if l > 0 {
		n += 1 + l + sovTemplateHandlerService(uint64(l))
	}
	l = len(m.ApiVersion)
	if l > 0 {
		n += 1 + l + sovTemplateHandlerService(uint64(l))
	}
	l = len(m.ApiOperation)
	if l > 0 {
		n += 1 + l + sovTemplateHandlerService(uint64(l))
	}
	l = len(m.ApiKey)
	if l > 0 {
		n += 1 + l + sovTemplateHandlerService(uint64(l))
	}
	if m.Timestamp != nil {
		l = m.Timestamp.Size()
		n += 1 + l + sovTemplateHandlerService(uint64(l))
	}
	l = len(m.Name)
	if l > 0 {
		n += 5 + l + sovTemplateHandlerService(uint64(l))
	}
	return n
}

func (m *Type) Size() (n int) {
	var l int
	_ = l
	return n
}

func (m *InstanceParam) Size() (n int) {
	var l int
	_ = l
	l = len(m.Api)
	if l > 0 {
		n += 1 + l + sovTemplateHandlerService(uint64(l))
	}
	l = len(m.ApiVersion)
	if l > 0 {
		n += 1 + l + sovTemplateHandlerService(uint64(l))
	}
	l = len(m.ApiOperation)
	if l > 0 {
		n += 1 + l + sovTemplateHandlerService(uint64(l))
	}
	l = len(m.ApiKey)
	if l > 0 {
		n += 1 + l + sovTemplateHandlerService(uint64(l))
	}
	l = len(m.Timestamp)
	if l > 0 {
		n += 1 + l + sovTemplateHandlerService(uint64(l))
	}
	return n
}

func sovTemplateHandlerService(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozTemplateHandlerService(x uint64) (n int) {
	return sovTemplateHandlerService(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *HandleApiKeyRequest) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&HandleApiKeyRequest{`,
		`Instance:` + strings.Replace(fmt.Sprintf("%v", this.Instance), "InstanceMsg", "InstanceMsg", 1) + `,`,
		`AdapterConfig:` + strings.Replace(fmt.Sprintf("%v", this.AdapterConfig), "Any", "google_protobuf1.Any", 1) + `,`,
		`DedupId:` + fmt.Sprintf("%v", this.DedupId) + `,`,
		`}`,
	}, "")
	return s
}
func (this *InstanceMsg) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&InstanceMsg{`,
		`Api:` + fmt.Sprintf("%v", this.Api) + `,`,
		`ApiVersion:` + fmt.Sprintf("%v", this.ApiVersion) + `,`,
		`ApiOperation:` + fmt.Sprintf("%v", this.ApiOperation) + `,`,
		`ApiKey:` + fmt.Sprintf("%v", this.ApiKey) + `,`,
		`Timestamp:` + strings.Replace(fmt.Sprintf("%v", this.Timestamp), "TimeStamp", "istio_mixer_adapter_model_v1beta12.TimeStamp", 1) + `,`,
		`Name:` + fmt.Sprintf("%v", this.Name) + `,`,
		`}`,
	}, "")
	return s
}
func (this *Type) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&Type{`,
		`}`,
	}, "")
	return s
}
func (this *InstanceParam) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&InstanceParam{`,
		`Api:` + fmt.Sprintf("%v", this.Api) + `,`,
		`ApiVersion:` + fmt.Sprintf("%v", this.ApiVersion) + `,`,
		`ApiOperation:` + fmt.Sprintf("%v", this.ApiOperation) + `,`,
		`ApiKey:` + fmt.Sprintf("%v", this.ApiKey) + `,`,
		`Timestamp:` + fmt.Sprintf("%v", this.Timestamp) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringTemplateHandlerService(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *HandleApiKeyRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTemplateHandlerService
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: HandleApiKeyRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: HandleApiKeyRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Instance", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTemplateHandlerService
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTemplateHandlerService
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Instance == nil {
				m.Instance = &InstanceMsg{}
			}
			if err := m.Instance.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AdapterConfig", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTemplateHandlerService
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTemplateHandlerService
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.AdapterConfig == nil {
				m.AdapterConfig = &google_protobuf1.Any{}
			}
			if err := m.AdapterConfig.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DedupId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTemplateHandlerService
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTemplateHandlerService
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DedupId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTemplateHandlerService(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTemplateHandlerService
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *InstanceMsg) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTemplateHandlerService
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: InstanceMsg: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: InstanceMsg: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Api", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTemplateHandlerService
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTemplateHandlerService
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Api = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ApiVersion", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTemplateHandlerService
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTemplateHandlerService
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ApiVersion = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ApiOperation", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTemplateHandlerService
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTemplateHandlerService
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ApiOperation = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ApiKey", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTemplateHandlerService
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTemplateHandlerService
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ApiKey = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Timestamp", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTemplateHandlerService
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTemplateHandlerService
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Timestamp == nil {
				m.Timestamp = &istio_mixer_adapter_model_v1beta12.TimeStamp{}
			}
			if err := m.Timestamp.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 72295727:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTemplateHandlerService
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTemplateHandlerService
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTemplateHandlerService(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTemplateHandlerService
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Type) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTemplateHandlerService
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Type: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Type: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTemplateHandlerService(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTemplateHandlerService
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *InstanceParam) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTemplateHandlerService
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: InstanceParam: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: InstanceParam: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Api", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTemplateHandlerService
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTemplateHandlerService
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Api = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ApiVersion", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTemplateHandlerService
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTemplateHandlerService
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ApiVersion = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ApiOperation", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTemplateHandlerService
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTemplateHandlerService
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ApiOperation = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ApiKey", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTemplateHandlerService
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTemplateHandlerService
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ApiKey = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Timestamp", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTemplateHandlerService
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTemplateHandlerService
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Timestamp = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTemplateHandlerService(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTemplateHandlerService
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipTemplateHandlerService(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTemplateHandlerService
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTemplateHandlerService
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTemplateHandlerService
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthTemplateHandlerService
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowTemplateHandlerService
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipTemplateHandlerService(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthTemplateHandlerService = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTemplateHandlerService   = fmt.Errorf("proto: integer overflow")
)

func init() {
	proto.RegisterFile("mixer/template/apikey/template_handler_service.proto", fileDescriptorTemplateHandlerService)
}

var fileDescriptorTemplateHandlerService = []byte{
	// 540 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x93, 0x31, 0x6f, 0x13, 0x31,
	0x14, 0xc7, 0xef, 0xda, 0x90, 0x36, 0x4e, 0x83, 0x90, 0x5b, 0xc4, 0x35, 0x20, 0x53, 0x05, 0x09,
	0x3a, 0x54, 0x3e, 0xb5, 0xb0, 0x31, 0xb5, 0x5d, 0x28, 0x15, 0x02, 0x5d, 0x2b, 0x16, 0x86, 0xc8,
	0xc9, 0xbd, 0x5e, 0xad, 0xe4, 0xce, 0xe6, 0xce, 0x89, 0x7a, 0x1b, 0xe2, 0x13, 0x20, 0x75, 0x67,
	0x66, 0xe3, 0x0b, 0xf0, 0x01, 0x2a, 0xa6, 0x8a, 0x89, 0x05, 0x89, 0x3b, 0x3a, 0x30, 0x76, 0x64,
	0x44, 0xf6, 0x39, 0x25, 0x20, 0x54, 0x31, 0xb1, 0xf9, 0xbd, 0xf7, 0x7b, 0x7e, 0xef, 0xff, 0xfc,
	0x8c, 0x1e, 0xc4, 0xfc, 0x08, 0x52, 0x5f, 0x41, 0x2c, 0x87, 0x4c, 0x81, 0xcf, 0x24, 0x1f, 0x40,
	0x7e, 0x61, 0x77, 0x0f, 0x59, 0x12, 0x0e, 0x21, 0xed, 0x66, 0x90, 0x8e, 0x79, 0x1f, 0xa8, 0x4c,
	0x85, 0x12, 0xb8, 0x5e, 0x61, 0xed, 0xa5, 0x48, 0x44, 0xc2, 0xb8, 0x7c, 0x7d, 0xaa, 0xa2, 0xed,
	0xb5, 0xea, 0x4e, 0x16, 0x32, 0xa9, 0x20, 0xf5, 0x63, 0x11, 0xc2, 0xd0, 0x1f, 0xaf, 0xf7, 0x40,
	0xb1, 0x75, 0x1f, 0x8e, 0x14, 0x24, 0x19, 0x17, 0x49, 0x66, 0xe9, 0xe5, 0x48, 0x88, 0x68, 0x08,
	0xbe, 0xb1, 0x7a, 0xa3, 0x03, 0x9f, 0x25, 0xb9, 0x0d, 0xdd, 0xbb, 0xec, 0xa2, 0xfe, 0x21, 0xf4,
	0x07, 0x16, 0xbc, 0x7b, 0x19, 0xa8, 0x72, 0x69, 0xfb, 0xee, 0xbc, 0x75, 0xd1, 0xe2, 0x23, 0xa3,
	0x68, 0x53, 0xf2, 0x5d, 0xc8, 0x03, 0x78, 0x39, 0x82, 0x4c, 0x61, 0x1f, 0xcd, 0xf3, 0x24, 0x53,
	0x2c, 0xe9, 0x83, 0xe7, 0xae, 0xb8, 0xab, 0xcd, 0x8d, 0x45, 0x5a, 0x49, 0xa4, 0x3b, 0xd6, 0xff,
	0x24, 0x8b, 0x82, 0x0b, 0x08, 0x3f, 0x44, 0x57, 0x6d, 0xb1, 0x6e, 0x5f, 0x24, 0x07, 0x3c, 0xf2,
	0x66, 0x4c, 0xda, 0x12, 0xad, 0xd4, 0xd0, 0x89, 0x1a, 0xba, 0x99, 0xe4, 0x41, 0xcb, 0xb2, 0xdb,
	0x06, 0xc5, 0xcb, 0x68, 0x3e, 0x84, 0x70, 0x24, 0xbb, 0x3c, 0xf4, 0x66, 0x57, 0xdc, 0xd5, 0x46,
	0x30, 0x67, 0xec, 0x9d, 0xb0, 0x53, 0xb8, 0xa8, 0x39, 0x55, 0x11, 0x5f, 0x43, 0xb3, 0x4c, 0x72,
	0xd3, 0x53, 0x23, 0xd0, 0x47, 0x7c, 0x1b, 0x35, 0x99, 0xe4, 0xdd, 0x31, 0xa4, 0x7a, 0x88, 0xa6,
	0x6c, 0x23, 0x40, 0x4c, 0xf2, 0xe7, 0x95, 0x07, 0xdf, 0x41, 0x2d, 0x0d, 0x08, 0x09, 0x29, 0x53,
	0x1a, 0xa9, 0x4a, 0x2c, 0x30, 0xc9, 0x9f, 0x4e, 0x7c, 0xf8, 0x06, 0x9a, 0xd3, 0xd0, 0x00, 0x72,
	0xaf, 0x66, 0xc2, 0xfa, 0x45, 0x77, 0x21, 0xc7, 0x8f, 0x51, 0x43, 0xf1, 0x18, 0x32, 0xc5, 0x62,
	0xe9, 0x5d, 0x31, 0x9a, 0xd6, 0x28, 0xcf, 0x14, 0x17, 0xd4, 0xcc, 0x98, 0x5a, 0x29, 0xd4, 0xcc,
	0x98, 0xda, 0x19, 0xd3, 0x7d, 0x1e, 0xc3, 0x9e, 0xce, 0x09, 0x7e, 0xa5, 0xe3, 0xeb, 0xa8, 0x96,
	0xb0, 0x18, 0xbc, 0xf7, 0x1f, 0x3f, 0x74, 0x4c, 0x11, 0x63, 0x76, 0xea, 0xa8, 0xb6, 0x9f, 0x4b,
	0xd0, 0x8f, 0xd1, 0x9a, 0x68, 0x7d, 0xc6, 0x52, 0x16, 0xff, 0x7f, 0xb5, 0xb7, 0xfe, 0x54, 0xdb,
	0x98, 0xea, 0x7f, 0x23, 0xfd, 0x7d, 0x59, 0xf6, 0xaa, 0x2f, 0x80, 0x5f, 0xa0, 0x85, 0x69, 0x37,
	0xbe, 0x39, 0x59, 0x95, 0xbf, 0x6c, 0x56, 0x9b, 0xfe, 0xc3, 0xf0, 0xb6, 0xf5, 0x26, 0x07, 0x90,
	0x8d, 0x86, 0x6a, 0x6b, 0xeb, 0xa4, 0x20, 0xce, 0x69, 0x41, 0x9c, 0xcf, 0x05, 0x71, 0xce, 0x0b,
	0xe2, 0xbc, 0x2a, 0x89, 0xfb, 0xae, 0x24, 0xce, 0x49, 0x49, 0xdc, 0xd3, 0x92, 0xb8, 0x5f, 0x4b,
	0xe2, 0x7e, 0x2f, 0x89, 0x73, 0x5e, 0x12, 0xf7, 0xcd, 0x37, 0xe2, 0xfc, 0xf8, 0x74, 0x76, 0x3c,
	0xe3, 0xbc, 0xfe, 0x72, 0x76, 0x3c, 0x63, 0x7f, 0x65, 0xaf, 0x6e, 0x96, 0xef, 0xfe, 0xcf, 0x00,
	0x00, 0x00, 0xff, 0xff, 0x8e, 0x12, 0x0f, 0x46, 0xdc, 0x03, 0x00, 0x00,
}
