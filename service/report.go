package service

import (
	"health_checker/config"
	"health_checker/service/common"
)

func Report(nodeId, projectName, model string) error {
	fetchInfo := common.GetNvidiaFetchInfo()
	_ = fetchInfo
	err := common.Post(config.ConfigInfo.DBCHealthCheckReportUrl(), map[string]interface{}{
		"node_id":      nodeId,
		"project_name": projectName,
		"model":        model,
	})
	if err != nil {
		return err
	}

	return nil
}
