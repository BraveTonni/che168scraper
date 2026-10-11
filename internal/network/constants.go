package network

import "time"

const maxAttempts = 5
const retryTimeout = 3 * time.Second
const defaultTimeout = 10 * time.Second
