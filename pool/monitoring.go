// Copyright (C) MongoDB, Inc. 2017-present.
//
// Licensed under the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at http://www.apache.org/licenses/LICENSE-2.0

package pool

import (
	"context"
)

// CommandStartedEvent represents an event generated when a command is sent to a server.
type CommandStartedEvent struct {
	//Command      bson.Raw
	DatabaseName string
	CommandName  string
	RequestID    int64
	ConnectionID string
}

// CommandFinishedEvent represents a generic command finishing.
type CommandFinishedEvent struct {
	DurationNanos int64
	CommandName   string
	RequestID     int64
	ConnectionID  string
}

// CommandSucceededEvent represents an event generated when a command's execution succeeds.
type CommandSucceededEvent struct {
	CommandFinishedEvent
	//Reply bson.Raw
}

// CommandFailedEvent represents an event generated when a command's execution fails.
type CommandFailedEvent struct {
	CommandFinishedEvent
	Failure string
}

// CommandMonitor represents a monitor that is triggered for different events.
type CommandMonitor struct {
	Started   func(context.Context, *CommandStartedEvent)
	Succeeded func(context.Context, *CommandSucceededEvent)
	Failed    func(context.Context, *CommandFailedEvent)
}

// strings for pool command monitoring reasons
const (
	ReasonPoolClosed        = "poolClosed"
	ReasonStale             = "stale"
	ReasonConnectionErrored = "connectionError"
	ReasonTimedOut          = "timeout"
)

// strings for pool command monitoring types
const (
	ConnectionClosed   = "ConnectionClosed"
	Created            = "ConnectionPoolCreated"
	ConnectionCreated  = "ConnectionCreated"
	GetFailed          = "ConnectionCheckOutFailed"
	GetSucceeded       = "ConnectionCheckedOut"
	ConnectionReturned = "ConnectionCheckedIn"
	Cleared            = "ConnectionPoolCleared"
	Closed             = "ConnectionPoolClosed"
)

// MonitorPoolOptions contains pool options as formatted in pool events
type MonitorPoolOptions struct {
	MaxPoolSize uint64 `json:"maxPoolSize"`
	MinPoolSize uint64 `json:"minPoolSize"`
}

// Event contains all information summarizing a pool event
type Event struct {
	Type         string              `json:"type"`
	Address      string              `json:"address"`
	ConnectionID uint64              `json:"connectionId"`
	PoolOptions  *MonitorPoolOptions `json:"options"`
	Reason       string              `json:"reason"`
}

// Monitor is a function that allows the user to gain access to events occurring in the pool
type Monitor struct {
	Event func(*Event)
}
