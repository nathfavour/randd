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

## Initial Benchmark Results (Phase 1)
| Pipeline | Mutation Payload (Tokens) | Relative Efficiency |
| :--- | :--- | :--- |
| **Pipeline A (Text)** | 63 | 1.0x (Baseline) |
| **Pipeline B (AST JSON)** | 85 | 0.74x |
| **Pipeline C (Hybrid)** | 25 | **2.52x** |

### Observations
- **Token Economics:** Pipeline C reduces token consumption by over 60% compared to traditional text search/replace.
- **Reliability:** Pipeline C ensures syntax correctness by performing mutations at the AST level, whereas Pipeline A is prone to breaking formatting or missing edge cases.
- **JSON Overhead:** Pipeline B (Raw AST JSON) is the most expensive due to structural boilerplate (brackets, keys, nesting).
