package service

import (
	"github.com/sirupsen/logrus"
	"health_checker/service/common"
)

func Report(nodeId, projectName, model string) error {
	fetchInfo := common.GetNvidiaFetchInfo()
	logrus.Infof("%v", map[string]interface{}{
		"node_id":        nodeId,
		"project_name":   projectName,
		"model":          model,
		"GPUName":        fetchInfo.GPUName,
		"UtilizationGPU": fetchInfo.UtilizationGPU,
		"MemoryTotal":    fetchInfo.MemoryTotal,
		"MemoryUsed":     fetchInfo.MemoryUsed,
	})

	//err := common.Post(config.ConfigInfo.DBCHealthCheckReportUrl(), map[string]interface{}{
	//	"node_id":        nodeId,
	//	"project_name":   projectName,
	//	"model":          model,
	//	"GPUName":        fetchInfo.GPUName,
	//	"UtilizationGPU": fetchInfo.UtilizationGPU,
	//	"MemoryTotal":    fetchInfo.MemoryTotal,
	//	"MemoryUsed":     fetchInfo.MemoryUsed,
	//})
	//if err != nil {
	//	return err
	//}

	return nil
}
