package apsarav2

import (
	"context"
	"net/http"

	"yunion.io/x/apsarav2/pkg/apsarav2/file"
	"yunion.io/x/jsonutils"
	"yunion.io/x/onecloud/pkg/appsrv"
	"yunion.io/x/pkg/errors"
)

func InitHandlers(app *appsrv.Application) {
	app.AddHandler("GET", "/<resources>", remotefile)
	app.AddHandler("POST", "/<resources>", remotefile)
}

var remotefile = func(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	params, _, _ := appsrv.FetchEnv(ctx, w, r)
	resource, isExist := params["<resources>"]
	if !isExist {
		appsrv.SendJSON(w, jsonutils.Marshal(map[string]string{
			"code": string(errors.ErrInvalidFormat),
			"msg":  "miss resources",
			"err":  "miss resources",
		}))
		return
	}
	obj, err := file.ReadRemoteFile(resource)
	if err != nil {
		appsrv.SendJSON(w, jsonutils.Marshal(map[string]string{
			"code": string(errors.ErrNotFound),
			"msg":  err.Error(),
			"err":  err.Error(),
		}))
		return
	}
	appsrv.SendJSON(w, obj)
}
