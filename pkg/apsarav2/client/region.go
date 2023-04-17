package client

import (
	"yunion.io/x/jsonutils"
	"yunion.io/x/pkg/util/httputils"
)

type SRegion struct {
	Id   string
	Name string
}

func getRegions() (jsonutils.JSONObject, error) {
	return request("region/list", httputils.GET, nil)
}

func getRegionDetails(res jsonutils.JSONObject) ([]SRegion, []string, error) {
	regions := []SRegion{}
	regionIds := []string{}
	objArr, err := res.GetArray("data")
	if err != nil {
		return nil, nil, err
	}
	for _, obj := range objArr {
		regionId, _ := obj.GetString("regionId")
		regionName, _ := obj.GetString("regionName")
		regionIds = append(regionIds, regionId)
		regions = append(regions, SRegion{Id: regionId, Name: regionName})
	}
	return regions, regionIds, nil
}
