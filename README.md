# ATHENA: Go Formatter & Import-Fixer (Experimental)

> **Note:** This is an experimental project under active development. Please do not submit merge or commit requests at this time.

## Project Overview
ATHENA is a fast, incremental formatter and import-fixer for Go. It processes only changed files, runs formatters in parallel, and uses intelligent caching for speed and efficiency.

## Phase 1 Goals
- Detect changed files using modification time and size
- Run `gofmt` and `goimports` only on changed files
- Execute formatters in parallel using a worker pool
- Provide auto-import functionality (leveraging `goimports`)
- Write changes atomically to files
- Support a `--check` mode to verify formatting without writing
- Output results in JSON for CI integration
- Report basic timing and concise results

## Status
Currently in Phase 1: building core functionality and validating speed improvements on real repositories.

---
For more details, see `PhaseGoals.md`.
