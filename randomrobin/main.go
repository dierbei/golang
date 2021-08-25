package main

import (
	"fmt"
	"math/rand"
)

type Random struct {
	Servers []string
}

func (r *Random) next() string {
	return r.Servers[rand.Intn(len(r.Servers))]
}

func main() {
	r := Random{
		Servers: []string{"192.168.0.1", "192.168.0.2", "192.168.0.3"},
	}

	for i := 0; i < 10; i++ {
		fmt.Println(r.next())
	}
}
