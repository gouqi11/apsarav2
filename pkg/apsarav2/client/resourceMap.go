package client

var resourceListActiionMap map[string]string
var resourceDetailActiionMap map[string]string
var resourceRespIdMap map[string]string
var resourceReqIdMap map[string]string

func InitResourceMap() {
	resourceListActiionMap = map[string]string{
		// "hosts": "cloudLoadBalancer/read/getBackendServerList",
		"instances":     "cloudEcs/read/getCloudEcsList",
		"dbinstances":   "rds/instance/getRdsInstanceList",
		"disks":         "cloudDisk/read/getCloudDiskList",
		"loadbalancers": "cloudLoadBalancer/read/getCloudSlbList",
		"buckets":       "cloudOss/read/getBucketList",
		"secgroups":     "secgroups/read/getSecurityGroupList",
		"vpcs":          "vpcs/read/getVpcs",
		"eips":          "eips/getEipList",
		// "networks":      "network/getNetworksByVpcId",
	}

	resourceDetailActiionMap = map[string]string{
		"instances":     "cloudEcs/read/getEcsDetailByInstanceId",              //regionId,instanceId
		"dbinstances":   "rds/instance/getRdsInstanceAttribute",                //regionId,instanceId
		"disks":         "cloudDisk/read/getCloudDiskInfo",                     //regionId,diskId
		"loadbalancers": "cloudLoadBalancer/read/getCloudSlbEntryBaseInfoById", //regionId,loadBalancerId
		"buckets":       "cloudOss/read/getBucketInfo",                         //regionId,buckerName
		"vpcs":          "vpcs/read/",                                          //vpcId
	}

	resourceRespIdMap = map[string]string{
		"instances":     "instanceId",     //regionId,instanceId
		"dbinstances":   "dBInstanceId",   //regionId,instanceId
		"disks":         "diskId",         //regionId,diskId
		"loadbalancers": "loadBalancerId", //regionId,loadBalancerId
		"buckets":       "instanceName",   //regionId,buckerName
		"vpcs":          "vpcId",          //vpcId
		// "eips":        "eips/getEipList",
	}

	resourceReqIdMap = map[string]string{
		"instances":     "instanceId",     //regionId,instanceId
		"dbinstances":   "instanceId",     //regionId,instanceId
		"disks":         "diskId",         //regionId,diskId
		"loadbalancers": "loadBalancerId", //regionId,loadBalancerId
		"buckets":       "bucketName",     //regionId,bucketName
		"vpcs":          "vpcId",          //vpcId
		// "eips":        "eips/getEipList",
	}

}
