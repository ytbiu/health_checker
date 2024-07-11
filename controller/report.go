package controller

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
	"health_checker/service"
	"health_checker/service/common"
	"net/http"
)

type Response struct {
	Code   int    `json:"code"`
	ErrMsg string `json:"err_msg"`
}

type ReportReqFromAIProxyService struct {
	NodeId  string                  `json:"node_id" binding:"required"`
	Project string                  `json:"project" binding:"required"`
	Models  []service.RegisterModel `json:"models" binding:"required"`
}

type ReportData struct {
	NodeId         string                  `json:"node_id"`
	Project        string                  `json:"project"`
	Models         []service.RegisterModel `json:"models"`
	GPUName        string                  `json:"gpu_name"`
	UtilizationGPU string                  `json:"utilization_gpu"`
	MemoryTotal    string                  `json:"memory_total"`
	MemoryUsed     string                  `json:"memory_used"`
}

func Report(c *gin.Context) {
	upgrader := websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool { return true },
	}
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		logrus.Error("build WebSocket connect:", err)
		return
	}
	defer conn.Close()

	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			logrus.Error("read message:", err)
			return
		}

		var data ReportReqFromAIProxyService
		err = json.Unmarshal(message, &data)
		if err != nil {
			logrus.Println("Unmarshal received data err:", err)
			return
		}
		logrus.Debugf("received data: %+v", data)

		fetchInfo := common.GetNvidiaFetchInfo()
		fullData := ReportData{
			NodeId:         data.NodeId,
			Project:        data.Project,
			Models:         data.Models,
			GPUName:        fetchInfo.GPUName,
			UtilizationGPU: fetchInfo.UtilizationGPU,
			MemoryTotal:    fetchInfo.MemoryTotal,
			MemoryUsed:     fetchInfo.MemoryUsed,
		}

		logrus.Infof("fulldata: %+v", fullData)
		//conn := service.GetReportToClusterWsConn()
		//b, _ := json.Marshal(fullData)
		//if err := conn.WriteMessage(websocket.TextMessage, b); err != nil {
		//	logrus.Error("write to DBC Health check cluster err : ", err)
		//}
	}
}
