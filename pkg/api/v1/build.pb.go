// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: api/v1/build.proto

package v1

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type BuildRun struct {
	Spec                 *BuildSpec      `protobuf:"bytes,1,opt,name=spec,proto3" json:"spec,omitempty"`
	Config               *BuildRunConfig `protobuf:"bytes,2,opt,name=config,proto3" json:"config,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *BuildRun) Reset()         { *m = BuildRun{} }
func (m *BuildRun) String() string { return proto.CompactTextString(m) }
func (*BuildRun) ProtoMessage()    {}
func (*BuildRun) Descriptor() ([]byte, []int) {
	return fileDescriptor_05c5040d8276e5da, []int{0}
}
func (m *BuildRun) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BuildRun.Unmarshal(m, b)
}
func (m *BuildRun) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BuildRun.Marshal(b, m, deterministic)
}
func (m *BuildRun) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BuildRun.Merge(m, src)
}
func (m *BuildRun) XXX_Size() int {
	return xxx_messageInfo_BuildRun.Size(m)
}
func (m *BuildRun) XXX_DiscardUnknown() {
	xxx_messageInfo_BuildRun.DiscardUnknown(m)
}

var xxx_messageInfo_BuildRun proto.InternalMessageInfo

func (m *BuildRun) GetSpec() *BuildSpec {
	if m != nil {
		return m.Spec
	}
	return nil
}

func (m *BuildRun) GetConfig() *BuildRunConfig {
	if m != nil {
		return m.Config
	}
	return nil
}

// This contains a few extra things at build time
type BuildRunConfig struct {
	BuildEnvVars         *BuildEnvVars      `protobuf:"bytes,1,opt,name=build_env_vars,json=buildEnvVars,proto3" json:"build_env_vars,omitempty"`
	ComputedBuildVars    *ComputedBuildVars `protobuf:"bytes,2,opt,name=computed_build_vars,json=computedBuildVars,proto3" json:"computed_build_vars,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *BuildRunConfig) Reset()         { *m = BuildRunConfig{} }
func (m *BuildRunConfig) String() string { return proto.CompactTextString(m) }
func (*BuildRunConfig) ProtoMessage()    {}
func (*BuildRunConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_05c5040d8276e5da, []int{1}
}
func (m *BuildRunConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BuildRunConfig.Unmarshal(m, b)
}
func (m *BuildRunConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BuildRunConfig.Marshal(b, m, deterministic)
}
func (m *BuildRunConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BuildRunConfig.Merge(m, src)
}
func (m *BuildRunConfig) XXX_Size() int {
	return xxx_messageInfo_BuildRunConfig.Size(m)
}
func (m *BuildRunConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_BuildRunConfig.DiscardUnknown(m)
}

var xxx_messageInfo_BuildRunConfig proto.InternalMessageInfo

func (m *BuildRunConfig) GetBuildEnvVars() *BuildEnvVars {
	if m != nil {
		return m.BuildEnvVars
	}
	return nil
}

func (m *BuildRunConfig) GetComputedBuildVars() *ComputedBuildVars {
	if m != nil {
		return m.ComputedBuildVars
	}
	return nil
}

type BuildEnvVars struct {
	BuildId              string   `protobuf:"bytes,1,opt,name=build_id,json=buildId,proto3" json:"build_id,omitempty"`
	TaggedVersion        string   `protobuf:"bytes,2,opt,name=tagged_version,json=taggedVersion,proto3" json:"tagged_version,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BuildEnvVars) Reset()         { *m = BuildEnvVars{} }
func (m *BuildEnvVars) String() string { return proto.CompactTextString(m) }
func (*BuildEnvVars) ProtoMessage()    {}
func (*BuildEnvVars) Descriptor() ([]byte, []int) {
	return fileDescriptor_05c5040d8276e5da, []int{2}
}
func (m *BuildEnvVars) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BuildEnvVars.Unmarshal(m, b)
}
func (m *BuildEnvVars) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BuildEnvVars.Marshal(b, m, deterministic)
}
func (m *BuildEnvVars) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BuildEnvVars.Merge(m, src)
}
func (m *BuildEnvVars) XXX_Size() int {
	return xxx_messageInfo_BuildEnvVars.Size(m)
}
func (m *BuildEnvVars) XXX_DiscardUnknown() {
	xxx_messageInfo_BuildEnvVars.DiscardUnknown(m)
}

var xxx_messageInfo_BuildEnvVars proto.InternalMessageInfo

func (m *BuildEnvVars) GetBuildId() string {
	if m != nil {
		return m.BuildId
	}
	return ""
}

func (m *BuildEnvVars) GetTaggedVersion() string {
	if m != nil {
		return m.TaggedVersion
	}
	return ""
}

// Holds all values that are computed at build time according to build conditions
type ComputedBuildVars struct {
	// indicates if this build is a release
	Release bool `protobuf:"varint,1,opt,name=release,proto3" json:"release,omitempty"`
	// the image tag to use for all images created during this build session
	ImageTag string `protobuf:"bytes,2,opt,name=image_tag,json=imageTag,proto3" json:"image_tag,omitempty"`
	// the container repo and organization to be used for all images created during this build session
	// ex: gcr.io/solo-public/
	ContainerPrefix string `protobuf:"bytes,3,opt,name=container_prefix,json=containerPrefix,proto3" json:"container_prefix,omitempty"`
	// the version associated with the source code being built
	// during a release, matches the semver tag
	// during a test build, matches the build id
	Version              string   `protobuf:"bytes,4,opt,name=version,proto3" json:"version,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ComputedBuildVars) Reset()         { *m = ComputedBuildVars{} }
func (m *ComputedBuildVars) String() string { return proto.CompactTextString(m) }
func (*ComputedBuildVars) ProtoMessage()    {}
func (*ComputedBuildVars) Descriptor() ([]byte, []int) {
	return fileDescriptor_05c5040d8276e5da, []int{3}
}
func (m *ComputedBuildVars) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ComputedBuildVars.Unmarshal(m, b)
}
func (m *ComputedBuildVars) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ComputedBuildVars.Marshal(b, m, deterministic)
}
func (m *ComputedBuildVars) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ComputedBuildVars.Merge(m, src)
}
func (m *ComputedBuildVars) XXX_Size() int {
	return xxx_messageInfo_ComputedBuildVars.Size(m)
}
func (m *ComputedBuildVars) XXX_DiscardUnknown() {
	xxx_messageInfo_ComputedBuildVars.DiscardUnknown(m)
}

