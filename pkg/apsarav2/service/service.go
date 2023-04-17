package service

import (
	"os"
	"time"

	"yunion.io/x/log"

	service "yunion.io/x/apsarav2/pkg/apsarav2"
	"yunion.io/x/onecloud/pkg/cloudcommon"
	app_common "yunion.io/x/onecloud/pkg/cloudcommon/app"
	"yunion.io/x/onecloud/pkg/cloudcommon/cronman"
	"yunion.io/x/onecloud/pkg/cloudcommon/db"
	common_options "yunion.io/x/onecloud/pkg/cloudcommon/options"
	"yunion.io/x/onecloud/pkg/image/models"

	"yunion.io/x/apsarav2/pkg/apsarav2/client"
	"yunion.io/x/apsarav2/pkg/apsarav2/options"
	_ "yunion.io/x/apsarav2/pkg/apsarav2/policy"
	_ "yunion.io/x/sqlchemy/backends"
)

func StartService() {
	opts := &options.Options
	baseOpts := &options.Options.BaseOptions
	commonOpts := &options.Options.CommonOptions
	dbOpts := &options.Options.DBOptions
	common_options.ParseOptions(opts, os.Args, "api.conf", "wz-api")

	if !db.CheckSync(opts.AutoSyncTable, false, true) {
		log.Fatalf("database schema not in sync!")
	}

	app_common.InitAuth(commonOpts, func() {
		log.Infof("Auth complete!!")
	})

	client.InitResourceMap()

	//consts.SetDataResp(true)
	cloudcommon.InitDB(&opts.DBOptions)
	app := app_common.InitApp(baseOpts, false)
	service.InitHandlers(app)
	defer cloudcommon.CloseDB()
	db.EnsureAppSyncDB(app, dbOpts, models.InitDB)

	cron := cronman.InitCronJobManager(true, 5)
	cron.AddJobAtIntervalsWithStartRun("SyncApsarav2FromCloud", time.Duration(3)*time.Hour, client.SyncFromCloud, true)
	cron.AddJobAtIntervalsWithStartRun("SyncApsarav2FromLocal", time.Duration(3)*time.Hour, client.SyncFromLocal, true)
	cron.AddJobAtIntervalsWithStartRun("SyncApsarav2Metric", time.Duration(5)*time.Minute, client.CronSyncMetricFromCloud, false)
	go cron.Start()
	app_common.ServeForever(app, baseOpts)
}
