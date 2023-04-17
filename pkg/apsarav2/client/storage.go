package client

import (
	"fmt"

	"yunion.io/x/apsarav2/pkg/apsarav2/file"
	"yunion.io/x/log"
)

type Storages struct {
	Id          string `json:"id"`
	ZoneId      string `json:"zone_id"`
	StorageType string `json:"storage_type"`
}

func getStorageDetails() ([]Storages, error) {
	disks := []SOldDisk{}
	err := file.ReadFileDetail("disks", &disks)
	if err != nil {
		log.Errorln("readFileDetai:", err)
		return nil, err
	}
	storageMap := make(map[string]struct{})
	res := []Storages{}
	for _, v := range disks {
		temp := Storages{
			Id:          fmt.Sprintf("%s-%s", v.Category, v.ZoneId),
			ZoneId:      v.ZoneId,
			StorageType: v.Category,
		}
		if _, isExist := storageMap[temp.Id]; isExist || len(v.Category) == 0 || len(v.ZoneId) == 0 {
			continue
		}
		storageMap[temp.Id] = struct{}{}
		res = append(res, temp)
	}
	return res, nil
}
