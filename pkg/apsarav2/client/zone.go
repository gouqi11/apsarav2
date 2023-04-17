package client

import (
	"net/url"

	"yunion.io/x/apsarav2/pkg/apsarav2/file"
	"yunion.io/x/jsonutils"
	"yunion.io/x/pkg/errors"
	"yunion.io/x/pkg/util/httputils"
)

type SZone struct {
	Id       string
	Name     string
	RegionId string
	Status   string
}

func getZones(regionId string) (jsonutils.JSONObject, error) {
	query := url.Values{}
	query.Set("regionId", regionId)
	return request("region/zones", httputils.GET, query)
}

func getZoneDetails() ([]SZone, error) {
	zonesObj, err := file.ReadFile("zones")
	if err != nil {
		return nil, errors.Wrap(err, "ReadFile")
	}
	temp := map[string]jsonutils.JSONObject{}
	zonesObj.Unmarshal(&temp)
	zones := []SZone{}
	for regionId, v := range temp {
		objArr, _ := v.GetArray()
		for _, obj := range objArr {
			zoneId, _ := obj.GetString("zoneId")
			zoneName, _ := obj.GetString("zoneName")
			zones = append(zones, SZone{Id: zoneId, Name: zoneName, RegionId: regionId, Status: "available"})
		}
	}
	return zones, nil
}
