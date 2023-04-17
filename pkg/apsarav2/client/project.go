package client

import (
	"yunion.io/x/apsarav2/pkg/apsarav2/file"
	"yunion.io/x/log"
)

type SProjects struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Status      string `json:"status" default:"available"`

	BaseInfo
}

func getProjectDetails() ([]SProjects, error) {
	projects := []SProjects{}
	idMap := make(map[string]struct{})
	for resource := range resourceDriverTable {
		obj, err := file.ReadFile(resource)
		if err != nil {
			return nil, err
		}
		objArr, _ := obj.GetArray()
		if len(objArr) == 0 {
			log.Infof("this is resource:%s,this is obj:%v", resource, obj)
		}
		for _, value := range objArr {
			projectId, _ := value.GetString("projectId")
			projectName, _ := value.GetString("projectName")
			if len(projectId) == 0 || len(projectName) == 0 {
				continue
			}
			if _, isExist := idMap[projectId]; isExist {
				continue
			}
			idMap[projectId] = struct{}{}
			projects = append(projects, SProjects{
				Id:     projectId,
				Name:   projectName,
				Status: "available",
				BaseInfo: BaseInfo{
					Tags:    map[string]string{},
					SysTags: map[string]string{},
				},
			})
		}
	}
	return projects, nil
}
