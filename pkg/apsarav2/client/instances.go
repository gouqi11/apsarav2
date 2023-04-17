package client

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/pkg/errors"
	"yunion.io/x/apsarav2/pkg/apsarav2/file"
	"yunion.io/x/cloudmux/pkg/apis/compute"
)

type SInstanceNic struct {
	ResourceBaseInfo

	Ip        string
	Mac       string
	Classic   bool
	Driver    string
	NetworkId string
	SubAddr   []string
}

type SOldInstances struct {
	SResourceDriverBase

	Cpu                 int
	Config              string
	DiskSize            int
	EipAddress          string
	ImageId             string
	InnerIpAddress      string
	InstanceId          string
	InstanceName        string
	InstanceStatus      string
	InsatnceType        string
	Memory              int
	NatIpAddress        string
	NetworkType         string
	OsName              string
	PhysicalHostName    string
	PrivateIpAddress    []string
	PublicIp            string
	RegionId            string
	SecurityGroupIdList string
	SerialNumber        string
	SystemDiskSize      string
	VpcId               string
	VSwitchId           string `json:"vSwitchId"`
	ZoneId              string

	SCommonInfo
}

type SInstances struct {
	HostId       string `json:"host_id"`
	Hostname     string `json:"hostname"`
	VcpuCount    int    `json:"vcpu_count"`
	VMemSize     int    `json:"vmem_size_mb"`
	InstanceType string `json:"instance_type"`
	Bandwidth    int    `json:"band_width"`
	Throughput   int    `json:"throughput"`
	EipId        string `json:"eip_id"`
	OsArch       string
	OsType       string
	OsName       string
	Disks        []SDisk `json:"disks"`
	Nics         []SInstanceNic
	Wires        []SWire

	ResourceBaseInfo
}

var instanceStatusMap = map[string]string{
	"running":  compute.VM_RUNNING,
	"stopped":  compute.VM_READY,
	"building": "creating",
	"resizing": "change_config",
}

func (instance SOldInstances) SyncFromLocal() error {
	res := []SInstances{}
	disks, err := getDiskDetails()
	if err != nil {
		return errors.Wrap(err, "getDiskDetails")
	}

	old := []SOldInstances{}
	err = file.ReadFileDetail("instances", &old)
	if err != nil {
		return errors.Wrap(err, "ReadFileDetail")
	}
	wires, err := getRemoteWire()
	if err != nil {
		return errors.Wrap(err, "getRemoteWire")
	}
	for _, v := range old {
		wire := getWireByVpcId(wires, v.VpcId)
		configArr := strings.Split(v.Config, "/")
		diskSizeGb, _ := strconv.Atoi(configArr[1])
		temp := SInstances{
			ResourceBaseInfo: ResourceBaseInfo{
				Id:       v.InstanceId,
				Name:     v.InstanceName,
				Status:   instanceStatusMap[v.InstanceStatus],
				BaseInfo: getBaseInfo(v),
			},
			HostId:       "host-1",
			Hostname:     fmt.Sprintf("instance-%s", v.ZoneId),
			VcpuCount:    v.Cpu,
			VMemSize:     diskSizeGb * 1024,
			InstanceType: v.InsatnceType,
			Nics: []SInstanceNic{
				{ResourceBaseInfo: ResourceBaseInfo{Id: fmt.Sprintf("%s-%s", v.InstanceId, v.VSwitchId)}, NetworkId: v.VSwitchId, Ip: v.PrivateIpAddress[0]},
			},
			OsName: v.OsName,
			Wires:  []SWire{*wire},
		}
		temp.Disks = getDisksByInstanceId(disks, v.InstanceId)
		temp.EipId, err = getEipIdByInsatnceId(v.InstanceId)
		if err != nil {
			return errors.Wrap(err, "getEipIdByInsatnceId")
		}
		res = append(res, temp)
	}
	return file.WriteRemoteFile(instance.Name(), res)
}

// func getPrivateIpAddress(instanceId string, PrivateIpAddress []string) []SInstanceNic {
// 	res := []SInstanceNic{}
// 	for _, ipAddr := range PrivateIpAddress {
// 		res = append(res, SInstanceNic{
// 			Ip: ipAddr,
// 			ResourceBaseInfo: ResourceBaseInfo{
// 				Id: fmt.Sprintf("%s-%s", instanceId, ipAddr),
// 			},
// 		})
// 	}
// 	return res
// }

func (instance SOldInstances) Name() string {
	return "instances"
}

func (instance SOldInstances) IsNeedDetails() bool {
	return false
}

func (instance *SOldInstances) InitBase() {
	instance.SResourceDriverBase = SResourceDriverBase{
		ResourceName:          instance.Name(),
		ResourceIsNeedDetails: instance.IsNeedDetails(),
	}
}

func init() {
	driver := &SOldInstances{}
	ResourceRegister(driver)
}
