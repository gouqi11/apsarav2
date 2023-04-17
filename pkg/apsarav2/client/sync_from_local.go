package client

import (
	"context"

	"github.com/pkg/errors"
	"yunion.io/x/apsarav2/pkg/apsarav2/file"
	"yunion.io/x/log"
	"yunion.io/x/onecloud/pkg/mcclient"
)

func SyncFromLocal(ctx context.Context, userCred mcclient.TokenCredential, isStart bool) {
	err := func() error {
		// regions
		regionsObj, err := file.ReadFile("regions")
		if err != nil {
			return errors.Wrap(err, "ReadFile")
		}
		regions, _, err := getRegionDetails(regionsObj)
		if err != nil {
			return errors.Wrap(err, "getRegionDetails")
		}
		file.WriteRemoteFile("regions", regions)

		// zones
		zones, err := getZoneDetails()
		if err != nil {
			return errors.Wrap(err, "getZoneDetails")
		}
		file.WriteRemoteFile("zones", zones)

		// projects
		project, err := getProjectDetails()
		if err != nil {
			return errors.Wrap(err, "getProjectDetails")
		}
		file.WriteRemoteFile("projects", project)

		// storages
		storages, err := getStorageDetails()
		if err != nil {
			return errors.Wrap(err, "getStorageDetails")
		}
		file.WriteRemoteFile("storages", storages)

		for resource, driver := range resourceDriverTable {
			// go func(resource string, driver IResourceDriver) {
			err := driver.SyncFromLocal()
			if err != nil {
				log.Errorln(errors.Wrapf(err, "resource:%s,SyncFromLocal", resource))
			}
			// }(resource, driver)
		}

		// wires
		wires, err := getWireDetails()
		if err != nil {
			return errors.Wrap(err, "getWireDetails")
		}
		file.WriteRemoteFile("wires", wires)

		// hosts
		hosts, err := getHostDetails()
		if err != nil {
			return errors.Wrap(err, "getHostDetails")
		}
		file.WriteRemoteFile("hosts", hosts)
		// networks
		networks, err := getNetworkDetails()
		if err != nil {
			return errors.Wrap(err, "getNetworkDetails")
		}
		file.WriteRemoteFile("networks", networks)
		return nil
	}()
	if err != nil {
		log.Errorln("SyncExternal", err)
	}
}
