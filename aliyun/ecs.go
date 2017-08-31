/*

LICENSE:  MIT
Author:   sine
Email:    sinerwr@gmail.com

*/

package aliyunSDK

type Aliyun struct {
	PageNumber int64
	TotalCount int64
	PageSize   int64
	RequestId  string
	Instances  Instances
}

type Instances struct {
	Instance []Instance
}

type Instance struct {
	InstanceId              string
	InstanceName            string
	Description             string
	ImageId                 string
	RegionId                string
	ZoneId                  string
	CPU                     int64
	Memory                  int64
	InstanceType            string
	InstanceTypeFamily      string
	HostName                string
	SerialNumber            string
	Status                  string
	SecurityGroupIds        SecurityGroupIds
	PublicIpAddress         PublicIpAddress
	InternetMaxBandwidthIn  int64
	InternetMaxBandwidthOut int64
	InternetChargeType      string
	CreationTime            string
	VpcAttributes           VpcAttributes
	EipAddress              EipAddress
	InnerIpAddress          InnerIpAddress
	InstanceNetworkType     string
	OperationLocks          OperationLocks
	InstanceChargeType      string
	SpotStrategy            string
	DeviceAvailable         bool
	DeploymentSetId         string
	NetworkInterfaces       NetworkInterfaces
	IoOptimized             bool
	ExpiredTime             string
	KeyPairName             string
	GPUAmount               int64
	GPUSpec                 string
	ClusterId               string
	OSType                  string
	OSName                  string
}

type SecurityGroupIds struct {
	SecurityGroupId []string
}

type PublicIpAddress struct {
	IpAddress []string
}

type VpcAttributes struct {
	VpcId            string
	VSwitchId        string
	PrivateIpAddress PrivateIpAddress
	NatIpAddress     string
}

type PrivateIpAddress struct {
	IpAddress []string
}

type EipAddress struct {
	AllocationId       string
	IpAddress          string
	Bandwidth          int64
	InternetChargeType string
}

type InnerIpAddress struct {
	IpAddress []string
}

type OperationLocks struct {
	LockReason []string
}

type NetworkInterfaces struct {
	NetworkInterface []NetworkInterface
}

type NetworkInterface struct {
	NetworkInterfaceId string
	PrimaryIpAddress   string
	MacAddress         string
}
