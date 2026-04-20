# Context Map

## Bounded Contexts
- **String Manipulation**: Handles UTF-8 byte splitting, reverse, and Unicode edge cases.  
  - Aggregates: `Split`, `Reverse`, `Trim`, `Join`, `Replace`, `Repeat`, `Echo`  
  - Domain Events: `StringSplit`, `UnicodeReversal`, `Byte