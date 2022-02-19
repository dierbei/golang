package lib

import (
	"fmt"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strings"
)

var chanInfo = regexp.MustCompile(`^Chain\s*(INPUT|FORWARD|OUTPUT)`)
var headerInfo = regexp.MustCompile(`^num\s+pkts\s+bytes`)

func isChanInfo(str string) bool {
	return chanInfo.MatchString(str)
}
func isHeaderInfo(str string) bool {
	return headerInfo.MatchString(str)
}

func printtable(header []string, data [][]string) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader(header)
	for _, v := range data {
		table.Append(v)
	}
	table.Render()
}
func filterlist(str string, istranslate bool) []string {
	list := strings.Split(str, " ")
	ret := []string{}
	for _, item := range list {
		if item = strings.Trim(item, " "); item != "" {
			if istranslate {
				item = translate(item)
			}
			ret = append(ret, item)
		}
	}
	return ret
}

//表格形式打印
func render(str string) {
	list := strings.Split(str, "\n")
	header := []string{}
	data := [][]string{}
	begin := true
	for _, item := range list {
		if item = strings.Trim(item, " "); item != "" {
			if isChanInfo(item) {
				begin = true
			} else {
				begin = false
			}
			if begin { //需要打印
				if len(header) > 0 {
					printtable(header, data)
					header = []string{}
					data = [][]string{}
				}
				fmt.Println("链信息", item)
			}
			if isHeaderInfo(item) {
				header = append(header, filterlist(item, true)...)
			}
			if !isChanInfo(item) && !isHeaderInfo(item) {
				data = append(data, filterlist(item, false))
			}

		}
	}
	printtable(header, data)

}

func init() {
	iptablesCMD.PersistentFlags().StringP("table", "t", "filter", "set table - filter、nat、")
	iptablesCMD.AddCommand(iptablesDropCMD)
	iptablesCMD.AddCommand(iptablesDeleteCMD)
	RootCmd.AddCommand(iptablesCMD)
}

var iptablesCMD = &cobra.Command{
	Use:   "iptables",
	Short: "iptables command",
	Run: func(cmd *cobra.Command, args []string) {
		table := mustFlag("table", "string", cmd).(string)
		session, err := SSHConnect("root", "123456", "192.168.244.143", 22)
		if err != nil {
			log.Fatal(err)
		}

		out, _ := session.StdoutPipe()
		err = session.Run(fmt.Sprintf("iptables -nvL -t %s --line-numbers", table))
		if err != nil {
			log.Fatal(err)
		}
		b, _ := ioutil.ReadAll(out)
		render(string(b))
	},
}
