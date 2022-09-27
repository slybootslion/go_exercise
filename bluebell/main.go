package main

import (
	"bluebell/dao/mysql"
	"bluebell/dao/redis"
	"bluebell/logger"
	"bluebell/routes"
	"bluebell/settings"
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"go.uber.org/zap"

	"github.com/spf13/viper"
)

func main() {
	// 加载配置项
	if err := settings.Init(); err != nil {
		fmt.Printf("init settings failed, err: %v\n", err)
		return
	}
	// 初始化日志
	if err := logger.Init(settings.Conf.LogConfig); err != nil {
		fmt.Printf("init logger failed, err: %v\n", err)
		return
	}
	defer zap.L().Sync()
	zap.L().Debug("logger init success ...")
	// db
	if err := mysql.Init(settings.Conf.MySQLConfig); err != nil {
		fmt.Printf("init mysql failed, err: %v\n", err)
		return
	}
	if err := redis.Init(settings.Conf.RedisConfig); err != nil {
		fmt.Printf("init redis failed, err: %v\n", err)
		return
	}
	defer mysql.Close()
	defer redis.Close()

	// 路由
	router := routes.Setup()

	// 启动服务
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", viper.GetInt("app.port")),
		Handler: router,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			zap.L().Fatal("listen: ", zap.Error(err))
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	zap.L().Info("Shutdown Server ...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		zap.L().Fatal("Server Shutdown", zap.Error(err))
	}

	zap.L().Info("Server exiting")
}
