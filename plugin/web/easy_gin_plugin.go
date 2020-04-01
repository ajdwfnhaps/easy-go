package web

import (
	easygin "github.com/ajdwfnhaps/easy-gin"
	mw "github.com/ajdwfnhaps/easy-gin/middleware"
	"github.com/ajdwfnhaps/easy-go/routers/api"
)

//UseEasyGinPlugin 使用easy-gin插件
func UseEasyGinPlugin(fPath string) {
	//创建应用程序 使用配置文件
	//r := easygin.New("conf/config.toml")
	r := easygin.Default(fPath)

	//使用跨域请求中间件
	r.UseCors()

	//使用logrus日志组件
	//指定api路径规则才记录日志
	apiPrefixes := []string{"/api/"}
	r.UseLogrusConf(mw.AllowPathPrefixNoSkipper(apiPrefixes...))

	//注册路由
	r.RegisterRouter(api.RouterHanlder)

	//使用swagger
	r.UseSwagger(SetSwaggerInfo)

	//使用静态站点中间件
	r.UseWWWRoot()

	//启动
	r.Run()
}