var xxx_messageInfo_ComputedBuildVars proto.InternalMessageInfo

func (m *ComputedBuildVars) GetRelease() bool {
	if m != nil {
		return m.Release
	}
	return false
}

func (m *ComputedBuildVars) GetImageTag() string {
	if m != nil {
		return m.ImageTag
	}
	return ""
}

func (m *ComputedBuildVars) GetContainerPrefix() string {
	if m != nil {
		return m.ContainerPrefix
	}
	return ""
}

func (m *ComputedBuildVars) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

// This gets checked into project repo
type BuildSpec struct {
	Config               *BuildConfig `protobuf:"bytes,1,opt,name=config,proto3" json:"config,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *BuildSpec) Reset()         { *m = BuildSpec{} }
func (m *BuildSpec) String() string { return proto.CompactTextString(m) }
func (*BuildSpec) ProtoMessage()    {}
func (*BuildSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_05c5040d8276e5da, []int{4}
}
func (m *BuildSpec) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BuildSpec.Unmarshal(m, b)
}
func (m *BuildSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BuildSpec.Marshal(b, m, deterministic)
}
func (m *BuildSpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BuildSpec.Merge(m, src)
}
func (m *BuildSpec) XXX_Size() int {
	return xxx_messageInfo_BuildSpec.Size(m)
}
func (m *BuildSpec) XXX_DiscardUnknown() {
	xxx_messageInfo_BuildSpec.DiscardUnknown(m)
}

var xxx_messageInfo_BuildSpec proto.InternalMessageInfo

func (m *BuildSpec) GetConfig() *BuildConfig {
	if m != nil {
		return m.Config
	}
	return nil
}

type BuildConfig struct {
	ReleaseContainerRegistry *ContainerRegistry `protobuf:"bytes,20,opt,name=release_container_registry,json=releaseContainerRegistry,proto3" json:"release_container_registry,omitempty"`
	// optional, if not provided, will use the same registry for release and test
	TestContainerRegistry *ContainerRegistry `protobuf:"bytes,21,opt,name=test_container_registry,json=testContainerRegistry,proto3" json:"test_container_registry,omitempty"`
	XXX_NoUnkeyedLiteral  struct{}           `json:"-"`
	XXX_unrecognized      []byte             `json:"-"`
	XXX_sizecache         int32              `json:"-"`
}

func (m *BuildConfig) Reset()         { *m = BuildConfig{} }
func (m *BuildConfig) String() string { return proto.CompactTextString(m) }
func (*BuildConfig) ProtoMessage()    {}
func (*BuildConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_05c5040d8276e5da, []int{5}
}
func (m *BuildConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BuildConfig.Unmarshal(m, b)
}
func (m *BuildConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BuildConfig.Marshal(b, m, deterministic)
}
func (m *BuildConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BuildConfig.Merge(m, src)
}
func (m *BuildConfig) XXX_Size() int {
	return xxx_messageInfo_BuildConfig.Size(m)
}
func (m *BuildConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_BuildConfig.DiscardUnknown(m)
}

var xxx_messageInfo_BuildConfig proto.InternalMessageInfo

func (m *BuildConfig) GetReleaseContainerRegistry() *ContainerRegistry {
	if m != nil {
		return m.ReleaseContainerRegistry
	}
	return nil
}

func (m *BuildConfig) GetTestContainerRegistry() *ContainerRegistry {
	if m != nil {
		return m.TestContainerRegistry
	}
	return nil
}

type ContainerRegistry struct {
	// Types that are valid to be assigned to Registry:
	//	*ContainerRegistry_Quay
	//	*ContainerRegistry_DockerHub
	//	*ContainerRegistry_Gcr
	Registry             isContainerRegistry_Registry `protobuf_oneof:"registry"`
	XXX_NoUnkeyedLiteral struct{}                     `json:"-"`
	XXX_unrecognized     []byte                       `json:"-"`
	XXX_sizecache        int32                        `json:"-"`
}

func (m *ContainerRegistry) Reset()         { *m = ContainerRegistry{} }
func (m *ContainerRegistry) String() string { return proto.CompactTextString(m) }
func (*ContainerRegistry) ProtoMessage()    {}
func (*ContainerRegistry) Descriptor() ([]byte, []int) {
	return fileDescriptor_05c5040d8276e5da, []int{6}
}
func (m *ContainerRegistry) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ContainerRegistry.Unmarshal(m, b)
}
func (m *ContainerRegistry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ContainerRegistry.Marshal(b, m, deterministic)
}
func (m *ContainerRegistry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ContainerRegistry.Merge(m, src)
}
func (m *ContainerRegistry) XXX_Size() int {
	return xxx_messageInfo_ContainerRegistry.Size(m)
}
func (m *ContainerRegistry) XXX_DiscardUnknown() {
	xxx_messageInfo_ContainerRegistry.DiscardUnknown(m)
}

var xxx_messageInfo_ContainerRegistry proto.InternalMessageInfo

type isContainerRegistry_Registry interface {
	isContainerRegistry_Registry()
}

type ContainerRegistry_Quay struct {
	Quay *QuayRegistry `protobuf:"bytes,10,opt,name=quay,proto3,oneof"`
}
type ContainerRegistry_DockerHub struct {
	DockerHub *DockerHubRegistry `protobuf:"bytes,11,opt,name=docker_hub,json=dockerHub,proto3,oneof"`
}
type ContainerRegistry_Gcr struct {
	Gcr *GoogleContainerRegistry `protobuf:"bytes,12,opt,name=gcr,proto3,oneof"`
}

func (*ContainerRegistry_Quay) isContainerRegistry_Registry()      {}
func (*ContainerRegistry_DockerHub) isContainerRegistry_Registry() {}
func (*ContainerRegistry_Gcr) isContainerRegistry_Registry()       {}

func (m *ContainerRegistry) GetRegistry() isContainerRegistry_Registry {
	if m != nil {
		return m.Registry
	}
	return nil
}

func (m *ContainerRegistry) GetQuay() *QuayRegistry {
	if x, ok := m.GetRegistry().(*ContainerRegistry_Quay); ok {
		return x.Quay
	}
	return nil
}

func (m *ContainerRegistry) GetDockerHub() *DockerHubRegistry {
	if x, ok := m.GetRegistry().(*ContainerRegistry_DockerHub); ok {
		return x.DockerHub
	}
	return nil
}

func (m *ContainerRegistry) GetGcr() *GoogleContainerRegistry {
	if x, ok := m.GetRegistry().(*ContainerRegistry_Gcr); ok {
		return x.Gcr
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*ContainerRegistry) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _ContainerRegistry_OneofMarshaler, _ContainerRegistry_OneofUnmarshaler, _ContainerRegistry_OneofSizer, []interface{}{
		(*ContainerRegistry_Quay)(nil),
		(*ContainerRegistry_DockerHub)(nil),
		(*ContainerRegistry_Gcr)(nil),
	}
}

func _ContainerRegistry_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*ContainerRegistry)
	// registry
	switch x := m.Registry.(type) {
	case *ContainerRegistry_Quay:
		_ = b.EncodeVarint(10<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Quay); err != nil {
			return err
		}
	case *ContainerRegistry_DockerHub:
		_ = b.EncodeVarint(11<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.DockerHub); err != nil {
			return err
		}
	case *ContainerRegistry_Gcr:
		_ = b.EncodeVarint(12<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Gcr); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("ContainerRegistry.Registry has unexpected type %T", x)
	}
	return nil
}

func _ContainerRegistry_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*ContainerRegistry)
	switch tag {
	case 10: // registry.quay
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(QuayRegistry)
		err := b.DecodeMessage(msg)
		m.Registry = &ContainerRegistry_Quay{msg}
		return true, err
	case 11: // registry.docker_hub
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(DockerHubRegistry)
		err := b.DecodeMessage(msg)
		m.Registry = &ContainerRegistry_DockerHub{msg}
		return true, err
	case 12: // registry.gcr
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(GoogleContainerRegistry)
		err := b.DecodeMessage(msg)
		m.Registry = &ContainerRegistry_Gcr{msg}
		return true, err
	default:
		return false, nil
	}
}

func _ContainerRegistry_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*ContainerRegistry)
	// registry
	switch x := m.Registry.(type) {
	case *ContainerRegistry_Quay:
		s := proto.Size(x.Quay)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *ContainerRegistry_DockerHub:
		s := proto.Size(x.DockerHub)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *ContainerRegistry_Gcr:
		s := proto.Size(x.Gcr)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type QuayRegistry struct {
	BaseUrl              string   `protobuf:"bytes,1,opt,name=base_url,json=baseUrl,proto3" json:"base_url,omitempty"`
	Organization         string   `protobuf:"bytes,2,opt,name=organization,proto3" json:"organization,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *QuayRegistry) Reset()         { *m = QuayRegistry{} }
func (m *QuayRegistry) String() string { return proto.CompactTextString(m) }
func (*QuayRegistry) ProtoMessage()    {}
func (*QuayRegistry) Descriptor() ([]byte, []int) {
	return fileDescriptor_05c5040d8276e5da, []int{7}
}
func (m *QuayRegistry) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QuayRegistry.Unmarshal(m, b)
}
func (m *QuayRegistry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QuayRegistry.Marshal(b, m, deterministic)
}
func (m *QuayRegistry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QuayRegistry.Merge(m, src)
}
func (m *QuayRegistry) XXX_Size() int {
	return xxx_messageInfo_QuayRegistry.Size(m)
}
func (m *QuayRegistry) XXX_DiscardUnknown() {
	xxx_messageInfo_QuayRegistry.DiscardUnknown(m)
}

