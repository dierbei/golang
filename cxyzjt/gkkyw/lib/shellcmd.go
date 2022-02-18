package lib

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
)

func init() {
	shellCMD.Flags().StringP("server", "s", "", "服务器地址")
	shellCMD.Flags().StringP("user", "u", "", "服务器用户名")
	shellCMD.Flags().StringP("password", "p", "", "服务器密码")
	shellCMD.Flags().StringP("command", "c", "", "待执行命令")

	RootCmd.AddCommand(shellCMD)
}

var shellCMD = &cobra.Command{
	Use:   "shell",
	Short: "执行shell命令",
	Run: func(cmd *cobra.Command, args []string) {
		// 获取参数值
		server := mustFlag("server", "string", cmd).(string)
		user := mustFlag("user", "string", cmd).(string)
		pwd := mustFlag("password", "string", cmd).(string)
		command := mustFlag("command", "string", cmd).(string)

		session, err := SSHConnect(user, pwd, server, 22)
		if err != nil {
			fmt.Println("密码错误，请重试")
		}
		defer session.Close()

		session.Stdin = os.Stdin
		session.Stdout = os.Stdout
		session.Stderr = os.Stderr
		err = session.Run(command)
		if err != nil {
			log.Fatal(err)
		}
	},
}
