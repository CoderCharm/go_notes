/*
* @Time    : 2020-11-17 10:42
* @Author  : CoderCharm
* @File    : main.go
* @Software: GoLand
* @Github  : github/CoderCharm
* @Email   : wg_python@163.com
* @Desc    :
**/
package main

import (
	"gin_study/core"
	"gin_study/global"
	"gin_study/initialize"
)

func main() {

	global.GIN_VP = core.Viper() // 初始化Viper

	global.GIN_LOG = core.Zap() // 初始化zap日志库

	global.GIN_DB = initialize.Gorm() // gorm连接数据库

	global.GIN_CRON = initialize.SysCron() // 初始化定时任务

	// 程序结束前关闭数据库链接
	db, _ := global.GIN_DB.DB()

	defer db.Close()

	//initialize.MysqlTables(global.GIN_DB)

	// 启动定时任务
	//go global.GIN_CRON.Start()

	//defer global.GIN_CRON.Stop()

	core.RunWindowsServer()

}
