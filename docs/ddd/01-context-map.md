# Context Map

## Bounded Contexts
- **String Manipulation Utilities** (primary context)
  - Handles UTF-8 byte splitting, surrogate pairs, combining marks
  - Aggregates: `Split`, `Reverse`, `EncodeBase64`
  - Subcontexts: 
    - `UTF8Decoder` (handles byte-to-code-point conversion)
    - `SurrogatePairHandler` (specializes in emoji/extended characters)
    - `ChineseCharacterProcessor` (specific to East Asian text)

## Aggregate Boundaries
- `Split(s, sep)`:
  - **Invariant**: When `sep == ""`, splits by UTF-8 byte (not code point)
  - **Domain Events**: 
    - `ByteSplitStarted` (when splitting begins)
    - `ByteSplitCompleted` (when all bytes processed)
- `Reverse(s)`:
  - **Invariant**: Preserves surrogate pairs as single code points
  - **Domain Events**: 
    - `ReversalStarted`
    - `ReversalCompleted`
- `EncodeBase64(data)`:
  - **Invariant**: Produces valid Base64 for any byte sequence
  - **Domain Events**: 
    - `EncodingStarted`
    - `EncodingCompleted`

## Ubiquitous Language
- **Byte** - UTF-8 encoded byte (not code point)
- **Code Point** - Unicode scalar value (e.g., "中" is U+4E2D)
- **Surrogate Pair** - Two 16-bit code units representing a single code point (e.g., emoji)
- **Combining Mark** - Diacritic or modifier that combines with a base character (e.g., "👨