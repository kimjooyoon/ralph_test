# Context Map

## Bounded Contexts
- **StringHelpers**: Handles string manipulation (Split, Reverse, Trim, Join)
  - Aggregate: String
  - Invariants: 
    - Split must split UTF-8 bytes for multi-byte characters (e.g., Chinese)
    - Reverse must handle surrogate pairs and combining marks
- **NumericOps**: Numeric range generation and conversion
  - Aggregate: NumberRange
  - Invariants: 
    - Range(start, end) returns inclusive sequence
    - ParseInt(s