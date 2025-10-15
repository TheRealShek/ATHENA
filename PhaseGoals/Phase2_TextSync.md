# Phase 2: Text Synchronization Engine

**Goal:** Implement algorithm that keeps text consistent across all clients despite concurrent edits.

## The Core Problem
Initial: "Hello"
User A: Inserts " World" at position 5 → "Hello World"
User B: Inserts "!" at position 5 → "Hello!"

Without sync algorithm: Conflicting states
With sync: Operations transform to reach consistent state

## Algorithm Choice: Operational Transformation (OT)

**Why OT for MVP:**
- Simpler to understand than CRDT
- Proven (Google Docs, Firepad use variants)
- Sufficient for text editing

### 2.1 Operation Types
```go
type Operation struct {
    Type     string // "insert", "delete", "retain"
    Position int    // Where in document
    Text     string // What to insert (for insert ops)
    Length   int    // How many chars (for delete/retain)
    UserID   string
    Version  int    // Document version when created
}
```

**Examples:**
- Insert "x" at position 5: `{Type: "insert", Position: 5, Text: "x"}`
- Delete 3 chars at position 2: `{Type: "delete", Position: 2, Length: 3}`

### 2.2 Transform Function
transform(op1, op2) → (op1', op2')
Purpose: Adjust operations so they can be applied in any order and reach the same final state.

**Key Rules:**
- If both insert at same position: One goes first (break tie by user ID)
- If one inserts before another's operation: Adjust positions
- If one deletes text another is editing: Handle carefully

### 2.3 Server-Side Document State
```go
type Document struct {
    Content    string        // Current text
    Version    int           // Increments on each operation
    Operations []Operation   // History (for late-joiners)
}
```

**Server Responsibilities:**
- Receive operation from client
- Transform against concurrent operations
- Apply to document
- Broadcast transformed operation to other clients
- Increment version

### 2.4 Client-Side State
```go
type ClientState struct {
    Content       string
    Version       int
    PendingOps    []Operation // Sent but not acknowledged
    Buffer        []Operation // Created during pending
}
```

## Deliverables
- Single-file text syncs between clients
- Concurrent edits don't cause corruption
- Late-joiners receive full document state
- Basic version tracking

## Testing Strategy
- Two browser tabs editing same document
- Rapid concurrent typing stress test
- Network delay simulation (throttle connection)
- Disconnect/reconnect scenarios
