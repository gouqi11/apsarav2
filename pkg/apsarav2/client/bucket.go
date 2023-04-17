package client

import (
	"net/url"
	"strconv"

	"yunion.io/x/apsarav2/pkg/apsarav2/file"
	"yunion.io/x/cloudmux/pkg/cloudprovider"
	"yunion.io/x/cloudmux/pkg/multicloud/remotefile"
	"yunion.io/x/log"
	"yunion.io/x/pkg/errors"
	"yunion.io/x/pkg/util/httputils"
)

type SOldBucket struct {
	SResourceDriverBase

	CreateDate         string
	Creator            string
	ExpiredDate        string
	IsAliType          bool
	InstanceName       string
	RegionId           string
	StoreSpace         int    //容量
	StorageUtilization int64  `json:"storageUtilization"` // 存量
	BucketAcl          string `json:"buckertAcl"`         //

	SCommonInfo
}

type SBucket struct {
	ResourceBaseInfo

	RegionId     string
	MaxPart      int
	MaxPartBytes int64
	Acl          string
	Location     string
	StorageClass string

	Stats SBucketStats
	Limit SBucketStats
}

type SBucketStats struct {
	SizeBytes   int64
	ObjectCount int
}

func (bucket SOldBucket) SyncFromCloud() error {
	resObj, err := GetResourceList(resourceListActiionMap[bucket.Name()])
	if err != nil {
		return errors.Wrapf(err, "GetResource %s error:%v", resourceListActiionMap[bucket.Name()], err)
	}
	resOld := []SOldBucket{}
	resObj.Unmarshal(&resOld)
	for i := 0; i < len(resOld); i++ {
		query := url.Values{}
		query.Set("bucketName", resOld[i].InstanceName)
		query.Set("regionId", resOld[i].RegionId)
		aclObj, err := request("cloudOss/read/getBucketACL", httputils.GET, query)
		if err != nil {
			log.Errorln(errors.Wrap(err, "get bucketACL"))
			continue
		}
		acl, _ := aclObj.GetString("data", "bucketAcl")
		storageObj, err := request("monitoross/current/monitor", httputils.GET, query)
		if err != nil {
			log.Errorln(errors.Wrap(err, "get bucketACL"))
			continue
		}
		storageArr, _ := storageObj.GetArray("data", "rows")
		var stat int64
		for _, storage := range storageArr {
			statStr, _ := storage.GetString("storageUtilization")
			stat, _ = strconv.ParseInt(statStr, 10, 64)
		}
		resOld[i].BucketAcl = acl
		resOld[i].StorageUtilization = stat
	}
	return file.WriteFile(bucket.Name(), resOld)
}

func (bucket SOldBucket) SyncFromLocal() error {
	old := []SOldBucket{}
	err := file.ReadFileDetail("buckets", &old)
	if err != nil {
		return err
	}

	res := []remotefile.SBucket{}
	for _, v := range old {
		temp := remotefile.SBucket{
			SResourceBase: getCloudBaseInfo(v),
			RegionId:      v.RegionId,
			Acl:           v.BucketAcl,
			Limit: cloudprovider.SBucketStats{
				SizeBytes:   int64(v.StoreSpace) * 1024 * 1024 * 1024 * 1024,
				ObjectCount: 0,
			},
			Stats: cloudprovider.SBucketStats{
				SizeBytes:   v.StorageUtilization,
				ObjectCount: 0,
			},
		}
		temp.Id = v.InstanceName
		temp.Name = v.InstanceName
		res = append(res, temp)
	}
	return file.WriteRemoteFile(bucket.Name(), res)
}

func (bucket SOldBucket) Name() string {
	return "buckets"
}

func (bucket SOldBucket) IsNeedDetails() bool {
	return true
}

func (bucket *SOldBucket) InitBase() {
	bucket.SResourceDriverBase = SResourceDriverBase{
		ResourceName:          bucket.Name(),
		ResourceIsNeedDetails: bucket.IsNeedDetails(),
	}
}

func init() {
	driver := &SOldBucket{}
	ResourceRegister(driver)
}