var xxx_messageInfo_QuayRegistry proto.InternalMessageInfo

func (m *QuayRegistry) GetBaseUrl() string {
	if m != nil {
		return m.BaseUrl
	}
	return ""
}

func (m *QuayRegistry) GetOrganization() string {
	if m != nil {
		return m.Organization
	}
	return ""
}

type DockerHubRegistry struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DockerHubRegistry) Reset()         { *m = DockerHubRegistry{} }
func (m *DockerHubRegistry) String() string { return proto.CompactTextString(m) }
func (*DockerHubRegistry) ProtoMessage()    {}
func (*DockerHubRegistry) Descriptor() ([]byte, []int) {
	return fileDescriptor_05c5040d8276e5da, []int{8}
}
func (m *DockerHubRegistry) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DockerHubRegistry.Unmarshal(m, b)
}
func (m *DockerHubRegistry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DockerHubRegistry.Marshal(b, m, deterministic)
}
func (m *DockerHubRegistry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DockerHubRegistry.Merge(m, src)
}
func (m *DockerHubRegistry) XXX_Size() int {
	return xxx_messageInfo_DockerHubRegistry.Size(m)
}
func (m *DockerHubRegistry) XXX_DiscardUnknown() {
	xxx_messageInfo_DockerHubRegistry.DiscardUnknown(m)
}

