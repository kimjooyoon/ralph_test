# Context Map

## Bounded Contexts

1. **String Manipulation**
   - Handles UTF-8 byte sequences, surrogate pairs, and Unicode code points
   - Contains: Split, Reverse, EncodeBase64, Match, Replace, Join, Trim, Repeat

2. **Pattern Matching**
   - Supports regex-like wildcards (* and ?) for substring searches
   - Contains: Match, ContainsWildcard

3. **Data Encoding**
   - Provides base64 and hexadecimal encoding/decoding
   - Contains: EncodeBase64, EncodeHex

4. **Numeric Operations**
   - Generates ranges and converts strings to integers
   - Contains: Range, ParseInt

## Aggregate Boundaries

- **String Manipulation Aggregate**
  - Ensures UTF-8 correctness for multi-byte characters
  - Invariants: 
    - Split must split on UTF-8 bytes, not code points
    - Reverse must treat surrogate pairs as single code points
    - EncodeBase64 must handle UTF-8 input correctly

- **Pattern Matching Aggregate**
  - Enforces regex-like wildcard matching semantics
  - Invariants: 
    - Match must validate pattern syntax
    - ContainsWildcard must handle * and ? wildcards

## Domain Events

- `