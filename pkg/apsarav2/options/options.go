package options

import (
	common_options "yunion.io/x/onecloud/pkg/cloudcommon/options"
)

type ApiServiceOptions struct {
	DebugRequest bool `default:"false"`

	Url          string `default:"http://192.168.86.241/gateway/api"`
	Method       string `default:"GET" choices:"DELETE|PUT|GET|POST"`
	AccessKeyId  string `default:"F9OEFyBXR6MddLnw"`
	AccessSecret string `default:"sb97t9KPRNiWlaEK4DHAnQHjTntb03"`

	RemoteFileLocation string `default:"/opt/yunion/apsaraJson"`
	FileLocation       string `default:"/opt/yunion/apsaraJson/original"`

	ClientInit     bool `default:"true"`
	EdgeClientInit bool `default:"true"`

	common_options.DBOptions
	common_options.CommonOptions
}

var (
	Options ApiServiceOptions
)
