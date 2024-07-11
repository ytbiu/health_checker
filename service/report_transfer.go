package service

import (
	"health_checker/config"
	"health_checker/service/common"
)

type RegisterModel struct {
	Model string
}

func Report(nodeId, project string, models []RegisterModel) error {
	fetchInfo := common.GetNvidiaFetchInfo()
	return common.Post(config.ConfigInfo.DBCHealthCheckReportUrl(), map[string]interface{}{
		"node_id":        nodeId,
		"project":        project,
		"models":         models,
		"GPUName":        fetchInfo.GPUName,
		"UtilizationGPU": fetchInfo.UtilizationGPU,
		"MemoryTotal":    fetchInfo.MemoryTotal,
		"MemoryUsed":     fetchInfo.MemoryUsed,
	})
}
