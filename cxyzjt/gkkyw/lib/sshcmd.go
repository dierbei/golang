package lib

import (
	"fmt"
	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/terminal"
	"log"
	"os"
	"syscall"

	"github.com/spf13/cobra"
)

func init() {
	sshCMD.Flags().StringP("server", "s", "", "服务器地址")
	sshCMD.Flags().StringP("user", "u", "", "服务器用户名")
	//sshCMD.Flags().StringP("password", "p", "", "服务器密码")

	RootCmd.AddCommand(sshCMD)
}

var sshCMD = &cobra.Command{
	Use:   "ssh",
	Short: "SSH远程连接",
	Run: func(cmd *cobra.Command, args []string) {
		// 获取参数值
		server := mustFlag("server", "string", cmd).(string)
		user := mustFlag("user", "string", cmd).(string)
		//pwd := mustFlag("password", "string", cmd).(string)

		var session *ssh.Session
		connectCount := 0
		for connectCount < 3 {
			fmt.Println("请输入密码：")
			pwd, err := terminal.ReadPassword(int(syscall.Stdin))
			if err != nil {
				log.Fatal("ReadPassword failed")
			}

			session, err = SSHConnect(user, string(pwd), server, 22)
			if err != nil {
				fmt.Println("密码错误，请重试")
			}
			connectCount++

			if session != nil {
				break
			}
		}

		// 如果为空说明三次连接全部失败
		if session == nil {
			log.Fatal("SSHConnect failed")
		}

		// 连接成功建立连接
		err := session.RequestPty("", 0, 0, ShellModes)
		if err != nil {
			log.Fatal(err)
		}
		session.Stdin = os.Stdin
		session.Stdout = os.Stdout
		session.Stderr = os.Stderr
		err = session.Run("sh")
		if err != nil {
			log.Fatal(err)
		}
		session.Close()
	},
}
