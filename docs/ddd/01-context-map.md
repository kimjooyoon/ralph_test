# Context Map

## Bounded Contexts
- **String Handling** (domain/split.go, domain/trim.go, domain/join.go)
  - Aggregate: `Split` with UTF-8 byte splitting for multi-byte characters
  - Invariant: `Split("中文", "")` returns ["中", "文"]
  - Domain Events: Character splitting, byte iteration

- **AI Code Generation** (domain/generatecode.go)
  - Aggregate: `GenerateCode` for AI-style code snippets
  - Invariant: `GenerateCode("for i in range(5): print(i)")` returns valid Python code
  - Domain Events: Code pattern recognition,