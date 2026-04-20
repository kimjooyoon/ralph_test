# Context Map

## Bounded Contexts

### String Manipulation
- **Aggregate**: `Split` (UTF-8 byte splitting)
  - **Invariant**: `Split("中文", "")` returns `["中", "文"]` (byte-level splitting)
  - **Domain Events**: `SplitCompleted`, `ByteSplitValidation`
  - **Ubiquitous Language**: 
    - `Split(separator