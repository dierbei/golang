package main

import "fmt"

type RoundRobin struct {
	servers []string
	current int
}

// -1 0 0
// 0 1 1
// 1 2 2
// 2 3 3
// 3 4 4
// 4 5 0
func (r *RoundRobin) next() string {
	r.current++
	r.current = r.current % len(r.servers)
	return r.servers[r.current]
}

func main() {
	r := &RoundRobin{
		servers: []string{"192.168.0.1", "192.168.0.2", "192.168.0.3"},
		current: -1,
	}

	for i := 0; i < 10; i++ {
		fmt.Printf("| %d | %s |\n", i+1, r.next())
	}
}