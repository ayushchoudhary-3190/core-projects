package loadbalancer

import "sync"

type LoadBalancer struct{
	mu sync.Mutex
	backends []string
	current  int
}

type Request struct{
	Method string
	Path   string
	Headers map[string]string
}