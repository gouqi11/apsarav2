package client

import (
	"fmt"

	"yunion.io/x/apsarav2/pkg/apsarav2/file"
	"yunion.io/x/pkg/errors"
)

type SWire struct {
	ResourceBaseInfo

	WireId    string
	VpcId     string
	ZoneId    string
	Bandwidth int
}

func getWireDetails() ([]SWire, error) {
	zones, err := getZoneDetails()
	if err != nil {
		return nil, errors.Wrap(err, "getZoneDetails")
	}
	vpcsOld := []SOldVpc{}
	file.ReadFileDetail("vpcs", &vpcsOld)
	res := []SWire{}
	for _, vpc := range vpcsOld {
		zoneId := zones[0].Id
		for _, zone := range zones {
			if zone.RegionId == vpc.RegionId {
				zoneId = zone.Id
			}
		}
		temp := SWire{
			VpcId:  vpc.VpcId,
			ZoneId: zoneId,
			WireId: fmt.Sprintf("%s-%s", vpc.VpcId, vpc.RegionId),
			ResourceBaseInfo: ResourceBaseInfo{
				Id:   fmt.Sprintf("%s-%s", vpc.VpcId, vpc.RegionId),
				Name: fmt.Sprintf("%s-%s", vpc.VpcName, vpc.RegionName),
			},
		}
		res = append(res, temp)
	}
	return res, nil
}

func getRemoteWire() ([]SWire, error) {
	obj, err := file.ReadRemoteFile("wires.json")
	if err != nil {
		return nil, err
	}
	res := []SWire{}
	obj.Unmarshal(&res)
	return res, nil
}

func getWireByVpcId(wires []SWire, id string) *SWire {
	for _, v := range wires {
		if v.VpcId == id {
			return &v
		}
	}
	return nil
}
