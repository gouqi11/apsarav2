package client

import (
	"github.com/pkg/errors"
	"yunion.io/x/apsarav2/pkg/apsarav2/file"
)

type SOldVpc struct {
	SResourceDriverBase

	CidrBlock       string
	DefaultSGId     string `json:"defaultSGId"`
	Description     string
	ExtNetworkId    string
	ExtNetworkName  string
	ExternalGateway string
	NgwId           string
	RegionId        string
	RegionName      string
	RegionType      string
	ReservedType    string
	Status          string
	VRouterId       string
	VSwitchCount    int
	VpcId           string
	VpcName         string

	SCommonInfo
}

type SVpc struct {
	ResourceBaseInfo

	RegionId  string
	CidrBlock string
	IsDefault bool
}

func (vpc SOldVpc) SyncFromCloud() error {
	res, err := GetVpcList(resourceListActiionMap[vpc.Name()])
	if err != nil {
		return errors.Wrapf(err, "GetResource %s ", resourceListActiionMap[vpc.Name()])
	}
	return file.WriteFile(vpc.Name(), res)
}

func (vpc SOldVpc) SyncFromLocal() error {
	old := []SOldVpc{}
	err := file.ReadFileDetail("vpcs", &old)
	if err != nil {
		return err
	}
	res := []SVpc{}
	for _, v := range old {
		temp := SVpc{
			ResourceBaseInfo: ResourceBaseInfo{
				Id:       v.VpcId,
				Name:     v.VpcName,
				Status:   "available",
				BaseInfo: getBaseInfo(v),
			},
			RegionId:  v.RegionId,
			CidrBlock: v.CidrBlock,
		}
		res = append(res, temp)
	}
	return file.WriteRemoteFile(vpc.Name(), res)
}

func (vpc SOldVpc) Name() string {
	return "vpcs"
}

func (vpc SOldVpc) IsNeedDetails() bool {
	return false
}

func (vpc *SOldVpc) InitBase() {
	vpc.SResourceDriverBase = SResourceDriverBase{
		ResourceName:          vpc.Name(),
		ResourceIsNeedDetails: vpc.IsNeedDetails(),
	}
}

func init() {
	driver := &SOldVpc{}
	ResourceRegister(driver)
}
