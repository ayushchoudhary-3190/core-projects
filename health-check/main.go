package main

import (
	"encoding/json"
	"net/http"
)

// CheckResult represents the result of a single health check.
type CheckResult struct {
	Status string `json:"status"`
	Error  string `json:"error,omitempty"`
}

// HealthResponse is the JSON response from the health endpoint.
type HealthResponse struct {
	Status string                 `json:"status"`
	Checks map[string]CheckResult `json:"checks"`
}

// HealthChecker manages health check functions for dependencies.
type HealthChecker struct {
	// TODO: Add fields for storing named check functions
	checks 	map[string]func() error
}

// NewHealthChecker creates a new health checker.
// TODO: Implement this function
func NewHealthChecker() *HealthChecker {
	return &HealthChecker{checks : make(map[string]func() error)}
}

// AddCheck registers a named health check function.
// The function returns nil if healthy, or an error if unhealthy.
// TODO: Implement this function
func (hc *HealthChecker) AddCheck(name string, check func() error) {
	hc.checks[name] = check
}

// Handler returns an http.Handler that serves the health check endpoint.
// - Runs all registered checks
// - Returns 200 and {"status":"healthy",...} if all pass
// - Returns 503 and {"status":"unhealthy",...} if any fail
// TODO: Implement this function 
func (hc *HealthChecker) Handler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		response := HealthResponse{
			Status : "healthy",
			Checks : make(map[string]CheckResult),
		}
		for name,check := range hc.checks{
			if err := check();err!=nil{
				response.Status = "unhealthy"
				response.Checks[name] = CheckResult{
					Status: "unhealthy",
					Error : err.Error(),
				}
			}else{
				response.Checks[name] = CheckResult{
					Status : "healthy",
				}
			}
		}
		w.Header().Set("Content-Type","application/json")
		if response.Status == "unhealthy"{
			w.WriteHeader(http.StatusServiceUnavailable)
		}else{
			w.WriteHeader(http.StatusOK)
		}
		json.NewEncoder(w).Encode(response)
	})
}

func main() {}