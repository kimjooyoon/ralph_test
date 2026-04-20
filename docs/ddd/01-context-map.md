# Context Map (v2)

## Bounded Contexts
- **String Manipulation** (primary context)
  - `Split`, `Join`, `Trim`, `Reverse`, `EncodeBase64`, `EncodeHex`
  - `Match`, `ContainsWildcard`, `Range`, `ParseInt`
- **Data Analysis** (secondary context)
  - `Average`, `GenerateCode`, `Eval`, `DecodeImage`
- **Unicode Handling** (cross-cutting concern)
  - `Split` with UTF-8 byte splitting
  - `Reverse` with surrogate pairs
  - `TestSplitSurrogatePairs`

## Aggregate Boundaries
- `Split` aggregates: `s` (string), `sep` (separator)
- `Reverse` aggregates: `s` (string)
- `EncodeBase64` aggregates: `data` ([]byte)

## Domain Events
- `StringSplitEvent` (when `Split` completes)
- `StringReversedEvent` (when `Reverse` completes)
- `EncodingCompletedEvent` (when base64/hex encoding completes)

## Invariants
- `Split` must split UTF-8 bytes, not Unicode code points
- `Reverse` must preserve surrogate pairs as single code points
- `Encode