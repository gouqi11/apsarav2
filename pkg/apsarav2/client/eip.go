package client

import (
	"yunion.io/x/apsarav2/pkg/apsarav2/file"
	"yunion.io/x/pkg/errors"
)

type SOldEip struct {
	SResourceDriverBase

	AllocationId   string
	AllocationTime string
	Bandwidth      int
	IpAddress      string
	InstanceId     string
	RegionId       string
	Status         string

	SCommonInfo
}

type SEip struct {
	IpAddr      string `json:"ip_addr"`
	RegionId    string `json:"region_id"`
	AssociateId string `json:"associate_id"`
	Bandwidth   int    `json:"bandwidth"`
	Mode        string

	ResourceBaseInfo
}

var eipStatusMap = map[string]string{
	"inUse":     "running",
	"Available": "ready",
}

func (eip SOldEip) SyncFromLocal() error {
	res := []SEip{}
	old := []SOldEip{}
	err := file.ReadFileDetail("eips", &old)
	if err != nil {
		return errors.Wrap(err, "ReadFileDetail")
	}
	for _, v := range old {
		temp := SEip{
			RegionId: v.RegionId,
			ResourceBaseInfo: ResourceBaseInfo{
				Id:       v.AllocationId,
				Name:     v.AllocationId,
				Status:   eipStatusMap[v.Status],
				BaseInfo: getBaseInfo(v),
			},
			IpAddr:      v.IpAddress,
			AssociateId: v.InstanceId,
			Bandwidth:   v.Bandwidth,
		}
		res = append(res, temp)
	}
	return file.WriteRemoteFile(eip.Name(), res)
}

func getEipIdByInsatnceId(instanceId string) (string, error) {
	old := []SOldEip{}
	err := file.ReadFileDetail("eips", &old)
	if err != nil {
		return "", errors.Wrap(err, "ReadFileDetail")
	}
	for _, v := range old {
		if v.InstanceId == instanceId {
			return v.AllocationId, nil
		}
	}
	return "", nil
}

func (eip SOldEip) Name() string {
	return "eips"
}

func (eip SOldEip) IsNeedDetails() bool {
	return false
}

func (eip *SOldEip) InitBase() {
	eip.SResourceDriverBase = SResourceDriverBase{
		ResourceName:          eip.Name(),
		ResourceIsNeedDetails: eip.IsNeedDetails(),
	}
}

func init() {
	driver := &SOldEip{}
	driver.InitBase()
	ResourceRegister(driver)
}
