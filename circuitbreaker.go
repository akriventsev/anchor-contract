package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// CircuitState represents the state of the circuit breaker
type CircuitState int

const (
	StateClosed CircuitState = iota
	StateOpen
	StateHalfOpen
)

func (s CircuitState) String() string {
	switch s {
	case StateClosed:
		return "CLOSED"
	case StateOpen:
		return "OPEN"
	case StateHalfOpen:
		return "HALF_OPEN"
	default:
		return "UNKNOWN"
	}
}

// CircuitBreakerConfig holds configuration for the circuit breaker
type CircuitBreakerConfig struct {
	FailureThreshold    int           // Number of failures before opening circuit
	SuccessThreshold    int           // Number of successes to close circuit from half-open
	Timeout             time.Duration // How long to wait before trying half-open
	MaxRequests         int           // Max requests allowed in half-open state
	FailureCountReset   time.Duration // How long to wait before resetting failure count
}

// DefaultConfig returns a sensible default configuration
func DefaultConfig() *CircuitBreakerConfig {
	return &CircuitBreakerConfig{
		FailureThreshold:  5,
		SuccessThreshold:  3,
		Timeout:           30 * time.Second,
		MaxRequests:       3,
		FailureCountReset: 60 * time.Second,
	}
}

// CircuitBreaker implements the circuit breaker pattern
type CircuitBreaker struct {
	config        *CircuitBreakerConfig
	state         CircuitState
	failureCount  int
	successCount  int
	lastFailTime  time.Time
	lastResetTime time.Time
	mutex         sync.RWMutex
}

// NewCircuitBreaker creates a new circuit breaker with the given config
func NewCircuitBreaker(config *CircuitBreakerConfig) *CircuitBreaker {
	if config == nil {
		config = DefaultConfig()
	}
	
	return &CircuitBreaker{
		config:        config,
		state:         StateClosed,
		failureCount:  0,
		successCount:  0,
		lastFailTime:  time.Time{},
		lastResetTime: time.Now(),
	}
}

// Execute runs a function through the circuit breaker
func (cb *CircuitBreaker) Execute(ctx context.Context, fn func() (interface{}, error)) (interface{}, error) {
	cb.mutex.Lock()
	defer cb.mutex.Unlock()

	// Check if we should reset failure count
	if time.Since(cb.lastResetTime) > cb.config.FailureCountReset {
		cb.failureCount = 0
		cb.lastResetTime = time.Now()
	}

	// Check circuit state
	switch cb.state {
	case StateOpen:
		if time.Since(cb.lastFailTime) < cb.config.Timeout {
			return nil, fmt.Errorf("circuit breaker is OPEN")
		}
		// Transition to half-open
		cb.state = StateHalfOpen
		cb.successCount = 0

	case StateHalfOpen:
		if cb.successCount >= cb.config.SuccessThreshold {
			// Transition to closed
			cb.state = StateClosed
			cb.failureCount = 0
		} else if cb.failureCount >= cb.config.FailureThreshold {
			// Transition back to open
			cb.state = StateOpen
			cb.lastFailTime = time.Now()
			return nil, fmt.Errorf("circuit breaker is OPEN")
		}
	}

	// Execute the function
	result, err := fn()
	
	if err != nil {
		cb.failureCount++
		cb.lastFailTime = time.Now()
		
		// Check if we should open the circuit
		if cb.failureCount >= cb.config.FailureThreshold {
			cb.state = StateOpen
		}
		
		return nil, err
	}

	// Success
	cb.successCount++
	if cb.state == StateHalfOpen && cb.successCount >= cb.config.SuccessThreshold {
		cb.state = StateClosed
		cb.failureCount = 0
	}

	return result, nil
}

// GetState returns the current state of the circuit breaker
func (cb *CircuitBreaker) GetState() CircuitState {
	cb.mutex.RLock()
	defer cb.mutex.RUnlock()
	return cb.state
}

// GetStats returns statistics about the circuit breaker
func (cb *CircuitBreaker) GetStats() map[string]interface{} {
	cb.mutex.RLock()
	defer cb.mutex.RUnlock()
	
	return map[string]interface{}{
		"state":         cb.state.String(),
		"failureCount":  cb.failureCount,
		"successCount":  cb.successCount,
		"lastFailTime":  cb.lastFailTime,
		"lastResetTime": cb.lastResetTime,
	}
}

// Reset resets the circuit breaker to closed state
func (cb *CircuitBreaker) Reset() {
	cb.mutex.Lock()
	defer cb.mutex.Unlock()
	
	cb.state = StateClosed
	cb.failureCount = 0
	cb.successCount = 0
	cb.lastFailTime = time.Time{}
	cb.lastResetTime = time.Now()
}