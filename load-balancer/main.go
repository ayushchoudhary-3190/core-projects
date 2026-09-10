package main

import (
	"sync"
	"errors"
)

// LoadBalancer distributes requests across backends using round-robin.
type LoadBalancer struct {
	mu sync.Mutex
	// TODO: Add fields for backends list and current index
	backends []string
	current	int
}

// NewLoadBalancer creates a new round-robin load balancer.
// TODO: Implement this function
func NewLoadBalancer() *LoadBalancer {
	return &LoadBalancer{
		backends : make([]string,0),                  
	}
}

// AddBackend adds a backend address to the pool. Duplicates are ignored.
// TODO: Implement this function
func (lb *LoadBalancer) AddBackend(addr string) {
	lb.mu.Lock()
	defer lb.mu.Unlock()
	for _,b := range lb.backends{
		if addr == b {
			return
		}
	}
	lb.backends = append(lb.backends, addr)
}

// RemoveBackend removes a backend address from the pool.
// TODO: Implement this function
func (lb *LoadBalancer) RemoveBackend(addr string) {
	lb.mu.Lock()
	lb.mu.Unlock()
	for i,b := range lb.backends{
		if b == addr{
			lb.backends = append(lb.backends[:i],lb.backends[i+1:]...)
			if lb.current >= len(lb.backends){
				lb.current = 0
			}
			return
		}
		
	}

}

// NextBackend returns the next backend address in round-robin order.
// Returns an error if no backends are available.
// TODO: Implement this function
func (lb *LoadBalancer) NextBackend() (string, error) {
	lb.mu.Lock()
	lb.mu.Unlock()
	if len(lb.backends) == 0{
		return "" , errors.New("no new backends")
	}
	if lb.current>= len(lb.backends){
		lb.current = 0
	}
	b := lb.backends[lb.current]
	lb.current = lb.current+1 % len(lb.backends)

	return b, nil
}

// Backends returns a copy of the current backend list.
// TODO: Implement this function
func (lb *LoadBalancer) Backends() []string {
	result := make([]string,len(lb.backends))
	copy(result,lb.backends)
	return result
}

func main() {}