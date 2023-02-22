package initialize

func Run() {
	// 加载配置
	LoadConfig()
	// 加载数据库
	Mysql()
	// 加载redis
	Redis()
	// 启动定时任务
	go Cron()
	// 加载路由
	Router()
}
