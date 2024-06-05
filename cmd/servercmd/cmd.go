package servercmd

import (
	"context"
	"github.com/spf13/cobra"
	"preject/internal/app/httpserv"
)

var Cmd = &cobra.Command{
	Use:   "server",
	Short: "启动 web 服务",
	Long:  `启动 web 服务`,
	RunE: func(cmd *cobra.Command, args []string) error {
		httpPort, err := cmd.Flags().GetUint("http")
		if err != nil {
			httpPort = 8080
		}
		return Run(cmd.Context(), httpPort)
	},
}

// 命令行需要啥启动参数，就在这里注册
func initCmdParams(c *cobra.Command) {
	c.Flags().UintP("http", "p", 8080, "监听端口号")
}

func Run(ctx context.Context, httpPort uint) error {
	ch := make(chan error)
	go func() {
		ch <- httpserv.Start(ctx, httpPort)
	}()
	select {
	case <-ctx.Done():
		return ctx.Err()
	case err := <-ch:
		return err
	}
}

func init() {
	initCmdParams(Cmd)
}
