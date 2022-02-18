package lib

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"

	"github.com/spf13/cobra"
)

func init() {
	scpCMD.Flags().StringP("server", "s", "", "服务器地址")
	scpCMD.Flags().StringP("user", "u", "", "服务器用户名")
	scpCMD.Flags().StringP("password", "p", "", "服务器密码")
	scpCMD.Flags().StringP("source", "", "", "set source path")
	scpCMD.Flags().StringP("dest", "", "", "set remote path")

	RootCmd.AddCommand(scpCMD)
	//xxx scp -n xxx  -s 源文件路径 -d 目标路径
}

var scpCMD = &cobra.Command{
	Use:   "scp",
	Short: "scp 拷贝本地文件到远程",
	Run: func(cmd *cobra.Command, args []string) {
		// 参数
		server := mustFlag("server", "string", cmd).(string)
		user := mustFlag("user", "string", cmd).(string)
		pwd := mustFlag("password", "string", cmd).(string)
		localPath := mustFlag("source", "string", cmd).(string)
		remotePath := mustFlag("dest", "string", cmd).(string)

		// 建立连接
		session, err := SSHConnect(user, pwd, server, 22)
		if err != nil {
			log.Fatal(err)
		}
		//到这里说明远程主机连上了
		defer session.Close()

		// 远程输入流
		in, _ := session.StdinPipe()

		// 读取本地文件
		localFile, _ := os.Open(localPath)
		fileBytes, _ := ioutil.ReadAll(localFile)
		localReader := bytes.NewReader(fileBytes)
		defer localFile.Close()

		// 执行scp命令
		err = session.Start(fmt.Sprintf("/usr/bin/scp %s %s ", "-t", remotePath))
		_, err = fmt.Fprintln(in, "C0644", int64(len(fileBytes)), localFile.Name())
		if err != nil {
			log.Fatal(err)
		}

		// 拷贝内容
		n, err := io.Copy(in, localReader)
		if err != nil {
			log.Fatal(err)
		}
		// 添加结束符
		_, err = fmt.Fprint(in, "\x00")
		if err != nil {
			log.Fatal(err)
		}
		// 关闭远程输入流
		err = in.Close()
		if err != nil {
			log.Fatal(err)
		}

		if err = session.Wait(); err != nil {
			log.Fatal(err)
		}

		fmt.Println("传输了", n, "个字节")
	},
}
