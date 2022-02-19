package lib

import (
	"fmt"
	"github.com/spf13/cobra"
	"golang.org/x/crypto/ssh"
	"log"
	"net"
	"os"
	"os/exec"
	"runtime"
	"time"
)

func clearConsole(){
	if runtime.GOOS=="windows" || runtime.GOOS=="darwin"{
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}else if runtime.GOOS=="windows"{
		cmd := exec.Command("cmd", "/c", "cls")

		cmd.Stdout = os.Stdout
		cmd.Run()
	}else{
		log.Fatal("目前clear操作只写了Linux、darwin、Windows3个。其他没做，累的")
	}
}
func mustFlag(name,t string,cmd *cobra.Command) interface{}{
	switch t{
	case "string":
		if v,err:=cmd.Flags().GetString(name);err==nil && v!=""{
			return v
		}
	case "int":
		if v,err:=cmd.Flags().GetInt(name);err==nil && v!=0{
			return v
		}
	}

	log.Fatal(name," is required")
	return nil
}
var  ShellModes  = ssh.TerminalModes{
	ssh.ECHO: 0,
	ssh.ECHOCTL: 0,
	ssh.TTY_OP_ISPEED: 14400, // input speed = 14.4kbaud
	ssh.TTY_OP_OSPEED: 14400, // output speed = 14.4kbaud
}
func SSHConnect( user, password, host string, port int ) ( *ssh.Session, error ) {
	var (
		auth         []ssh.AuthMethod
		addr         string
		clientConfig *ssh.ClientConfig
		client       *ssh.Client
		session      *ssh.Session
		err          error
	)
	// get auth method
	auth = make([]ssh.AuthMethod, 0)
	auth = append(auth, ssh.Password(password))
	hostKeyCallbk := func(hostname string, remote net.Addr, key ssh.PublicKey) error {
		return nil
	}

	clientConfig = &ssh.ClientConfig{
		User:               user,
		Auth:               auth,
	    Timeout:             5 * time.Second,  //增加了  超时时间
		HostKeyCallback:    hostKeyCallbk,
	}
	// connet to ssh
	addr = fmt.Sprintf( "%s:%d", host, port )
	if client, err = ssh.Dial( "tcp", addr, clientConfig ); err != nil {
		return nil, err
	}
	if session, err = client.NewSession(); err != nil {
		return nil, err
	}
	return session, nil
}

func translate(head string) string  {

	t:=map[string]string{
		"num":"行号",
		"pkts":"包数量",
		"bytes":"字节数",
		"target":"目标行为",
		"prot":"协议",
		"in":"输入",
		"out":"输出",
		"source":"源地址",
		"destination":"目标地址",
		"opt":"选项",
	}
	if v,ok:=t[head];ok{
		return v
	}
	return head
}

//func getSession(remoteName string ) *ssh.Session  {
//	remote:=SysConfig.GetRemote(remoteName)
//	if remote==nil{log.Fatal("no such remote")}
//	session,err:=SSHConnect(remote.User,remote.Pwd,remote.Host,22)
//	if err!=nil{log.Fatal(err)}
//	return session
//}