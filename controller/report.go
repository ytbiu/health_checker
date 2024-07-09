package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"health_checker/service"
)

type Response struct {
	Code   int    `json:"code"`
	ErrMsg string `json:"err_msg"`
}

type ReportReq struct {
	NodeId  string                  `json:"node_id" binding:"required"`
	Project string                  `json:"project" binding:"required"`
	Models  []service.RegisterModel `json:"models" binding:"required"`
}

func Report(c *gin.Context) {
	var req ReportReq
	if err := c.Bind(&req); err != nil {
		c.JSON(200, Response{
			Code:   -1,
			ErrMsg: fmt.Sprintf("request bind err : %s", err),
		})
		return
	}

	if err := service.Report(req.NodeId, req.Project, req.Models); err != nil {
		logrus.Errorf("service.Report err : %s", err)
		c.JSON(200, Response{
			Code:   -2,
			ErrMsg: fmt.Sprintf("internal err : %s", err),
		})
		return
	}
	c.JSON(200, Response{})
}
