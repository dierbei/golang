package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	req, err := http.NewRequest("GET", "http://175.24.198.168:9090", nil)
	if err != nil {
		log.Fatal(err)
	}

	//req.Header.Add("Authorization", "Bearer kubeconfig-user-r8wp5.c-sld9n:hdkwqrwrblpmwkgj7m29vjgw7xqmsb2hqwnm7b9rgfgg7xdtft8lqr")
	rsp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	defer rsp.Body.Close()
	b, _ := ioutil.ReadAll(rsp.Body)
	fmt.Println(string(b))
}
