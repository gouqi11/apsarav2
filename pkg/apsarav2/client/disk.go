package client

import (
	"fmt"

	"yunion.io/x/apsarav2/pkg/apsarav2/file"
)

var DiskTypeMap = map[string]string{
	"system": "sys",
	"data":   "data",
}

type Attachment struct {
	InstanceId   string
	InstanceName string
}

type SOldDisk struct {
	SResourceDriverBase

	SCommonInfo

	Attachments          []Attachment
	AutoSnapshotPolicyId string
	Category             string
	CategoryName         string
	Detach               string
	Device               string
	DiskId               string
	DiskName             string
	DiskSize             int
	Encrypted            bool
	InstanceId           string
	Multiattach          string
	Portable             string
	RegionId             string
	Status               string
	Type                 string
	ZoneId               string
}

type SDisk struct {
	ResourceBaseInfo
	Id     string `json:"id"`
	Name   string `json:"name"`
	Status string `json:"status"`

	InstanceId   string
	ZoneId       string `json:"zone_id"`
	StorageId    string `json:"storage_id"`
	DiskFormat   string `json:"disk_format"`
	DiskSizeMb   int    `json:"disk_size_mb"`
	IsAutoDelete bool   `json:"is_auto_delete"`
	DiskType     string `json:"disk_type"`
	FsFormat     string `json:"fs_format"`
	Iops         int    `json:"iops"`
	Driver       string `json:"driver"`
	CacheMode    string `json:"cache_mode"`
	Mountpoint   string `json:"mountpoint"`
	AccessPath   string `json:"access_path"`
}

var diskStatusMap = map[string]string{
	"attached":  "ready",
	"available": "ready",
}

func getDiskDetails() ([]SDisk, error) {
	res := []SDisk{}
	old := []SOldDisk{}
	file.ReadFileDetail("disks", &old)
	for _, v := range old {
		temp := SDisk{
			ResourceBaseInfo: ResourceBaseInfo{
				Id:       v.DiskId,
				Name:     v.DiskName,
				Status:   diskStatusMap[v.Status],
				BaseInfo: getBaseInfo(v),
			},
			InstanceId: v.InstanceId,
			ZoneId:     v.ZoneId,
			StorageId:  fmt.Sprintf("%s-%s", v.Category, v.ZoneId),
			DiskFormat: "raw",
			DiskSizeMb: v.DiskSize * 1024,
			DiskType:   DiskTypeMap[v.Type],
		}
		base := getBaseInfo(v)
		temp.BaseInfo = base
		res = append(res, temp)
	}
	return res, nil
}

func (disk SOldDisk) SyncFromLocal() error {
	res := []SDisk{}
	old := []SOldDisk{}
	file.ReadFileDetail("disks", &old)
	for _, v := range old {
		temp := SDisk{
			ResourceBaseInfo: ResourceBaseInfo{
				Id:       v.DiskId,
				Name:     v.DiskName,
				Status:   diskStatusMap[v.Status],
				BaseInfo: getBaseInfo(v),
			},
			InstanceId: v.InstanceId,
			ZoneId:     v.ZoneId,
			StorageId:  fmt.Sprintf("%s-%s", v.Category, v.ZoneId),
			DiskFormat: "raw",
			DiskSizeMb: v.DiskSize * 1024,
			DiskType:   DiskTypeMap[v.Type],
		}
		base := getBaseInfo(v)
		temp.BaseInfo = base
		res = append(res, temp)
	}
	return file.WriteRemoteFile(disk.Name(), res)
}

func getDisksByInstanceId(disks []SDisk, instanceId string) []SDisk {
	res := []SDisk{}
	for i := 0; i < len(disks); i++ {
		if disks[i].InstanceId == instanceId {
			res = append(res, disks[i])
		}
	}
	return res
}

func (dbInstance SOldDisk) Name() string {
	return "disks"
}

func (dbInstance SOldDisk) IsNeedDetails() bool {
	return false
}

func (disk *SOldDisk) InitBase() {
	disk.SResourceDriverBase = SResourceDriverBase{
		ResourceName:          disk.Name(),
		ResourceIsNeedDetails: disk.IsNeedDetails(),
	}
}
func init() {
	driver := &SOldDisk{}
	driver.InitBase()
	ResourceRegister(driver)
}
