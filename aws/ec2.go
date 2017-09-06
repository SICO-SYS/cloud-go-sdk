/*

LICENSE:  MIT
Author:   sine
Email:    sinerwr@gmail.com

*/

package awsSDK

import (
	"encoding/xml"
)

type EC2DescribeInstances struct {
	DescribeInstances xml.Name      `xml:"DescribeInstancesResponse"`
	RequestID         string        `xml:"requestId"`
	NextToken         string        `xml:"nextToken"`
	ReservationSet    []Reservation `xml:"reservationSet>item"`
}

type Reservation struct {
	OwnerId       string     `xml:"ownerId"`
	ReservationId string     `xml:"reservationId"`
	InstancesSet  []Instance `xml:"instancesSet>item"`
}

type Instance struct {
	AmiLaunchIndex           int64                        `xml:"amiLaunchIndex"`
	Architecture             string                       `xml:"architecture"`
	BlockDeviceMapping       []InstanceBlockDeviceMapping `xml:"blockDeviceMapping>item"`
	ClientToken              string                       `xml:"clientToken"`
	DnsName                  string                       `xml:"dnsName"`
	EbsOptimized             bool                         `xml:"ebsOptimized"`
	ElasticGpuAssociationSet []ElasticGpuAssociation      `xml:"elasticGpuAssociationSet>item"`
	EnaSupport               bool                         `xml:"enaSupport"`
	GroupSet                 []GroupIdentifier            `xml:"groupSet>item"`
	Hypervisor               string                       `xml:"hypervisor"`
	IamInstanceProfile       IamInstanceProfile           `xml:"iamInstanceProfile"`
	ImageID                  string                       `xml:"imageID"`
	InstanceID               string                       `xml:"instanceId"`
	InstanceLifecycle        string                       `xml:"instanceLifecycle"`
	InstanceState            InstanceState                `xml:"InstanceState"`
	InstanceType             string                       `xml:"instanceType"`
	IpAddress                string                       `xml:"ipAddress"`
	KernelID                 string                       `xml:"kernelId"`
	KeyName                  string                       `xml:"keyName"`
	LaunchTime               string                       `xml:"launchTime"`
	Monitoring               Monitoring                   `xml:"monitoring"`
	NetworkInterfaceSet      []InstanceNetworkInterface   `xml:"networkInterfaceSet>item"`
	Placement                Placement                    `xml:"placement"`
	Platform                 string                       `xml:"platform"`
	PrivateDnsName           string                       `xml:"privateDnsName"`
	PrivateIpAddress         string                       `xml:"privateIpAddress"`
	ProductCodes             []ProductCode                `xml:"productCodes>item"`
	RamdiskID                string                       `xml:"ramdiskId"`
	Reason                   string                       `xml:"reason"`
	RootDeviceName           string                       `xml:"rootDeviceName"`
	RootDeviceType           string                       `xml:"rootDeviceType"`
	SourceDestCheck          bool                         `xml:"sourceDestCheck"`
	SpotInstanceRequestID    string                       `xml:"spotInstanceRequestId"`
	SriovNetSupport          string                       `xml:"sriovNetSupport"`
	StateReason              StateReason                  `xml:"stateReason"`
	SubnetID                 string                       `xml:"subnetId"`
	TagSet                   []Tag                        `xml:"tagSet>item"`
	VirtualizationType       string                       `xml:"virtualizationType"`
	VpcID                    string                       `xml:"vpcId"`
}

type InstanceBlockDeviceMapping struct {
	DeviceName string                 `xml:"deviceName"`
	Ebs        EbsInstanceBlockDevice `xml:"ebs"`
}

type EbsInstanceBlockDevice struct {
	AttachTime          string `xml:"attachTime"`
	DeleteOnTermination bool   `xml:"deleteOnTermination"`
	Status              string `xml:"status"`
	VolumeID            string `xml:"volumeId"`
}

type ElasticGpuAssociation struct {
	ElasticGpuAssociationID    string `xml:"elasticGpuAssociationId"`
	ElasticGpuAssociationState string `xml:"elasticGpuAssociationState"`
	ElasticGpuAssociationTime  string `xml:"elasticGpuAssociationTime"`
	ElasticGpuID               string `xml:"elasticGpuId"`
}

type GroupIdentifier struct {
	GroupID   string `xml:"groupId"`
	GroupName string `xml:"groupName"`
}

type IamInstanceProfile struct {
	ARN string `xml:"arn"`
	ID  string `xml:"id"`
}

type InstanceState struct {
	Code int64  `xml:"code"`
	Name string `xml:"name"`
}

type Monitoring struct {
	State string `xml:"state"`
}

type InstanceNetworkInterface struct {
	Association           InstanceNetworkInterfaceAssociation `xml:"association"`
	Attachment            InstanceNetworkInterfaceAttachment  `xml:"attachment"`
	Description           string                              `xml:"description"`
	GroupSet              []GroupIdentifier                   `xml:"groupSet>item"`
	Ipv6AddressesSet      []InstanceIpv6Address               `xml:"ipv6AddressesSet>item"`
	MacAddress            string                              `xml:"macAddress"`
	NetworkInterfaceID    string                              `xml:"networkInterfaceId"`
	OwnerID               string                              `xml:"ownerId"`
	PrivateDnsName        string                              `xml:"privateDnsName"`
	PrivateIpAddress      string                              `xml:"privateIpAddress"`
	PrivateIpAddressesSet []InstancePrivateIpAddress          `xml:"privateIpAddressesSet>item"`
	SourceDestCheck       bool                                `xml:"sourceDestCheck"`
	Status                string                              `xml:"status"`
	SubnetID              string                              `xml:"subnetId"`
	VpcID                 string                              `xml:"vpcId"`
}

type InstanceNetworkInterfaceAssociation struct {
	IpOwnerID     string `xml:"ipOwnerId"`
	PublicDnsName string `xml:"publicDnsName"`
	PublicIp      string `xml:"publicIp"`
}

type InstanceNetworkInterfaceAttachment struct {
	attachmentId        string `xml:"attachmentId"`
	attachTime          string `xml:"attachTime"`
	deleteOnTermination bool   `xml:"deleteOnTermination"`
	deviceIndex         int64  `xml:"deviceIndex"`
	status              string `xml:"status"`
}

type InstanceIpv6Address struct {
	Ipv6Address string `xml:"ipv6Address"`
}

type InstancePrivateIpAddress struct {
	Association      InstanceNetworkInterfaceAssociation `xml:"association"`
	Primary          bool                                `xml:"primary"`
	PrivateDnsName   string                              `xml:"privateDnsName"`
	PrivateIpAddress string                              `xml:"privateIpAddress"`
}

type Placement struct {
	affinity         string `xml:"affinity"`
	availabilityZone string `xml:"availabilityZone"`
	groupName        string `xml:"groupName"`
	hostId           string `xml:"hostId"`
	spreadDomain     string `xml:"spreadDomain"`
	tenancy          string `xml:"tenancy"`
}
type ProductCode struct {
	ProductCode string `xml:"productCode"`
	Type        string `xml:"type"`
}
type StateReason struct {
	Code    string `xml:"code"`
	Message string `xml:"message"`
}
type Tag struct {
	Key   string `xml:"key"`
	Value string `xml:"value"`
}
