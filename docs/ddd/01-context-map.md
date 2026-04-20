# Context Map

## Bounded Contexts
- **String Manipulation**
  - Aggregates: `Split`, `Join`, `Reverse`, `Trim`
  - Domain Events: `CharacterSplit`, `StringReversed`, `SubstringTrimmed`
  - Invariants:
    - `Split("中文