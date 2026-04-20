# Context Map

## Bounded Contexts
- **String Manipulation**: Handles UTF-8 aware splitting, reversing, and encoding.  
  - **Aggregates**:  
    - `Split` (UTF-8 byte splitting, surrogate pair handling)  
    - `Reverse` (Unicode-compliant reversal)  
    - `Join` (UTF-8 aware concatenation)  
  - **Domain Events**:  
    - `StringSplitEvent` (triggered on successful UTF-8 splitting)  
    - `StringReverseEvent` (triggered on successful Unicode reversal)  

## Ubiquitous Language
- **Split**: UTF-8 byte splitting for multi-byte characters (