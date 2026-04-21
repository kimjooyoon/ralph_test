# Context Map — String Manipulation Domain

## Bounded Context: String Processing
- **Aggregate Roots**: `Split`, `Reverse`, `EncodeBase64`
- **Domain Events**: 
  - `StringSplitCompleted` (when `Split` finishes processing)
  - `StringReversed` (when `Reverse` completes)
- **Invariants**:
  - `Split(s, "")` returns UTF-8 byte array for non-empty `s` (e.g., "中文" → ["中", "文"])
  - `Reverse` handles surrogate pairs as single code points (e.g., emoji)
  - `EncodeBase64` produces valid base64 for any byte array

## Bounded Context