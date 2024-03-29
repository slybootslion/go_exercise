package miniblog

import (
	"context"
	"errors"
	"fmt"
	"github.com/slybootslion/miniblog-t/internal/pkg/core"
	"github.com/slybootslion/miniblog-t/internal/pkg/errno"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/slybootslion/miniblog-t/internal/pkg/log"
	mw "github.com/slybootslion/miniblog-t/internal/pkg/middleware"
	"github.com/slybootslion/miniblog-t/pkg/version/verflag"
)

var cfgFile string

// NewMiniBlogCommand 创建一个 *cobra.Command 对象. 之后，可以使用 Command 对象的 Execute 方法来启动应用程序.
func NewMiniBlogCommand() *cobra.Command {
	cmd := &cobra.Command{
		// 指定命令的名字，该名字会出现在帮助信息中
		Use: "miniblog",
		// 命令的简短描述
		Short: "A good Go practical project",
		// 命令的详细描述
		Long: `A good Go practical project, used to create user with basic information.

Find more miniblog information at:
	https://github.com/marmotedu/miniblog#readme`,

		// 命令出错时，不打印帮助信息。不需要打印帮助信息，设置为 true 可以保持命令出错时一眼就能看到错误信息
		SilenceUsage: true,
		// 指定调用 cmd.Execute() 时，执行的 Run 函数，函数执行失败会返回错误信息
		RunE: func(cmd *cobra.Command, args []string) error {
			// 如果 `--version=true`，则打印版本并退出
			verflag.PrintAndExitIfRequested()

			// 初始化日志
			log.Init(logOptions())
			defer log.Sync() // Sync 将缓存中的日志刷新到磁盘文件中

			return run()
		},
		// 这里设置命令运行时，不需要指定命令行参数
		Args: func(cmd *cobra.Command, args []string) error {
			for _, arg := range args {
				if len(arg) > 0 {
					return fmt.Errorf("%q does not take any arguments, got %q", cmd.CommandPath(), args)
				}
			}

			return nil
		},
	}

	// 以下设置，使得 initConfig 函数在每个命令运行时都会被调用以读取配置
	cobra.OnInitialize(initConfig)

	// 在这里您将定义标志和配置设置。

	// Cobra 支持持久性标志(PersistentFlag)，该标志可用于它所分配的命令以及该命令下的每个子命令
	cmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "The path to the miniblog configuration file. Empty string for no configuration file.")

	// Cobra 也支持本地标志，本地标志只能在其所绑定的命令上使用
	cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	// 添加 --version 标志
	verflag.AddFlags(cmd.PersistentFlags())

	return cmd
}

// run 函数是实际的业务代码入口函数.
func run() error {
	// 设置Gin模式
	gin.SetMode(viper.GetString("runmode"))
	// 创建Gin引擎
	g := gin.New()
	// gin.Recovery() 中间件，用来捕获任何panic，并恢复
	mws := []gin.HandlerFunc{gin.Recovery(), mw.NoCache, mw.Cors, mw.Secure, mw.RequestId()}
	g.Use(mws...)
	// 注册404
	g.NoRoute(func(c *gin.Context) {
		core.WriteResponse(c, errno.ErrPageNotFond, nil)
	})
	// 注册 /healthz
	g.GET("/healthz", func(c *gin.Context) {
		log.C(c).Infow("healthz function called.")
		core.WriteResponse(c, nil, map[string]string{"status": "ok"})
	})
	// 创建http server实例
	httpsrv := &http.Server{Addr: viper.GetString("addr"), Handler: g}
	// 运行http服务器
	// 打印一条日志，用来提示HTTP服务已经启动，方便排障
	log.Infow("start to listening th incoming requests on http address",
		"addr",
		viper.GetString("addr"))
	go func() {
		if err := httpsrv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalw(err.Error())
		}
	}()

	// 等待终端信号优雅的关闭服务器（10秒超时）
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Infow("Shutting down server ...")
	// 创建ctx用于通知服务器goroutine，它有10秒时间完成当前正在处理的请求
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	// 10秒内优雅关闭服务，超过10秒就超时退出
	if err := httpsrv.Shutdown(ctx); err != nil {
		log.Errorw("insecure server forced to shutdown", "err", err)
		return err
	}
	log.Infow("Server exiting")

	return nil
}
