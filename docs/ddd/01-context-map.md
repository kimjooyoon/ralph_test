# Context Map

## Bounded Contexts
- **String Manipulation** (domain/split.go, domain/reverse.go, domain/trim.go)
  - Aggregate: StringSegment
  - Invariants: 
    - `Split(s, sep)` must split UTF-8 bytes for multi-byte characters (Chinese, emoji)
    - `Reverse(s)` must handle surrogate pairs and combining marks as single code points
    - `Trim(s)` must preserve valid UTF-8 sequences

- **Numeric Processing** (domain/parseint.go, domain/average.go, domain/range.go)
  - Aggregate: NumericRange
  - Invariants: 
    - `ParseInt(s)` must validate UTF-8 numeric strings
    - `Average(nums)` must floor mean of integers
    - `Range(start, end)` must generate inclusive integer sequences

- **Pattern Matching** (domain/match.go, domain/containswildcard.go)
  - Aggregate: PatternMatcher
  - Invariants: 
    - `Match(pattern, s)` must validate regex-like syntax
    - `ContainsWildcard(s, pattern)` must support * and ? wildcards

- **Encoding/Decoding** (domain/encodebase64.go, domain/decodeimage.go)
  - Aggregate: DataEncoder
  - Invariants: 
    - `EncodeBase64(data