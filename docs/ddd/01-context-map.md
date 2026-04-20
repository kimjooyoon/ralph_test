# Context Map

## Bounded Contexts
- **String Manipulation**: Handles UTF-8 byte splitting, reversing, and basic string operations.  
  - Aggregate: `Split` (splits strings into UTF-8 byte segments)  
  - Invariant: `Split("中文", "")`