# Context Map — String Manipulation

## Bounded Context: String DSL
- **Primary responsibility**: UTF-8 aware string splitting, reversing, and transformation
- **Key domain terms**:
  - `Split(s string, sep string)` → UTF-8 byte splitting (for `sep == ""`)
  - `Reverse(s string)` →