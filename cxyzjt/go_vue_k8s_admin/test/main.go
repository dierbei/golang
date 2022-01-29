package main

import (
	"github.com/gin-gonic/gin"
	"go_vue_k8s_admin/src/helpers"
	"go_vue_k8s_admin/src/wscore"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/tools/remotecommand"
	"log"
)

func main() {
	config, err := clientcmd.BuildConfigFromFlags("", "config")
	if err != nil {
		log.Fatal(err)
	}
	config.Insecure = true
	client, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatal(err)
	}

	r := gin.New()
	r.GET("/", func(c *gin.Context) {
		wsClient, err := wscore.Upgrader.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			log.Println(err)
			return
		}
		shellClient := wscore.NewWsShellClient(wsClient)
		err = helpers.HandleCommand(client, config, []string{"sh"}).
			Stream(remotecommand.StreamOptions{
				Stdin:  shellClient,
				Stdout: shellClient,
				Stderr: shellClient,
				Tty:    true,
			})
		if err != nil {
			log.Println(err)
		}

	})
	r.Run(":8080")
}
