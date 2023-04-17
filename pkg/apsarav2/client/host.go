package client

import (
	"fmt"

	"github.com/pkg/errors"
	"yunion.io/x/apsarav2/pkg/apsarav2/file"
	"yunion.io/x/cloudmux/pkg/multicloud/remotefile"
)

type SHost struct {
	ZoneId        string `json:"zone_id"`
	Enabled       bool   `json:"enabled"`
	HostStatus    string `json:"host_status"`
	AccessIp      string `json:"access_ip"`
	AccessMac     string `json:"access_mac"`
	Sn            string `json:"sn"`
	CpuCount      int    `json:"cpu_count"`
	NodeCount     int    `json:"node_count"`
	CpuDesc       string `json:"cpu_desc"`
	CpuMbz        int    `json:"cpu_mbz"`
	MemSizeMb     int    `json:"mem_size_mb"`
	StorageSizeMb int    `json:"storage_size_mb"`

	Wires []remotefile.SWire
	ResourceBaseInfo
}

func getHostDetails() ([]remotefile.SHost, error) {
	zones, err := getZoneDetails()
	if err != nil {
		return nil, errors.Wrap(err, "getZoneDetails")
	}
	res := []remotefile.SHost{}

	obj, err := file.ReadRemoteFile("wires.json")
	if err != nil {
		return nil, err
	}
	wires := []remotefile.SWire{}
	obj.Unmarshal(&wires)
	for _, zone := range zones {
		temp := remotefile.SHost{
			SResourceBase: remotefile.SResourceBase{
				Id:       "host-1",
				Name:     fmt.Sprintf("instance-%s", zone.Id),
				Emulated: true,
			},
			Wires:  wires,
			ZoneId: zone.Id,
		}
		res = append(res, temp)
	}
	return res, nil
}
