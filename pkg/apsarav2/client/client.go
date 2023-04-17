package client

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/pkg/errors"
	"yunion.io/x/apsarav2/pkg/apsarav2/options"
	"yunion.io/x/jsonutils"
	"yunion.io/x/pkg/util/httputils"
)

func request(action string, method httputils.THttpMethod, query url.Values) (jsonutils.JSONObject, error) {
	url := fmt.Sprintf("%s/%s", options.Options.Url, action)
	if len(query) > 0 {
		url = fmt.Sprintf("%s?%s", url, query.Encode())
	}
	header := setHeader(method)
	client := httputils.GetTimeoutClient(time.Minute * 2)
	_, resp, err := httputils.JSONRequest(client, context.Background(), method, url, header, nil, options.Options.DebugRequest)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func setHeader(method httputils.THttpMethod) http.Header {
	params := url.Values{}
	params.Set("SignatureMethod", "HMAC-SHA256")
	params.Set("Version", "20160701")
	params.Set("AccessKeyId", options.Options.AccessKeyId)
	var sign = func(method string, opts url.Values, secret string) string {
		stringToSign := strings.Replace(opts.Encode(), "+", "%20", -1)
		stringToSign = strings.Replace(stringToSign, "*", "%2A", -1)
		stringToSign = strings.Replace(stringToSign, "%7E", "~", -1)
		stringToSign = url.QueryEscape(stringToSign)
		signStr := method + "&%2F&" + stringToSign
		h := hmac.New(sha256.New, []byte(options.Options.AccessSecret+"&"))
		h.Write([]byte(signStr))
		return base64.StdEncoding.EncodeToString((h.Sum(nil)))
	}
	key := sign(string(method), params, options.Options.AccessSecret)
	header := http.Header{}
	header.Set("AccessKeyId", options.Options.AccessKeyId)
	header.Set("Signature", key)
	header.Set("SignatureMethod", "HMAC-SHA256")
	header.Set("Version", "20160701")
	return header
}

func GetResourceList(action string) (*jsonutils.JSONArray, error) {
	resArr := new(jsonutils.JSONArray)
	pageSize := 50
	pageNum := 1
	query := url.Values{}
	query.Set("pageSize", fmt.Sprintf("%d", pageSize))
	for {
		query.Set("pageNum", fmt.Sprintf("%d", pageNum))
		res, err := request(action, httputils.GET, query)
		if err != nil {
			return nil, errors.Wrapf(err, "get %s resource", action)
		}
		total, _ := res.GetString("data", "totalRows")
		totalInt, _ := strconv.Atoi(total)
		resTemp, _ := res.GetArray("data", "rows")
		resArr.Add(resTemp...)
		if resArr.Length() >= totalInt {
			break
		}
		pageNum++
	}
	return resArr, nil
}

func GetVpcList(action string) (*jsonutils.JSONArray, error) {
	resArr := new(jsonutils.JSONArray)
	query := url.Values{}
	res, err := request(action, httputils.GET, query)
	if err != nil {
		return nil, errors.Wrapf(err, "get %s resource", action)
	}
	resTemp, _ := res.GetArray("data", "result")
	resArr.Add(resTemp...)
	return resArr, nil
}
