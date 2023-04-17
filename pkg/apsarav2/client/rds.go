package client

import (
	"fmt"
	"time"

	"yunion.io/x/apsarav2/pkg/apsarav2/file"
	"yunion.io/x/cloudmux/pkg/multicloud/remotefile"
)

type SRdsSku struct {
	Cpu int
	Mem int
}

var rdsSkuMap = map[string]SRdsSku{
	"60-240": {
		Cpu: 1,
		Mem: 2,
	},
	"150-600": {
		Cpu: 2,
		Mem: 4,
	},
	"300-1200": {
		Cpu: 4,
		Mem: 4,
	},
	"300-2400": {
		Cpu: 2,
		Mem: 8,
	},
	"500-600": {
		Cpu: 2,
		Mem: 8,
	},
	"600-2400": {
		Cpu: 2,
		Mem: 8,
	},
	"1500-6000": {
		Cpu: 4,
		Mem: 8,
	},
	"2000-12000": {
		Cpu: 4,
		Mem: 16,
	},
	"2000-24000": {
		Cpu: 8,
		Mem: 16,
	},
	"2000-48000": {
		Cpu: 8,
		Mem: 32,
	},
}

type SOldDbInstance struct {
	SResourceDriverBase

	IPAddress            string      `json:"ipAddress"`
	DBInstanceCreateTime string      `json:"dBInstanceCreateTime"`
	NetworkType          string      `json:"networkType"`
	DepartmentName       string      `json:"departmentName"`
	ProjectName          string      `json:"projectName"`
	CPU                  int         `json:"cpu"`
	RegionName           string      `json:"regionName"`
	RegionType           string      `json:"regionType"`
	DtDescription        interface{} `json:"dtDescription"`

	AccountMaxQuantity    int       `json:"accountMaxQuantity"`
	AvailabilityValue     string    `json:"availabilityValue"`
	BackupSize            int       `json:"backupSize"`
	ConnectionMode        string    `json:"connectionMode"`
	ConnectionString      string    `json:"connectionString"`
	CreationTime          time.Time `json:"creationTime"`
	DBInstanceClass       string    `json:"dBInstanceClass"`
	DBInstanceID          string    `json:"dBInstanceId"`
	DBInstanceMemory      int       `json:"dBInstanceMemory"`
	DBInstanceName        string    `json:"dBInstanceName"`
	DBInstanceStatus      string    `json:"dBInstanceStatus"`
	DBInstanceStorage     int       `json:"dBInstanceStorage"`
	DBInstanceType        string    `json:"dBInstanceType"`
	DBMaxQuantity         int       `json:"dBMaxQuantity"`
	DataDiskSize          int       `json:"dataDiskSize"`
	DepartmentID          string    `json:"departmentId"`
	DiskUsed              int       `json:"diskUsed"`
	Engine                string    `json:"engine"`
	EngineVersion         string    `json:"engineVersion"`
	ExpireDate            string    `json:"expireDate"`
	ExpireTime            time.Time `json:"expireTime"`
	InstanceNetworkType   string    `json:"instanceNetworkType"`
	LockMode              string    `json:"lockMode"`
	LockReason            string    `json:"lockReason"`
	LogDiskSize           int       `json:"logDiskSize"`
	MaxConnections        int       `json:"maxConnections"`
	MaxIOPS               int       `json:"maxIOPS"`
	Port                  int       `json:"port"`
	ProjectID             string    `json:"projectId"`
	ReadOnlyDBInstanceIds string    `json:"readOnlyDBInstanceIds"`
	RegionID              string    `json:"regionId"`
	RequestID             string    `json:"requestId"`
	TestTime              string    `json:"testTime"`
	VSwitchID             string    `json:"vSwitchId"`
	Vip                   string    `json:"vip"`
	Vpcid                 string    `json:"vpcid"`
	ZoneID                string    `json:"zoneId"`
}

func (dbInstance SOldDbInstance) SyncFromLocal() error {
	old := []SOldDbInstance{}
	file.ReadFileDetail(dbInstance.Name(), &old)
	res := []remotefile.SDBInstance{}
	for _, v := range old {
		temp := remotefile.SDBInstance{
			RegionId:      v.RegionID,
			Port:          v.Port,
			Engine:        v.Engine,
			EngineVersion: v.EngineVersion,
			InstanceType:  v.DBInstanceClass,
			VpcId:         v.Vpcid,
			// VcpuCount:             v.CPU,
			VcpuCount:  rdsSkuMap[fmt.Sprintf("%d-%d", v.MaxConnections, v.DBInstanceMemory)].Cpu,
			VmemSizeMb: rdsSkuMap[fmt.Sprintf("%d-%d", v.MaxConnections, v.DBInstanceMemory)].Cpu * 1024,
			// VmemSizeMb:            v.DBInstanceMemory,
			DiskSizeGb:            v.DBInstanceStorage,
			DiskSizeUsedGb:        v.DiskUsed,
			Category:              v.Engine,
			ConnectionStr:         v.ConnectionString,
			Zone1Id:               v.ZoneID,
			Iops:                  v.MaxConnections,
			StorageType:           "default",
			InternalConnectionStr: v.IPAddress,

			SResourceBase: remotefile.SResourceBase{
				Id:        v.DBInstanceID,
				Name:      v.DBInstanceName,
				Status:    "running",
				CreatedAt: v.CreationTime,
				ProjectId: v.ProjectID,
			},
		}

		res = append(res, temp)
	}
	return file.WriteRemoteFile(dbInstance.Name(), res)
}

func (dbInstance SOldDbInstance) Name() string {
	return "dbinstances"
}

func (dbInstance SOldDbInstance) IsNeedDetails() bool {
	return true
}

func (dbInstance *SOldDbInstance) InitBase() {
	dbInstance.SResourceDriverBase = SResourceDriverBase{
		ResourceName:          dbInstance.Name(),
		ResourceIsNeedDetails: dbInstance.IsNeedDetails(),
	}
}

func init() {
	driver := &SOldDbInstance{}
	driver.InitBase()
	ResourceRegister(driver)
}
