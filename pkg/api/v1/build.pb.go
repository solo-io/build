// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: api/v1/build.proto

package v1

import (
	fmt "fmt"
	math "math"

	proto "github.com/gogo/protobuf/proto"
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
	Version string `protobuf:"bytes,4,opt,name=version,proto3" json:"version,omitempty"`
	// the URI of the repository to be used to publish the helm charts for this build, e.g. gs://supergloo-helm/
	HelmRepository       string   `protobuf:"bytes,5,opt,name=helm_repository,json=helmRepository,proto3" json:"helm_repository,omitempty"`
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

func (m *ComputedBuildVars) GetHelmRepository() string {
	if m != nil {
		return m.HelmRepository
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
	// target helm repo for release builds
	ReleaseHelmRepository *HelmChartRepository `protobuf:"bytes,30,opt,name=release_helm_repository,json=releaseHelmRepository,proto3" json:"release_helm_repository,omitempty"`
	// target helm repo for non-release builds
	TestHelmRepository   *HelmChartRepository `protobuf:"bytes,31,opt,name=test_helm_repository,json=testHelmRepository,proto3" json:"test_helm_repository,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
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

func (m *BuildConfig) GetReleaseHelmRepository() *HelmChartRepository {
	if m != nil {
		return m.ReleaseHelmRepository
	}
	return nil
}

func (m *BuildConfig) GetTestHelmRepository() *HelmChartRepository {
	if m != nil {
		return m.TestHelmRepository
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

type HelmChartRepository struct {
	// Types that are valid to be assigned to RepositoryType:
	//	*HelmChartRepository_GoogleCloudStorage
	RepositoryType       isHelmChartRepository_RepositoryType `protobuf_oneof:"repository_type"`
	XXX_NoUnkeyedLiteral struct{}                             `json:"-"`
	XXX_unrecognized     []byte                               `json:"-"`
	XXX_sizecache        int32                                `json:"-"`
}

func (m *HelmChartRepository) Reset()         { *m = HelmChartRepository{} }
func (m *HelmChartRepository) String() string { return proto.CompactTextString(m) }
func (*HelmChartRepository) ProtoMessage()    {}
func (*HelmChartRepository) Descriptor() ([]byte, []int) {
	return fileDescriptor_05c5040d8276e5da, []int{10}
}
func (m *HelmChartRepository) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HelmChartRepository.Unmarshal(m, b)
}
func (m *HelmChartRepository) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HelmChartRepository.Marshal(b, m, deterministic)
}
func (m *HelmChartRepository) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HelmChartRepository.Merge(m, src)
}
func (m *HelmChartRepository) XXX_Size() int {
	return xxx_messageInfo_HelmChartRepository.Size(m)
}
func (m *HelmChartRepository) XXX_DiscardUnknown() {
	xxx_messageInfo_HelmChartRepository.DiscardUnknown(m)
}

var xxx_messageInfo_HelmChartRepository proto.InternalMessageInfo

type isHelmChartRepository_RepositoryType interface {
	isHelmChartRepository_RepositoryType()
}

type HelmChartRepository_GoogleCloudStorage struct {
	GoogleCloudStorage *GoogleCloudStorage `protobuf:"bytes,11,opt,name=google_cloud_storage,json=googleCloudStorage,proto3,oneof"`
}

func (*HelmChartRepository_GoogleCloudStorage) isHelmChartRepository_RepositoryType() {}

func (m *HelmChartRepository) GetRepositoryType() isHelmChartRepository_RepositoryType {
	if m != nil {
		return m.RepositoryType
	}
	return nil
}

func (m *HelmChartRepository) GetGoogleCloudStorage() *GoogleCloudStorage {
	if x, ok := m.GetRepositoryType().(*HelmChartRepository_GoogleCloudStorage); ok {
		return x.GoogleCloudStorage
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*HelmChartRepository) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _HelmChartRepository_OneofMarshaler, _HelmChartRepository_OneofUnmarshaler, _HelmChartRepository_OneofSizer, []interface{}{
		(*HelmChartRepository_GoogleCloudStorage)(nil),
	}
}

func _HelmChartRepository_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*HelmChartRepository)
	// repository_type
	switch x := m.RepositoryType.(type) {
	case *HelmChartRepository_GoogleCloudStorage:
		_ = b.EncodeVarint(11<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.GoogleCloudStorage); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("HelmChartRepository.RepositoryType has unexpected type %T", x)
	}
	return nil
}

func _HelmChartRepository_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*HelmChartRepository)
	switch tag {
	case 11: // repository_type.google_cloud_storage
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(GoogleCloudStorage)
		err := b.DecodeMessage(msg)
		m.RepositoryType = &HelmChartRepository_GoogleCloudStorage{msg}
		return true, err
	default:
		return false, nil
	}
}

func _HelmChartRepository_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*HelmChartRepository)
	// repository_type
	switch x := m.RepositoryType.(type) {
	case *HelmChartRepository_GoogleCloudStorage:
		s := proto.Size(x.GoogleCloudStorage)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type GoogleCloudStorage struct {
	BucketUrl            string   `protobuf:"bytes,1,opt,name=bucket_url,json=bucketUrl,proto3" json:"bucket_url,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GoogleCloudStorage) Reset()         { *m = GoogleCloudStorage{} }
func (m *GoogleCloudStorage) String() string { return proto.CompactTextString(m) }
func (*GoogleCloudStorage) ProtoMessage()    {}
func (*GoogleCloudStorage) Descriptor() ([]byte, []int) {
	return fileDescriptor_05c5040d8276e5da, []int{11}
}
func (m *GoogleCloudStorage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GoogleCloudStorage.Unmarshal(m, b)
}
func (m *GoogleCloudStorage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GoogleCloudStorage.Marshal(b, m, deterministic)
}
func (m *GoogleCloudStorage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GoogleCloudStorage.Merge(m, src)
}
func (m *GoogleCloudStorage) XXX_Size() int {
	return xxx_messageInfo_GoogleCloudStorage.Size(m)
}
func (m *GoogleCloudStorage) XXX_DiscardUnknown() {
	xxx_messageInfo_GoogleCloudStorage.DiscardUnknown(m)
}

var xxx_messageInfo_GoogleCloudStorage proto.InternalMessageInfo

func (m *GoogleCloudStorage) GetBucketUrl() string {
	if m != nil {
		return m.BucketUrl
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
	proto.RegisterType((*HelmChartRepository)(nil), "build.solo.io.HelmChartRepository")
	proto.RegisterType((*GoogleCloudStorage)(nil), "build.solo.io.GoogleCloudStorage")
}

func init() { proto.RegisterFile("api/v1/build.proto", fileDescriptor_05c5040d8276e5da) }

var fileDescriptor_05c5040d8276e5da = []byte{
	// 676 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x55, 0xcd, 0x4e, 0xdb, 0x4c,
	0x14, 0x25, 0x90, 0x0f, 0x92, 0x4b, 0x08, 0x5f, 0x06, 0x10, 0x2e, 0x88, 0x96, 0x8e, 0xd4, 0x3f,
	0xa9, 0x0a, 0x02, 0xd4, 0x4d, 0x37, 0x15, 0xa1, 0x55, 0xc3, 0xa2, 0x12, 0x1d, 0x7e, 0x54, 0xb1,
	0xa8, 0x35, 0xb1, 0x87, 0xc1, 0xc5, 0xf1, 0xb8, 0xe3, 0x71, 0xd4, 0x74, 0xd3, 0xa7, 0xe9, 0x3b,
	0xf4, 0x35, 0xfa, 0x38, 0xdd, 0x55, 0xbe, 0x76, 0xfe, 0x6c, 0x53, 0xb1, 0xbc, 0xc7, 0xe7, 0x9e,
	0x7b, 0xe6, 0x9e, 0x99, 0x04, 0x08, 0x0f, 0xbd, 0xbd, 0xc1, 0xfe, 0x5e, 0x2f, 0xf6, 0x7c, 0xb7,
	0x1d, 0x6a, 0x65, 0x14, 0x59, 0x49, 0x8b, 0x48, 0xf9, 0xaa, 0xed, 0x29, 0xaa, 0xa0, 0xd6, 0x49,
	0x00, 0x16, 0x07, 0xe4, 0x25, 0x54, 0xa3, 0x50, 0x38, 0x56, 0x65, 0xb7, 0xf2, 0x7c, 0xf9, 0xc0,
	0x6a, 0xcf, 0x30, 0xdb, 0x48, 0x3b, 0x0b, 0x85, 0xc3, 0x90, 0x45, 0x5e, 0xc1, 0xa2, 0xa3, 0x82,
	0x6b, 0x4f, 0x5a, 0xf3, 0xc8, 0xdf, 0x29, 0xe3, 0xb3, 0x38, 0x38, 0x46, 0x12, 0xcb, 0xc8, 0xf4,
	0x67, 0x05, 0x9a, 0xb3, 0x9f, 0xc8, 0x11, 0x34, 0xb1, 0xd5, 0x16, 0xc1, 0xc0, 0x1e, 0x70, 0x1d,
	0x65, 0x0e, 0xb6, 0xcb, 0x14, 0xdf, 0x05, 0x83, 0x4b, 0xae, 0x23, 0xd6, 0xe8, 0x4d, 0x55, 0xe4,
	0x14, 0xd6, 0x1c, 0xd5, 0x0f, 0x63, 0x23, 0x5c, 0x3b, 0xd5, 0x42, 0x9d, 0xd4, 0xd9, 0x6e, 0x4e,
	0xe7, 0x38, 0x63, 0xa2, 0x1e, 0x8a, 0xb5, 0x9c, 0x3c, 0x44, 0x4f, 0xa1, 0x31, 0x3d, 0x8f, 0x3c,
	0x80, 0x5a, 0x2a, 0xec, 0xb9, 0x68, 0xaf, 0xce, 0x96, 0xb0, 0x3e, 0x71, 0xc9, 0x13, 0x68, 0x1a,
	0x2e, 0xa5, 0x70, 0xed, 0x81, 0xd0, 0x91, 0xa7, 0x02, 0x9c, 0x5b, 0x67, 0x2b, 0x29, 0x7a, 0x99,
	0x82, 0xf4, 0x57, 0x05, 0x5a, 0x85, 0xd1, 0xc4, 0x82, 0x25, 0x2d, 0x7c, 0xc1, 0x23, 0x81, 0xb2,
	0x35, 0x36, 0x2a, 0xc9, 0x36, 0xd4, 0xbd, 0x3e, 0x97, 0xc2, 0x36, 0x5c, 0x66, 0x8a, 0x35, 0x04,
	0xce, 0xb9, 0x24, 0x2f, 0xe0, 0x7f, 0x47, 0x05, 0x86, 0x7b, 0x81, 0xd0, 0x76, 0xa8, 0xc5, 0xb5,
	0xf7, 0xcd, 0x5a, 0x40, 0xce, 0xea, 0x18, 0x3f, 0x45, 0x38, 0x99, 0x30, 0xf2, 0x55, 0x4d, 0x8d,
	0x67, 0x25, 0x79, 0x06, 0xab, 0x37, 0xc2, 0xef, 0xdb, 0x5a, 0x84, 0x2a, 0xf2, 0x8c, 0xd2, 0x43,
	0xeb, 0x3f, 0x64, 0x34, 0x13, 0x98, 0x8d, 0x51, 0xfa, 0x06, 0xea, 0xe3, 0xf8, 0xc9, 0xc1, 0x38,
	0xf8, 0x34, 0xa6, 0xad, 0xb2, 0x98, 0x72, 0xa9, 0xff, 0x99, 0x87, 0xe5, 0x29, 0x9c, 0x7c, 0x86,
	0xad, 0xec, 0x98, 0xf6, 0xe4, 0x18, 0x5a, 0x48, 0x2f, 0x32, 0x7a, 0x68, 0xad, 0xdf, 0x11, 0x5b,
	0x46, 0x64, 0x19, 0x8f, 0x59, 0x99, 0x46, 0xe1, 0x0b, 0xf9, 0x04, 0x9b, 0x46, 0x44, 0xa6, 0x4c,
	0x7c, 0xe3, 0x9e, 0xe2, 0x1b, 0x89, 0x40, 0x51, 0xf9, 0x0a, 0x36, 0x47, 0xce, 0xf3, 0xbb, 0x7b,
	0x88, 0xca, 0x34, 0xa7, 0xdc, 0x15, 0x7e, 0xff, 0xf8, 0x86, 0x6b, 0x33, 0xd9, 0x27, 0xdb, 0xc8,
	0x24, 0xba, 0x33, 0x6b, 0x26, 0xe7, 0xb0, 0x8e, 0xae, 0xf3, 0xc2, 0x8f, 0xee, 0x2d, 0x4c, 0x92,
	0xfe, 0x59, 0x55, 0xfa, 0x1b, 0xef, 0x5d, 0xfe, 0x1c, 0xfb, 0x50, 0xfd, 0x1a, 0xf3, 0xa1, 0x05,
	0xa5, 0x4f, 0xed, 0x63, 0xcc, 0x87, 0x23, 0x6a, 0x77, 0x8e, 0x21, 0x95, 0x1c, 0x01, 0xb8, 0xca,
	0xb9, 0x15, 0xda, 0xbe, 0x89, 0x7b, 0xd6, 0x72, 0xe9, 0x1e, 0xdf, 0x22, 0xa1, 0x1b, 0xf7, 0xa6,
	0xba, 0xeb, 0xee, 0x08, 0x24, 0xaf, 0x61, 0x41, 0x3a, 0xda, 0x6a, 0x60, 0xef, 0xd3, 0x5c, 0xef,
	0x7b, 0xa5, 0xa4, 0x5f, 0x0c, 0xb3, 0x3b, 0xc7, 0x92, 0xa6, 0x0e, 0x40, 0x6d, 0x14, 0x22, 0xfd,
	0x00, 0x8d, 0x69, 0x8b, 0xf8, 0x3a, 0x93, 0x48, 0x62, 0xed, 0x8f, 0x5f, 0x27, 0x8f, 0xc4, 0x85,
	0xf6, 0x09, 0x85, 0x86, 0xd2, 0x92, 0x07, 0xde, 0x77, 0x6e, 0x26, 0x6f, 0x73, 0x06, 0xa3, 0x6b,
	0xd0, 0x2a, 0x18, 0xa7, 0x67, 0xb0, 0x79, 0x87, 0xa3, 0x7f, 0x8d, 0xdb, 0x01, 0x08, 0xb5, 0xfa,
	0x22, 0x1c, 0x93, 0xfc, 0x52, 0xa4, 0xc3, 0xea, 0x19, 0x72, 0xe2, 0xd2, 0x1f, 0xb0, 0x56, 0x92,
	0x1b, 0xb9, 0x80, 0x75, 0x89, 0xb3, 0x6c, 0xc7, 0x57, 0xb1, 0x6b, 0x47, 0x46, 0x69, 0x2e, 0x45,
	0xb6, 0xe4, 0xc7, 0xe5, 0x8b, 0x4a, 0x98, 0x67, 0x29, 0xb1, 0x3b, 0xc7, 0x88, 0x2c, 0xa0, 0x9d,
	0x16, 0xac, 0x4e, 0xae, 0x91, 0x6d, 0x86, 0xa1, 0xa0, 0x87, 0x40, 0x8a, 0xed, 0x89, 0xeb, 0x5e,
	0xec, 0xdc, 0x0a, 0x33, 0x75, 0xa4, 0x7a, 0x8a, 0x5c, 0x68, 0xbf, 0x53, 0xbd, 0x9a, 0x1f, 0xec,
	0xf7, 0x16, 0xf1, 0x1f, 0xe4, 0xf0, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x68, 0x37, 0x1e, 0xa0,
	0x57, 0x06, 0x00, 0x00,
}
