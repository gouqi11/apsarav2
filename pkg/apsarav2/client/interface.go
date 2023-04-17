package client

import (
	"net/url"

	"github.com/pkg/errors"
	"yunion.io/x/apsarav2/pkg/apsarav2/file"
	"yunion.io/x/jsonutils"
	"yunion.io/x/log"
	"yunion.io/x/pkg/util/httputils"
)

var (
	resourceDriverTable = make(map[string]IResourceDriver)
)

type IResourceDriver interface {
	Name() string
	IsNeedDetails() bool
	SyncFromCloud() error
	SyncFromLocal() error
	InitBase()
}

func ResourceRegister(driver IResourceDriver) {
	resourceDriverTable[driver.Name()] = driver
	driver.InitBase()
}

func GetResourceDriver(resourceDriver string) IResourceDriver {
	driver, isExist := resourceDriverTable[resourceDriver]
	if isExist {
		return driver
	}
	return nil
}

type SResourceDriverBase struct {
	ResourceName          string
	ResourceIsNeedDetails bool
}

func (base SResourceDriverBase) Name() string {
	return ""
}

func (base SResourceDriverBase) SyncFromCloud() error {
	resObj, err := GetResourceList(resourceListActiionMap[base.ResourceName])
	if err != nil {
		return errors.Wrapf(err, "GetResource %s error:%v", resourceListActiionMap[base.ResourceName], err)
	}
	if base.ResourceIsNeedDetails {
		resOld, _ := resObj.GetArray()
		// 需要获取详情的资源
		for i := 0; i < len(resOld); i++ {
			regionId, _ := resOld[i].GetString("regionId")
			departmentId, _ := resOld[i].GetString("departmentId")
			respIdKey := resourceRespIdMap[base.ResourceName]
			reqIdKey := resourceReqIdMap[base.ResourceName]
			idValue, _ := jsonutils.Marshal(resOld[i]).GetString(respIdKey)
			query := url.Values{}
			query.Set("regionId", regionId)
			query.Set(reqIdKey, idValue)
			query.Set("departmentId", departmentId)
			detail, err := request(resourceDetailActiionMap[base.ResourceName], httputils.GET, query)
			if err != nil {
				return errors.Wrap(err, "get detail error:")
			}
			detailData, _ := detail.Get("data")
			err = jsonutils.Update(&resOld[i], detailData)
			if err != nil {
				log.Errorln("this is update err:", err)
			}
		}
		err = file.WriteFile(base.ResourceName, resOld)
	} else {
		err = file.WriteFile(base.ResourceName, resObj)
	}
	return err
}
