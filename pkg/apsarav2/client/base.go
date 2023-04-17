package client

import (
	"time"

	"yunion.io/x/cloudmux/pkg/multicloud/remotefile"
	"yunion.io/x/jsonutils"
	"yunion.io/x/log"
)

type SCommonInfo struct {
	CreatedAt    string
	CreateDate   string
	CreationTime string
	// Slb
	CreateTime string
	// Eip
	AllocationTime string
	// Disk
	StartTime      string
	Department     string
	DepartmentId   string
	DepartmentName string
	ProjectId      string
	ProjectName    string
	Description    string
}

type IResourceDriverBase interface {
}

type BaseInfo struct {
	ProjectId string
	Emulated  bool              `json:"emulated"`
	CreatedAt string            `json:"created_at"`
	Tags      map[string]string `json:"tags"`
	SysTags   map[string]string `json:"sys_tags"`
}

type ResourceBaseInfo struct {
	Id     string `json:"id"`
	Name   string `json:"name"`
	Status string

	BaseInfo
}

func getBaseInfo(val interface{}) BaseInfo {
	obj := jsonutils.Marshal(val)

	commonInfo := SCommonInfo{}
	err := obj.Unmarshal(&commonInfo)
	if err != nil {
		log.Errorln("this is unmarshal error:", err)
	}
	res := BaseInfo{}
	if len(commonInfo.StartTime) > 0 {
		res.CreatedAt = commonInfo.StartTime
	}
	if len(commonInfo.CreateTime) > 0 {
		res.CreatedAt = commonInfo.CreateTime
	}
	if len(commonInfo.CreatedAt) > 0 {
		res.CreatedAt = commonInfo.CreatedAt
	}
	if len(commonInfo.CreateDate) > 0 {
		res.CreatedAt = commonInfo.CreateDate
	}
	if len(commonInfo.CreationTime) > 0 {
		res.CreatedAt = commonInfo.CreationTime
	}
	if len(commonInfo.AllocationTime) > 0 {
		res.CreatedAt = commonInfo.AllocationTime
	}

	res.ProjectId = commonInfo.ProjectId
	res.Tags = getCommonTag(commonInfo)
	res.SysTags = make(map[string]string)
	return res
}

func getCloudBaseInfo(val interface{}) remotefile.SResourceBase {
	obj := jsonutils.Marshal(val)

	commonInfo := SCommonInfo{}
	err := obj.Unmarshal(&commonInfo)
	if err != nil {
		log.Errorln("this is unmarshal error:", err)
	}

	res := remotefile.SResourceBase{}
	if len(commonInfo.StartTime) > 0 {
		res.CreatedAt, _ = time.Parse("2006-01-02 15:04:05", commonInfo.StartTime)
	}
	if len(commonInfo.CreateTime) > 0 {
		res.CreatedAt, _ = time.Parse("2006-01-02 15:04:05", commonInfo.CreateTime)
	}
	if len(commonInfo.CreatedAt) > 0 {
		res.CreatedAt, _ = time.Parse("2006-01-02 15:04:05", commonInfo.CreatedAt)
	}
	if len(commonInfo.CreateDate) > 0 {
		res.CreatedAt, _ = time.Parse("2006-01-02 15:04:05", commonInfo.CreateDate)
	}
	if len(commonInfo.CreationTime) > 0 {
		res.CreatedAt, _ = time.Parse("2006-01-02 15:04:05", commonInfo.CreationTime)
	}
	if len(commonInfo.AllocationTime) > 0 {
		res.CreatedAt, _ = time.Parse("2006-01-02 15:04:05", commonInfo.AllocationTime)
	}
	res.Tags = getCommonTag(commonInfo)
	res.SysTags = make(map[string]string)
	res.ProjectId = commonInfo.ProjectId
	return res
}

func getCommonTag(common SCommonInfo) map[string]string {
	tags := make(map[string]string)
	if len(common.Department) > 0 {
		tags["departmentId"] = common.Department
	}
	if len(common.DepartmentId) > 0 {
		tags["departmentId"] = common.DepartmentId
	}
	if len(common.DepartmentName) > 0 {
		tags["departmentName"] = common.DepartmentName
	}
	if len(common.ProjectId) > 0 {
		tags["projectId"] = common.ProjectId
	}
	if len(common.ProjectName) > 0 {
		tags["projectName"] = common.ProjectName
	}
	return tags
}
