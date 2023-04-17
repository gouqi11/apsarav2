package client

import (
	"time"

	"yunion.io/x/apsarav2/pkg/apsarav2/file"
	"yunion.io/x/pkg/errors"
)

type SOldEcsMetric struct {
	InternetRx        int    `json:"internetRx"`
	BpsWrite          int    `json:"bpsWrite"`
	RegionName        string `json:"regionName"`
	InstanceID        string `json:"instanceId"`
	InstanceStatus    string `json:"instanceStatus"`
	ProjectID         string `json:"projectId"`
	InternetTx        int    `json:"internetTx"`
	DepartmentName    string `json:"departmentName"`
	DiskUtilization   int    `json:"diskUtilization"`
	RegionType        string `json:"regionType"`
	CPUUtilization    int    `json:"cpuUtilization"`
	BpsRead           int    `json:"bpsRead"`
	CollectTime       string `json:"collectTime"`
	InstanceName      string `json:"instanceName"`
	MemoryUtilization int    `json:"memoryUtilization"`
	DepartmentID      string `json:"departmentId"`
	ProjectName       string `json:"projectName"`
	AgentStatus       string `json:"agentStatus"`
	RegionID          string `json:"regionId"`
}

type SOldRdsMetric struct {
	CPUUtilization    int     `json:"cpuUtilization"`
	DepartmentID      string  `json:"departmentId"`
	DepartmentName    string  `json:"departmentName"`
	DatabaseType      string  `json:"databaseType"`
	InstanceID        string  `json:"instanceId"`
	InstanceName      string  `json:"instanceName"`
	IopsUtilization   int     `json:"iopsUtilization"`
	MemoryUtilization float64 `json:"memoryUtilization"`
	NetTraffic        int     `json:"netTraffic"`
	ProjectID         string  `json:"projectId"`
	ProjectName       string  `json:"projectName"`
	RegionID          string  `json:"regionId"`
	RegionName        string  `json:"regionName"`
	RegionType        string  `json:"regionType"`
	Sessions          int     `json:"sessions"`
	SpaceUsage        float64 `json:"spaceUsage"`
	InstanceStatus    string  `json:"instanceStatus"`
}

type SMetric struct {
	MetricType string          `json:"metric_type"`
	Id         string          `json:"id"`
	Values     []SMetricValues `json:"values"`
}

type SMetricValues struct {
	Timestamp time.Time
	Value     float64
}

func SyncMetricFromCloud() error {
	ecsObj, err := GetResourceList("monitorecs/current/monitor")
	if err != nil {
		return err
	}
	err = file.WriteFile("ecsMetric", ecsObj)
	if err != nil {
		return errors.Wrap(err, "get ecsMetric")
	}
	rdsObj, err := GetResourceList("monitorrds/current/monitor")
	if err != nil {
		return err
	}
	err = file.WriteFile("rdsMetric", rdsObj)
	if err != nil {
		return errors.Wrap(err, "get ecsMetric")
	}
	return nil
}

func SyncMetricFromLocal() error {
	obj, err := file.ReadFile("ecsMetric")
	if err != nil {
		return err
	}
	old := []SOldEcsMetric{}
	obj.Unmarshal(&old)
	res := []map[string]SMetric{}
	for _, v := range old {

		res = append(res, map[string]SMetric{
			"vm_cpu.usage_active": {
				MetricType: "vm_cpu.usage_active",
				Values: []SMetricValues{
					{
						Timestamp: time.Now(),
						Value:     float64(v.CPUUtilization),
					},
				},
				Id: v.InstanceID,
			},
		})
		res = append(res, map[string]SMetric{
			"vm_mem.used_percent": {
				MetricType: "vm_mem.used_percent",
				Values: []SMetricValues{
					{
						Timestamp: time.Now(),
						Value:     float64(v.MemoryUtilization),
					},
				},
				Id: v.InstanceID,
			},
		})
		res = append(res, map[string]SMetric{
			"vm_disk.used_percent": {
				MetricType: "vm_disk.used_percent",
				Values: []SMetricValues{
					{
						Timestamp: time.Now(),
						Value:     float64(v.DiskUtilization),
					},
				},
				Id: v.InstanceID,
			}})
	}

	obj, err = file.ReadFile("rdsMetric")
	if err != nil {
		return err
	}
	oldRds := []SOldRdsMetric{}
	obj.Unmarshal(&oldRds)
	for _, v := range oldRds {
		res = append(res, map[string]SMetric{
			"rds_cpu.usage_active": {
				MetricType: "rds_cpu.usage_active",
				Values: []SMetricValues{
					{
						Timestamp: time.Now(),
						Value:     float64(v.CPUUtilization),
					},
				},
				Id: v.InstanceID,
			},
		})
		res = append(res, map[string]SMetric{
			"rds_mem.used_percent": {
				MetricType: "rds_mem.used_percent",
				Values: []SMetricValues{
					{
						Timestamp: time.Now(),
						Value:     float64(v.MemoryUtilization),
					},
				},
				Id: v.InstanceID,
			},
		})
		res = append(res, map[string]SMetric{
			"rds_disk.used_percent": {
				MetricType: "rds_disk.used_percent",
				Values: []SMetricValues{
					{
						Timestamp: time.Now(),
						Value:     float64(v.SpaceUsage),
					},
				},
				Id: v.InstanceID,
			}})
	}
	return file.WriteRemoteFile("metrics", res)
}
