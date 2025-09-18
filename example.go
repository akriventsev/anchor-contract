package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

// SimulateExternalService simulates an external service that can fail
type SimulateExternalService struct {
	failureRate float64 // 0.0 to 1.0, probability of failure
	callCount   int
}

func NewSimulateExternalService(failureRate float64) *SimulateExternalService {
	return &SimulateExternalService{
		failureRate: failureRate,
		callCount:   0,
	}
}

func (s *SimulateExternalService) Call() (interface{}, error) {
	s.callCount++
	
	// Simulate some processing time
	time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
	
	// Simulate failure based on failure rate
	if rand.Float64() < s.failureRate {
		return nil, fmt.Errorf("service call failed (call #%d)", s.callCount)
	}
	
	return fmt.Sprintf("success (call #%d)", s.callCount), nil
}

func main() {
	// Create a circuit breaker with custom configuration
	config := &CircuitBreakerConfig{
		FailureThreshold:  3,  // Open circuit after 3 failures
		SuccessThreshold:  2,  // Close circuit after 2 successes in half-open
		Timeout:           5 * time.Second,  // Wait 5 seconds before trying half-open
		MaxRequests:       2,  // Allow max 2 requests in half-open state
		FailureCountReset: 10 * time.Second, // Reset failure count after 10 seconds
	}
	
	cb := NewCircuitBreaker(config)
	
	// Create a service that fails 70% of the time
	service := NewSimulateExternalService(0.7)
	
	fmt.Println("=== Circuit Breaker Example ===")
	fmt.Printf("Service failure rate: %.0f%%\n", service.failureRate*100)
	fmt.Printf("Circuit breaker config: %d failures to open, %d successes to close\n", 
		config.FailureThreshold, config.SuccessThreshold)
	fmt.Println()
	
	ctx := context.Background()
	
	// Make multiple calls to demonstrate circuit breaker behavior
	for i := 0; i < 20; i++ {
		result, err := cb.Execute(ctx, service.Call)
		
		state := cb.GetState()
		stats := cb.GetStats()
		
		if err != nil {
			fmt.Printf("Call %2d: ERROR - %s [State: %s, Failures: %v]\n", 
				i+1, err.Error(), state.String(), stats["failureCount"])
		} else {
			fmt.Printf("Call %2d: SUCCESS - %s [State: %s, Successes: %v]\n", 
				i+1, result, state.String(), stats["successCount"])
		}
		
		// Small delay between calls
		time.Sleep(500 * time.Millisecond)
	}
	
	fmt.Println("\n=== Final Statistics ===")
	stats := cb.GetStats()
	for key, value := range stats {
		fmt.Printf("%s: %v\n", key, value)
	}
}