# Phase 3: Editor Interface Integration

**Goal:** Replace textarea with professional code editor that supports syntax highlighting and features.

## 3.1 Monaco Editor Setup
- Monaco is VS Code's editor engine
- Runs in browser
- Syntax highlighting for 50+ languages
- Built-in autocomplete, minimap, etc.

**Integration Points:**
```javascript
// Listen to content changes
editor.onDidChangeModelContent((event) => {
    // Convert Monaco's change to your Operation format
    const op = convertToOperation(event);
    sendToServer(op);
});

// Apply remote changes
function applyRemoteOperation(op) {
    // Convert Operation to Monaco edit
    const edit = convertToMonacoEdit(op);
    editor.executeEdits('remote', [edit]);
}
```

## 3.2 Change Event Handling
Monaco provides changes as ranges, you need operations.

## 3.3 Cursor Synchronization
Display other users' cursors:
```javascript
{
    type: "cursor",
    user_id: "alice",
    position: 42,
    color: "#FF5733"
}
// Render as decoration in Monaco
editor.deltaDecorations(oldDecorations, [{
    range: new Range(line, col, line, col),
    options: {
        className: 'remote-cursor',
        hoverMessage: 'Alice is here'
    }
}]);
```

## Deliverables
- Monaco editor embedded in web page
- Local edits convert to operations correctly
- Remote operations apply to Monaco
- Syntax highlighting works
- Cursor positions visible
