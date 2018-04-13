// Code generated by protoc-gen-go. DO NOT EDIT.
// source: 0.proto

/*
Package pb is a generated protocol buffer package.

It is generated from these files:
	0.proto
	app.proto
	cluster.proto
	job.proto
	pilot.proto
	repo.proto
	repo_indexer.proto
	runtime.proto
	task.proto

It has these top-level messages:
	CreateAppRequest
	CreateAppResponse
	ModifyAppRequest
	ModifyAppResponse
	DeleteAppRequest
	DeleteAppResponse
	App
	DescribeAppsRequest
	DescribeAppsResponse
	CreateAppVersionRequest
	CreateAppVersionResponse
	ModifyAppVersionRequest
	ModifyAppVersionResponse
	DeleteAppVersionRequest
	DeleteAppVersionResponse
	AppVersion
	DescribeAppVersionsRequest
	DescribeAppVersionsResponse
	GetAppVersionPackageRequest
	GetAppVersionPackageResponse
	CreateClusterRequest
	CreateClusterResponse
	ModifyClusterRequest
	ModifyClusterResponse
	ModifyClusterNodeRequest
	ModifyClusterNodeResponse
	DeleteClustersRequest
	DeleteClustersResponse
	UpgradeClusterRequest
	UpgradeClusterResponse
	RollbackClusterRequest
	RollbackClusterResponse
	ResizeClusterRequest
	ResizeClusterResponse
	AddClusterNodesRequest
	AddClusterNodesResponse
	DeleteClusterNodesRequest
	DeleteClusterNodesResponse
	UpdateClusterEnvRequest
	UpdateClusterEnvResponse
	ClusterCommon
	ClusterNode
	ClusterRole
	ClusterLoadbalancer
	ClusterLink
	Cluster
	DescribeClustersRequest
	DescribeClustersResponse
	DescribeClusterNodesRequest
	DescribeClusterNodesResponse
	StopClustersRequest
	StopClustersResponse
	StartClustersRequest
	StartClustersResponse
	RecoverClustersRequest
	RecoverClustersResponse
	CeaseClustersRequest
	CeaseClustersResponse
	CreateJobRequest
	CreateJobResponse
	Job
	DescribeJobsRequest
	DescribeJobsResponse
	HandleSubtaskRequest
	HandleSubtaskResponse
	GetSubtaskStatusRequest
	SubtaskStatus
	GetSubtaskStatusResponse
	PilotInfo
	FrontgateInfo
	DroneInfo
	ConfdInfo
	GetConfdInfoRequest
	StartConfdRequest
	StopConfdRequest
	RegisterMetadataRequest
	DeregisterMetadataRequest
	RegisterCmdRequest
	DeregisterCmdRequest
	CreateRepoRequest
	CreateRepoResponse
	ModifyRepoRequest
	ModifyRepoResponse
	DeleteRepoRequest
	DeleteRepoResponse
	RepoLabel
	RepoSelector
	Repo
	DescribeReposRequest
	DescribeReposResponse
	IndexRepoRequest
	IndexRepoResponse
	RepoEvent
	DescribeRepoEventsRequest
	DescribeRepoEventsResponse
	RuntimeLabel
	Runtime
	CreateRuntimeRequest
	CreateRuntimeResponse
	DescribeRuntimesRequest
	DescribeRuntimesResponse
	ModifyRuntimeRequest
	ModifyRuntimeResponse
	DeleteRuntimeRequest
	DeleteRuntimeResponse
	DescribeRuntimeProviderZonesRequest
	DescribeRuntimeProviderZonesResponse
	CreateTaskRequest
	CreateTaskResponse
	Task
	DescribeTasksRequest
	DescribeTasksResponse
*/
package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

func init() { proto.RegisterFile("0.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 162 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x62, 0x37, 0xd0, 0x2b, 0x28,
	0xca, 0x2f, 0xc9, 0x17, 0xe2, 0xca, 0x2f, 0x48, 0xcd, 0x2b, 0xc8, 0x2c, 0x29, 0xca, 0xac, 0x90,
	0xd2, 0x01, 0x0b, 0x25, 0xeb, 0xa6, 0xa7, 0xe6, 0xe9, 0x16, 0x97, 0x27, 0xa6, 0xa7, 0xa7, 0x16,
	0xe9, 0xe7, 0x17, 0x94, 0x64, 0xe6, 0xe7, 0x15, 0xeb, 0x27, 0xe6, 0xe5, 0xe5, 0x97, 0x24, 0x82,
	0xd9, 0x10, 0x9d, 0x4e, 0x15, 0x93, 0x1c, 0x0b, 0x85, 0x3c, 0xb8, 0x84, 0xfc, 0x0b, 0x52, 0xf3,
	0x02, 0xc0, 0x06, 0x28, 0x04, 0x14, 0xe5, 0x67, 0xa5, 0x26, 0x97, 0x28, 0x69, 0x63, 0x13, 0x15,
	0x12, 0xcd, 0x28, 0x29, 0x29, 0x28, 0xb6, 0xd2, 0xd7, 0x47, 0x58, 0xa9, 0x97, 0x99, 0x6f, 0xc4,
	0x6a, 0xa0, 0x67, 0xa0, 0x67, 0xa8, 0xc5, 0xc8, 0x68, 0x24, 0x90, 0x58, 0x50, 0x90, 0x93, 0x99,
	0x0c, 0xb6, 0x45, 0x3f, 0xab, 0x38, 0x3f, 0xcf, 0x0a, 0x43, 0x24, 0x8a, 0xa9, 0x20, 0x29, 0x89,
	0x0d, 0xec, 0x00, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0xaa, 0x3d, 0xcf, 0xeb, 0xc5, 0x00,
	0x00, 0x00,
}
