package client

import (
	"context"

	"yunion.io/x/apsarav2/pkg/apsarav2/file"
	"yunion.io/x/jsonutils"
	"yunion.io/x/log"
	"yunion.io/x/onecloud/pkg/mcclient"
)

func SyncFromCloud(ctx context.Context, userCred mcclient.TokenCredential, isStart bool) {
	regionRes, err := getRegions()
	if err != nil {
		log.Errorln("this is err:", err)
		return
	}
	file.WriteFile("regions", regionRes)

	_, regionIds, _ := getRegionDetails(regionRes)
	zoneRes := map[string]jsonutils.JSONObject{}
	for _, regionId := range regionIds {
		zone, err := getZones(regionId)
		if err != nil {
			log.Errorln("this is err:", err)
			return
		}
		ZoneData, _ := zone.Get("data")
		zoneRes[regionId] = ZoneData
	}
	file.WriteFile("zones", zoneRes)

	for resource, driver := range resourceDriverTable {
		go func(resource string, driver IResourceDriver) {
			err := driver.SyncFromCloud()
			if err != nil {
				log.Errorf("this is resource:%s,err:%v", resource, err)
				return
			}
		}(resource, driver)
	}
}

func CronSyncMetricFromCloud(ctx context.Context, userCred mcclient.TokenCredential, isStart bool) {
	err := SyncMetricFromCloud()
	if err != nil {
		log.Errorln("SyncMetricFromCloud:", err)
		return
	}
	err = SyncMetricFromLocal()
	if err != nil {
		log.Errorln("SyncMetricFromCloud:", err)
	}
}
