# Context Map — String Manipulation

## Bounded Context: String DSL
- **Domain**: String manipulation with UTF-8 awareness
- **Boundaries**:
  - `Split` handles UTF-8 byte splitting for multi-byte characters
  - `Reverse` processes surrogate pairs as single code points
  - `Trim` removes whitespace per Unicode definition
- **Events**:
  - `StringSplitEvent`: UTF-8 byte splitting completed
  - `StringReversedEvent`: Surrogate pairs processed correctly
- **Invariants**:
  - `Split("", "")` returns empty slice
  - `Split("中文", "")` returns ["中", "文"]
  - Surrogate pairs treated as single code points
- **Shared Language**:
  - "UTF-8 byte" vs "Unicode code point"
  - "Surrogate pair" as single code point
  - "Empty separator" yields