package main

import (
	"context"
	"fmt"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8sapi/lib"
	"log"
)

func main() {
	list, err := lib.K8sClient.CoreV1().Events("default").
		List(context.Background(), v1.ListOptions{})
	if err != nil {
		log.Fatal(err)
	}
	for _, item := range list.Items {
		fmt.Println(
			item.Name,
			item.Type,
			item.Reason,
			item.Message,
			item.Namespace,
			//item.InvolvedObject,
		)
	}
}
