# Go Formatter & Import-Fixer Development Plan

## Overview
Build a fast, incremental formatter & import-fixer for Go that only processes changed files, runs formatters in parallel, and intelligently caches results.

---

## Core Features
- Detects changed files (hash / mtime+size)
- Runs `gofmt`/`goimports` only on changed files
- Executes formatters in parallel using a worker pool
- Caches results (file → hash, import resolution) with intelligent invalidation
- Writes back atomically
- Reports concise CI-friendly results (`--check`, JSON output, timings)
- Auto-imports: removes unused imports(ONLY WHEN FILE IS BEING RUN) and adds missing imports

---

## Phase 1: Core & MVP

**Goal:** Get a working, fast tool end-to-end.

### Tasks
- Basic file change detection (mtime+size to start)
- Parallel worker pool wrapping `gofmt` + `goimports`
- Auto-import functionality (leverage `goimports`)
- Atomic writes to files
- `--check` mode (verify without writing)
- JSON output for CI integration
- Basic timing/reporting

### Testing
- Create test Go repo with messy formatting and missing imports
- Verify output correctness against raw `gofmt`/`goimports`
- Time your tool vs running formatters directly
- Test `--check` mode manually
- Test JSON output parsing

### Success Criteria
- Tool runs on a real repo and produces correct formatted output
- Measurable speed improvement over running formatters on everything
- `--check` and JSON modes work reliably

---

## Phase 2: Caching & Invalidation

**Goal:** Skip re-formatting already-processed files, dramatically reducing subsequent runs.

### Tasks
- Persistent cache layer (file hash → formatted result)
- Monitor `go.mod`/`go.sum` for changes to trigger invalidation
- Cache import resolution results (the expensive operation)
- Implement cache eviction policy
- Store cache locally (e.g., `.gofast-cache/`)

### Testing
- Test cache hit/miss scenarios
- Verify cache invalidates when `go.mod` changes
- Benchmark: subsequent runs should be 10x faster

### Success Criteria
- Second run on unchanged repo is near-instant
- Cache correctly invalidates on dependency changes

---

## Phase 3: Intelligent Optimization

**Goal:** Minimize unnecessary work through smart dependency tracking.

### Tasks
- Dependency graph tracking (know which files depend on which)
- Selective invalidation: if file A's imports change, only invalidate files that import A
- Batch `goimports` calls by package to reduce resolver overhead
- Profile and optimize hot paths

### Testing
- Benchmark against large repos
- Test scenarios where one file changes and verify only affected files are re-processed

### Success Criteria
- Handles large codebases efficiently
- Performance scales well with repo size

---

## Phase 4: Polish & Features

**Goal:** Production-ready experience and integration.

### Tasks
- Watch mode (`--watch` flag for continuous formatting)
- Progress bars and better CLI UX
- Improved error messages and logging
- Configuration file support (e.g., `.gofast.yml`)
- Integration tests against real repositories
- Comprehensive benchmark suite
- Content hashing as fallback to mtime+size

### Testing
- Run against multiple large Go projects
- Test all CLI flags and modes
- Verify error handling is robust

### Success Criteria
- Tool feels polished and production-ready
- Good CI/CD integration story

---

## Testing Strategy

### Basic Testing (Phase 1)
- Use Go's `testing` package
- Create `testdata/` directory with intentional format issues
- Compare output byte-for-byte with `gofmt`/`goimports`
- Manual timing: `time ./tool` vs `time gofmt ./...`

### Intermediate Testing (Phase 2+)
- Unit tests for cache logic (hits, misses, invalidation)
- Fixture repos with various scenarios
- Benchmark suite using `testing.B`

### Integration Testing (Phase 3+)
- Clone popular Go projects and test against them
- Measure real-world performance gains
- Test edge cases: empty repos, all files changed, no changes, etc.

---

## Architecture Notes

- **File Detection:** Start with mtime+size, upgrade to content hashing later if needed
- **Worker Pool:** Use buffered channels or goroutine pool pattern
- **Caching:** Simple file-based cache initially (JSON or binary format)
- **Atomicity:** Write to temp file, then atomic rename
- **Output:** JSON mode for CI, human-readable for CLI

---

## Next Steps

1. **Start with Phase 1** — validate the speed gains on a real repository
2. **Measure, don't guess** — use actual timing data to prioritize Phase 2+
3. **Ship incrementally** — Phase 1 alone is useful; add phases as you identify bottlenecks