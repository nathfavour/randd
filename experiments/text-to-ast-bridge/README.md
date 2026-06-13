# Experiment: Text-to-AST Hybrid Compiler Bridge

## Objective
Benchmark three code mutation strategies for LLMs:
1. **Pipeline A (Direct Text):** Search-and-replace / Diff.
2. **Pipeline B (Direct AST):** Mutating raw JSON AST.
3. **Pipeline C (Hybrid Bridge):** Intent-based DSL mapped to AST operations locally.

## Setup
- **Language:** Go
- **Tools:** `go/ast`, `go/parser`, `go/format`
- **Subject:** `test-subjects/simple_service.go`

## Target Mutation
Task: Modify `Greeter(name string)` to `Greeter(name string, title string)` and update the return string to include the title.
