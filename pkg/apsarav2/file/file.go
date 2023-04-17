package file

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"yunion.io/x/apsarav2/pkg/apsarav2/options"
	"yunion.io/x/jsonutils"
	"yunion.io/x/log"
)

func initDir(res string) {
	dir := fmt.Sprintf("%s/%s", options.Options.FileLocation, res)
	err := os.Mkdir(dir, 0755)
	if err != nil && !strings.Contains(err.Error(), "file exists") {
		log.Errorln("err:", err)
	}
}

func initRegionDir(res string) {
	dir := fmt.Sprintf("%s/%s", options.Options.FileLocation, res)
	err := os.Mkdir(dir, 0755)
	if err != nil && !strings.Contains(err.Error(), "file exists") {
		log.Errorln("err:", err)
	}
}

func ReadFile(resource string) (jsonutils.JSONObject, error) {
	filename := fmt.Sprintf("%s/%s.json", options.Options.FileLocation, resource)
	obj, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return jsonutils.Parse(obj)
}

func ReadFileDetail(resource string, retVal interface{}) error {
	filename := fmt.Sprintf("%s/%s.json", options.Options.FileLocation, resource)
	obj, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	jsonObj, err := jsonutils.Parse(obj)
	if err != nil {
		return err
	}
	return jsonObj.Unmarshal(retVal)
}

func WriteFile(res string, data interface{}) error {
	filename := fmt.Sprintf("%s/%s.json", options.Options.FileLocation, res)
	return ioutil.WriteFile(filename, []byte(jsonutils.Marshal(data).PrettyString()), 0755)
}

func WriteRemoteFile(res string, data interface{}) error {
	filename := fmt.Sprintf("%s/%s.json", options.Options.RemoteFileLocation, res)
	return ioutil.WriteFile(filename, []byte(jsonutils.Marshal(data).PrettyString()), 0755)
}

func ReadRemoteFile(resource string) (jsonutils.JSONObject, error) {
	filename := fmt.Sprintf("%s/%s", options.Options.RemoteFileLocation, resource)
	obj, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return jsonutils.Parse(obj)
}
