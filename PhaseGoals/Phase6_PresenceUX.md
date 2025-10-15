# Phase 6: Presence & User Experience

**Goal:** Make collaboration feel natural and aware.

## 6.1 User Presence Indicators
- Active users panel: user_id, name, avatar_color, current_file, cursor_position, last_seen
- Visual elements: user list, colored dots, file indicators

## 6.2 Cursor & Selection Sync
- Broadcast cursor movements (debounced)
- Render remote cursors (color-coded)

## 6.3 Activity Indicators
- Typing indicator
- File badges for active editors

## 6.4 Notifications
- System events: join/leave, file open/close
- Display as toast or activity feed

## 6.5 Conflict Warnings
- Highlight concurrent edits, show tooltip, optional line lock

## Deliverables
- See all active users
- Cursor positions visible and color-coded
- Notifications for join/leave events
- Typing indicators
- Visual feedback for conflicts
