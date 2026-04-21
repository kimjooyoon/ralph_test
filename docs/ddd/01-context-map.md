# Context Map

## Bounded Contexts
- **String Manipulation**
  - Aggregate Roots: `Split`, `Reverse`, `Trim`, `Join`, `Repeat`, `Replace`
  - Invariants:
    - `Split("中文", "")` → ["中", "文"] (UTF-8 byte splitting)
    - `Reverse("中文")` → "文中" (surrogate pairs handled)
    - `Trim("  abc  ")` → "abc" (whitespace removal)
    - `Join(["a", "b"], "-")` → "a-b"
    - `Repeat("a", 3)` → "aaa"
    - `Replace("abc", "a", "x")` → "xbc"

- **AI Code Generation**
  - Aggregate Roots: `GenerateCode`
  - Invariants:
    - `GenerateCode("for i in range(5): print(i)")` → same input (pure function