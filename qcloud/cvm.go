/*

LICENSE:  MIT
Author:   sine
Email:    sinerwr@gmail.com

*/

package qcloudSDK

type CVM struct {
	Response CVMResponse `json:"response"`
}

type CVMResponse struct {
	RequestId   string
	TotalCount  int64
	InstanceSet []CVMInstanceSet
}

type CVMInstanceSet struct {
	Placement           CVMPlacement
	InstanceId          string
	InstanceType        string
	CPU                 int64
	Memory              int64
	InstanceName        string
	InstanceChargeType  string
	SystemDisk          CVMSystemDisk
	DataDisks           []CVMDataDisk
	PrivateIpAddresses  []string
	PublicIpAddresses   []string
	InternetAccessible  CVMInternetAccessible
	VirtualPrivateCloud CVMVirtualPrivateCloud
	ImageId             string
	RenewFlag           string
	CreatedTime         string
	ExpiredTime         string
}

type CVMPlacement struct {
	Zone      string
	ProjectId int64
	HostIds   []string
}

type CVMSystemDisk struct {
	DiskType string
	DiskId   string
	DiskSize int64
}

type CVMDataDisk struct {
	DiskType string
	DiskId   string
	DiskSize int64
}
type CVMInternetAccessible struct {
	InternetChargeType      string
	InternetMaxBandwidthOut int64
	PublicIpAssigned        bool
}

type CVMVirtualPrivateCloud struct {
	VpcId              string
	SubnetId           string
	AsVpcGateway       bool
	PrivateIpAddresses []string
}
