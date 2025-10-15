# Phase 4: Session Management

**Goal:** Allow users to create/join editing sessions with unique identifiers.

## 4.1 Session Structure
```go
type Session struct {
    ID          string
    Document    *Document
    Clients     map[string]*Client
    CreatedAt   time.Time
    LastActive  time.Time
}
```

## 4.2 Session Lifecycle
- Creating Session: POST /api/session/create
- Joining Session: WebSocket connect to /ws?session=abc123
- Leaving Session: WebSocket disconnect

## 4.3 Session Storage
- In-Memory (MVP): `var sessions = make(map[string]*Session)`
- Redis (Production): Store document content, active user list, TTL for auto-cleanup

## 4.4 User Identification
- Simple Approach: UUID, name, color
- No auth required for MVP

## Deliverables
- Generate shareable session links
- Multiple users can join same session
- Session state persists during active use
- Users see who else is in session
