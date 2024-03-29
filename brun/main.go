package main

import (
	_ "github.com/slomo0808/account"
	"github.com/slomo0808/infra"
	"github.com/slomo0808/infra/comm"
	"github.com/tietang/props/ini"
)

func main() {
	// 获取程序运行文件所在的路径
	path := comm.GetCurrentPath()
	// 加载和解析配置文件
	conf := ini.NewIniFileCompositeConfigSource(path + "/config.ini")

	app := infra.New(conf)
	app.Start()
}
