# ATHENA

> **Note:** This is an experimental project under active development. Please do not submit merge or commit requests at this time.

## Project Overview

ATHENA is a real-time collaborative code editor system, enabling multiple developers to edit code files simultaneously with instant synchronization—similar to Google Docs, but for code. The core value is zero-friction collaboration, eliminating merge conflicts during active editing. Text synchronization is achieved using Operational Transformation (OT), a proven algorithm for consistent document state across clients, even with concurrent edits.

Key features:
- Real-time bidirectional communication via WebSockets
- Operational Transformation for text consistency
- Multi-user sessions and presence indicators
- Professional code editor integration (Monaco)

The system architecture includes a Go server (sync hub), WebSocket clients, and optional Redis for session persistence.


## Phase 1 Goals

**Foundation: WebSocket Communication**
- Build a Go server to handle WebSocket connections
- Implement basic message passing (send/receive JSON)
- Manage connection lifecycle and active connection pool
- Route messages between clients in the same session
- Design a message protocol for join, leave, broadcast, and keepalive (ping/pong)
- Ensure robust error handling and logging

Testing includes connecting with browser WebSocket APIs, using CLI tools, and simulating multiple clients.

## Status
Currently in Phase 1: building core functionality and validating speed improvements on real repositories.

---
For more details, see `PhaseGoals Folder`.
