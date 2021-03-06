/*
* @Time    : 2020-11-17 11:47
* @Author  : CoderCharm
* @File    : server.go
* @Software: GoLand
* @Github  : github/CoderCharm
* @Email   : wg_python@163.com
* @Desc    : 启动服务
**/
package core

import (
	"gin_study/global"
	"gin_study/initialize"
	"go.uber.org/zap"
	"net/http"
	"time"
)

func RunWindowsServer() {

	// 初始化redis
	initialize.Redis()

	// 初始化路由
	Router := initialize.Routers()

	// 取系统配置的地址 并格式化:
	address := global.GIN_CONFIG.System.Addr

	global.GIN_LOG.Info("当前地址为:", zap.Any("ipAddress", address))

	//_ = Router.Run(address) // gin启动web服务
	// 启用原生的 web服务器 方便其他配置
	s := &http.Server{
		Addr:           address,
		Handler:        Router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe().Error()
}
