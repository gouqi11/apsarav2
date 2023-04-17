package client

import (
	"github.com/pkg/errors"
	"yunion.io/x/apsarav2/pkg/apsarav2/file"
)

type SOldLoadbalancer struct {
	SResourceDriverBase

	SCommonInfo

	AddressType        string
	InternetChargeType string
	LoadBalancerId     string
	LoadBalancerName   string
	NetworkType        string
	RegionId           string
	SlbBandwidth       int
	SlbIp              string
	SlbStatus          string
	VSwitchId          string
	VpcId              string
}

type SLoadbalancer struct {
	ResourceBaseInfo

	RegionId     string
	Address      string
	AddressType  string
	NetworkType  string
	VpcId        string
	ZoneId       string
	Zone1Id      string
	InstanceType string
	ChargeType   string
	Bandwidth    int
	NetworkIds   []string
}

var loadbalancersStatusMap = map[string]string{
	"active":   "enabled",
	"inactive": "disabled",
	"locked":   "disabled",
}

var loadbalancersChargeTypeMap = map[string]string{
	"paybytraffic":   "traffic",
	"paybybandwidth": "bandwidth",
}

func (lb SOldLoadbalancer) SyncFromLocal() error {
	old := []SOldLoadbalancer{}
	file.ReadFileDetail(lb.Name(), &old)
	res := []SLoadbalancer{}
	zones, err := getZoneDetails()
	if err != nil {
		return errors.Wrap(err, "getZoneDetails")
	}
	for _, v := range old {
		zoneId := zones[0].Id
		for _, zone := range zones {
			if zone.RegionId == v.RegionId {
				zoneId = zone.Id
			}
		}
		temp := SLoadbalancer{
			RegionId:    v.RegionId,
			NetworkType: v.NetworkType,
			AddressType: v.AddressType,
			Address:     v.SlbIp,
			Bandwidth:   v.SlbBandwidth,
			ZoneId:      zoneId,
			VpcId:       v.VpcId,
			ChargeType:  loadbalancersChargeTypeMap[v.InternetChargeType],
			ResourceBaseInfo: ResourceBaseInfo{
				Id:       v.LoadBalancerId,
				Name:     v.LoadBalancerName,
				Status:   loadbalancersStatusMap[v.SlbStatus],
				BaseInfo: getBaseInfo(v),
			},
		}
		res = append(res, temp)
	}
	return file.WriteRemoteFile(lb.Name(), res)
}

func (lb SOldLoadbalancer) Name() string {
	return "loadbalancers"
}

func (lb SOldLoadbalancer) IsNeedDetails() bool {
	return true
}

func (lb *SOldLoadbalancer) InitBase() {
	lb.SResourceDriverBase = SResourceDriverBase{
		ResourceName:          lb.Name(),
		ResourceIsNeedDetails: lb.IsNeedDetails(),
	}
}

func init() {
	driver := &SOldLoadbalancer{}
	ResourceRegister(driver)
}
