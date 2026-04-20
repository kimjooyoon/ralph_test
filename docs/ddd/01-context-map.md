# Context Map

## Bounded Contexts
- **String Manipulation** (primary context)
  - Aggregates: String Processing, Unicode Handling
  - Domain Events: Character Split, Reverse, Encode
  - Invariants:
    - `Split("中文", "")` returns ["中", "文"] (UTF-8 byte split)
    - `Reverse` handles surrogate pairs as single code points
    - `EncodeBase64` produces valid base64 for non-ASCII input

- **Code Generation** (secondary context)
  - Aggregates: AI Pattern Matching, Template Rendering
  - Domain Events: Code Snippet Generation, Eval Simulation
  - Invariants:
    - `GenerateCode("for i in range(5): print(i)")` returns valid Python
    - `Eval("1+1")` returns "2" (mocked execution)

## Context Boundaries
- String operations are confined to `domain/` package
- Unicode handling is bounded to `Split`, `Reverse`, and `EncodeBase64`
- Code