package main

import "fmt"

type Server struct {
	Host          string
	Weight        int
	CurrentWeight int
}

// 用一个变量记录总权重
// 每一轮循环都叠加每个服务的当前权重
// 选取每一轮当前权重最高的服务
// 使用当前权重最高的服务扣减总权重
func getServer(servers []*Server) (s *Server) {
	allWeight := 0

	for _, server := range servers {
		allWeight += server.Weight
		server.CurrentWeight += server.Weight

		if s == nil || server.CurrentWeight > s.CurrentWeight {
			s = server
		}
	}

	s.CurrentWeight -= allWeight
	return
}

func main() {
	servers := []*Server{
		{"192.168.0.1", 5, 0},
		{"192.168.0.2", 2, 0},
		{"192.168.0.3", 1, 0},
	}

	for i := 0; i < 20; i++ {
		server := getServer(servers)
		fmt.Printf("| %s | %d | \n", server.Host, server.Weight)
	}
}
