# Phase 5: File Management

**Goal:** Support multiple files in a project, not just single document.

## 5.1 File Tree Structure
```go
type FileNode struct {
    Name     string
    Type     string // "file" or "directory"
    Content  *Document // nil for directories
    Children []*FileNode
}

type Project struct {
    SessionID string
    Root      *FileNode
}
```

## 5.2 File Operations
- Create File: `{ "type": "file_create", "path": "src/main.go", "content": "" }`
- Open File for Editing: `{ "type": "file_open", "path": "src/main.go" }`
- Switch Files: Unsubscribe/subscribe to file operations
- File Operations to Support: Create/delete/rename/move files, create directory

## 5.3 Conflict Resolution
- Lock file during rename or queue operation

## 5.4 UI Components
- File tree sidebar, open files tabs, indicators for active editors

## Deliverables
- Browse file tree structure
- Open multiple files (one active at a time)
- Create/delete/rename files
- See which files teammates have open
