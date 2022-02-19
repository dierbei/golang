package lib

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

func init() {
	iptablesDropCMD.Flags().IntP("port", "p", 0, "set drop-port")
}

var iptablesDropCMD = &cobra.Command{
	Use:   "drop",
	Short: "drop port",
	Run: func(cmd *cobra.Command, args []string) {
		port := mustFlag("port", "int", cmd).(int)
		session, err := SSHConnect("root", "123456", "192.168.244.143", 22)
		if err != nil {
			log.Fatal(err)
		}

		err = session.Run(fmt.Sprintf("iptables  -A INPUT -p tcp --dport %d -j DROP", port))
		if err != nil {
			log.Fatal(err)
		}
		iptablesCMD.Run(cmd, args)
	},
}
