package client

import (
	"fmt"

	"github.com/golang-plus/errors"
	"yunion.io/x/apsarav2/pkg/apsarav2/file"
	"yunion.io/x/pkg/util/netutils"
)

type SNetwork struct {
	ResourceBaseInfo

	WireId  string
	IpStart string
	IpEnd   string
	IpMask  int8
	Gatway  string
}

func getNetworkDetails() ([]SNetwork, error) {
	networkMap := map[string]struct{}{}
	old := []SOldInstances{}
	err := file.ReadFileDetail("instances", &old)
	if err != nil {
		return nil, errors.Wrap(err, "ReadFileDetail")
	}
	res := []SNetwork{}
	vpcList := []SVpc{}
	obj, err := file.ReadRemoteFile("vpcs.json")
	if err != nil {
		return nil, errors.Wrap(err, "ReadFileDetail")
	}
	obj.Unmarshal(&vpcList)
	vpcMap := map[string]SVpc{}
	for _, vpc := range vpcList {
		vpcMap[vpc.Id] = vpc
	}
	for _, v := range old {
		if _, isExist := networkMap[v.VSwitchId]; isExist {
			continue
		}
		prefix, _ := netutils.NewIPV4Prefix(vpcMap[v.VpcId].CidrBlock)
		iprange := prefix.ToIPRange()
		networkMap[v.VSwitchId] = struct{}{}
		temp := SNetwork{
			ResourceBaseInfo: ResourceBaseInfo{
				Id:     v.VSwitchId,
				Name:   v.VSwitchId,
				Status: "available",
			},
			IpStart: iprange.StartIp().String(),
			IpEnd:   iprange.EndIp().String(),
			WireId:  fmt.Sprintf("%s-%s", v.VpcId, v.RegionId),
		}
		res = append(res, temp)
	}
	return res, nil
}
