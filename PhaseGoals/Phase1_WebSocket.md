# Phase 1: Foundation — WebSocket Communication

**Goal:** Establish bidirectional real-time connection between clients and server.

## What You're Building
- Go server that handles WebSocket connections
- Basic message passing (send/receive JSON)
- Connection lifecycle management

## Components

### 1.1 WebSocket Server (Go)
**Core responsibilities:**
- Accept incoming WebSocket connections
- Maintain active connection pool
- Route messages between clients
- Handle disconnections gracefully

**Key Concepts:**
- One goroutine per WebSocket connection
- Use channels for message broadcasting
- Hub pattern for managing connections

### 1.2 Message Protocol Design
```json
{
  "type": "message_type",
  "payload": {...},
  "session_id": "abc123",
  "user_id": "alice",
  "timestamp": 1234567890
}
```

**Message Types to Implement:**
- `ping/pong` - Keep connection alive
- `join` - User enters session
- `leave` - User exits
- `broadcast` - Generic message distribution

## Deliverables
- Server accepts WebSocket connections
- Clients can send/receive messages
- Multiple clients in same "room"
- Basic error handling and logging

## Testing Strategy
- Connect with browser WebSocket API
- Use `websocat` CLI tool for debugging
- Simulate multiple clients with simple HTML page
