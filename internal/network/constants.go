package network

import "time"

const (
	maxAttempts    = 5
	retryTimeout   = 3 * time.Second
	defaultTimeout = 10 * time.Second
)
