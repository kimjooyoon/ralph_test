# Context Map

## Bounded Contexts

### String Manipulation
- **Aggregate Roots**: `Split`, `Join`, `Trim`, `Repeat`, `Replace`, `Reverse`
- **Domain Events**: 
  - `StringSplitEvent` (when split by UTF-8 bytes)
  - `StringReverseEvent` (when reversed with Unicode awareness)
- **Invariants**:
  - `SplitWithEmptySeparator` (returns UTF-8 bytes, not Unicode code points)