var xxx_messageInfo_DockerHubRegistry proto.InternalMessageInfo

type GoogleContainerRegistry struct {
	BaseUrl string `protobuf:"bytes,1,opt,name=base_url,json=baseUrl,proto3" json:"base_url,omitempty"`
	// The unique id of your Google Cloud project. Note that project id usually differs from project name.
	// https://cloud.google.com/resource-manager/docs/creating-managing-projects#identifying_projects
	ProjectId            string   `protobuf:"bytes,2,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GoogleContainerRegistry) Reset()         { *m = GoogleContainerRegistry{} }
func (m *GoogleContainerRegistry) String() string { return proto.CompactTextString(m) }
func (*GoogleContainerRegistry) ProtoMessage()    {}
func (*GoogleContainerRegistry) Descriptor() ([]byte, []int) {
	return fileDescriptor_05c5040d8276e5da, []int{9}
}
func (m *GoogleContainerRegistry) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GoogleContainerRegistry.Unmarshal(m, b)
}
func (m *GoogleContainerRegistry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GoogleContainerRegistry.Marshal(b, m, deterministic)
}
func (m *GoogleContainerRegistry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GoogleContainerRegistry.Merge(m, src)
}
func (m *GoogleContainerRegistry) XXX_Size() int {
	return xxx_messageInfo_GoogleContainerRegistry.Size(m)
}
func (m *GoogleContainerRegistry) XXX_DiscardUnknown() {
	xxx_messageInfo_GoogleContainerRegistry.DiscardUnknown(m)
}

var xxx_messageInfo_GoogleContainerRegistry proto.InternalMessageInfo

func (m *GoogleContainerRegistry) GetBaseUrl() string {
	if m != nil {
		return m.BaseUrl
	}
	return ""
}

func (m *GoogleContainerRegistry) GetProjectId() string {
	if m != nil {
		return m.ProjectId
	}
	return ""
}

func init() {
	proto.RegisterType((*BuildRun)(nil), "build.solo.io.BuildRun")
	proto.RegisterType((*BuildRunConfig)(nil), "build.solo.io.BuildRunConfig")
	proto.RegisterType((*BuildEnvVars)(nil), "build.solo.io.BuildEnvVars")
	proto.RegisterType((*ComputedBuildVars)(nil), "build.solo.io.ComputedBuildVars")
	proto.RegisterType((*BuildSpec)(nil), "build.solo.io.BuildSpec")
	proto.RegisterType((*BuildConfig)(nil), "build.solo.io.BuildConfig")
	proto.RegisterType((*ContainerRegistry)(nil), "build.solo.io.ContainerRegistry")
	proto.RegisterType((*QuayRegistry)(nil), "build.solo.io.QuayRegistry")
	proto.RegisterType((*DockerHubRegistry)(nil), "build.solo.io.DockerHubRegistry")
	proto.RegisterType((*GoogleContainerRegistry)(nil), "build.solo.io.GoogleContainerRegistry")
}

func init() { proto.RegisterFile("api/v1/build.proto", fileDescriptor_05c5040d8276e5da) }

var fileDescriptor_05c5040d8276e5da = []byte{
	// 541 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x54, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0x6d, 0xda, 0xa8, 0x8d, 0x27, 0x69, 0x20, 0x5b, 0xaa, 0x9a, 0x56, 0x95, 0x2a, 0x4b, 0x20,
	0x90, 0x50, 0xaa, 0x14, 0x71, 0xe1, 0x82, 0x9a, 0x82, 0x48, 0x0f, 0x48, 0x61, 0x0b, 0x15, 0xe2,
	0x80, 0xb5, 0xb1, 0xb7, 0xcb, 0x82, 0xeb, 0x35, 0xeb, 0xb5, 0x45, 0xf8, 0x09, 0xbe, 0x82, 0x6f,
	0xe1, 0xcc, 0x1f, 0x21, 0x8f, 0x37, 0x69, 0x12, 0xbb, 0x88, 0xe3, 0xbc, 0x7d, 0xf3, 0xde, 0xd3,
	0xcc, 0xd8, 0x40, 0x58, 0x22, 0x8f, 0xf3, 0xc1, 0xf1, 0x24, 0x93, 0x51, 0xd8, 0x4f, 0xb4, 0x32,
	0x8a, 0x6c, 0x97, 0x45, 0xaa, 0x22, 0xd5, 0x97, 0xca, 0x53, 0xd0, 0x1a, 0x16, 0x00, 0xcd, 0x62,
	0xf2, 0x04, 0x9a, 0x69, 0xc2, 0x03, 0xb7, 0x71, 0xd4, 0x78, 0xd4, 0x3e, 0x71, 0xfb, 0x4b, 0xcc,
	0x3e, 0xd2, 0x2e, 0x12, 0x1e, 0x50, 0x64, 0x91, 0x67, 0xb0, 0x19, 0xa8, 0xf8, 0x4a, 0x0a, 0x77,
	0x1d, 0xf9, 0x87, 0x75, 0x7c, 0x9a, 0xc5, 0x67, 0x48, 0xa2, 0x96, 0xec, 0xfd, 0x6a, 0x40, 0x77,
	0xf9, 0x89, 0x9c, 0x42, 0x17, 0x5b, 0x7d, 0x1e, 0xe7, 0x7e, 0xce, 0x74, 0x6a, 0x13, 0x1c, 0xd4,
	0x29, 0xbe, 0x8a, 0xf3, 0x4b, 0xa6, 0x53, 0xda, 0x99, 0x2c, 0x54, 0x64, 0x0c, 0x3b, 0x81, 0xba,
	0x4e, 0x32, 0xc3, 0x43, 0xbf, 0xd4, 0x42, 0x9d, 0x32, 0xd9, 0xd1, 0x8a, 0xce, 0x99, 0x65, 0xa2,
	0x1e, 0x8a, 0xf5, 0x82, 0x55, 0xc8, 0x1b, 0x43, 0x67, 0xd1, 0x8f, 0xdc, 0x87, 0x56, 0x29, 0x2c,
	0x43, 0x8c, 0xe7, 0xd0, 0x2d, 0xac, 0xcf, 0x43, 0xf2, 0x00, 0xba, 0x86, 0x09, 0xc1, 0x43, 0x3f,
	0xe7, 0x3a, 0x95, 0x2a, 0x46, 0x5f, 0x87, 0x6e, 0x97, 0xe8, 0x65, 0x09, 0x7a, 0x3f, 0x1b, 0xd0,
	0xab, 0x58, 0x13, 0x17, 0xb6, 0x34, 0x8f, 0x38, 0x4b, 0x39, 0xca, 0xb6, 0xe8, 0xac, 0x24, 0x07,
	0xe0, 0xc8, 0x6b, 0x26, 0xb8, 0x6f, 0x98, 0xb0, 0x8a, 0x2d, 0x04, 0xde, 0x31, 0x41, 0x1e, 0xc3,
	0xdd, 0x40, 0xc5, 0x86, 0xc9, 0x98, 0x6b, 0x3f, 0xd1, 0xfc, 0x4a, 0x7e, 0x77, 0x37, 0x90, 0x73,
	0x67, 0x8e, 0x8f, 0x11, 0x2e, 0x1c, 0x66, 0xb9, 0x9a, 0x65, 0x70, 0x5b, 0x7a, 0x2f, 0xc0, 0x99,
	0x6f, 0x95, 0x9c, 0xcc, 0xf7, 0x59, 0x4e, 0x7f, 0xbf, 0x6e, 0xfa, 0x2b, 0xcb, 0xfc, 0xdd, 0x80,
	0xf6, 0x02, 0x4e, 0x3e, 0xc1, 0xbe, 0x4d, 0xef, 0xdf, 0xa4, 0xd3, 0x5c, 0xc8, 0xd4, 0xe8, 0xa9,
	0x7b, 0xef, 0x96, 0x6d, 0x58, 0x22, 0xb5, 0x3c, 0xea, 0x5a, 0x8d, 0xca, 0x0b, 0xf9, 0x00, 0x7b,
	0x86, 0xa7, 0xa6, 0x4e, 0x7c, 0xf7, 0x3f, 0xc5, 0x77, 0x0b, 0x81, 0x0a, 0xec, 0xfd, 0xc1, 0xe5,
	0xac, 0xfa, 0x0d, 0xa0, 0xf9, 0x2d, 0x63, 0x53, 0x17, 0x6a, 0xef, 0xf1, 0x6d, 0xc6, 0xa6, 0x33,
	0xea, 0x68, 0x8d, 0x22, 0x95, 0x9c, 0x02, 0x84, 0x2a, 0xf8, 0xca, 0xb5, 0xff, 0x39, 0x9b, 0xb8,
	0xed, 0xda, 0x54, 0x2f, 0x91, 0x30, 0xca, 0x26, 0x0b, 0xdd, 0x4e, 0x38, 0x03, 0xc9, 0x73, 0xd8,
	0x10, 0x81, 0x76, 0x3b, 0xd8, 0xfb, 0x70, 0xa5, 0xf7, 0xb5, 0x52, 0x22, 0xaa, 0x8e, 0x66, 0xb4,
	0x46, 0x8b, 0xa6, 0x21, 0x40, 0x6b, 0x36, 0x12, 0xef, 0x0d, 0x74, 0x16, 0x23, 0xe2, 0x09, 0x17,
	0xab, 0xc9, 0x74, 0x34, 0x3f, 0x61, 0x96, 0xf2, 0xf7, 0x3a, 0x22, 0x1e, 0x74, 0x94, 0x16, 0x2c,
	0x96, 0x3f, 0x98, 0xb9, 0x39, 0xe0, 0x25, 0xcc, 0xdb, 0x81, 0x5e, 0x25, 0xb8, 0x77, 0x01, 0x7b,
	0xb7, 0x24, 0xfa, 0x97, 0xdd, 0x21, 0x40, 0xa2, 0xd5, 0x17, 0x1e, 0x98, 0xe2, 0x73, 0x2a, 0xcd,
	0x1c, 0x8b, 0x9c, 0x87, 0xc3, 0xe6, 0xc7, 0xf5, 0x7c, 0x30, 0xd9, 0xc4, 0x1f, 0xd6, 0xd3, 0xbf,
	0x01, 0x00, 0x00, 0xff, 0xff, 0xb6, 0x3d, 0x4c, 0xce, 0xc6, 0x04, 0x00, 0x00,
}
