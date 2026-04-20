# Context Map — String Manipulation

## Bounded Contexts
- **String Manipulation**: Handles UTF-8 byte splitting, surrogate pairs, and Unicode normalization.  
  - **Aggregates**:  
    - `Splitter`: Ensures UTF-8 byte splitting for multi-byte characters (e.g., Chinese, emoji).  
    - `UnicodeValidator`: Validates UTF-8 input to avoid invalid sequence errors.  
  - **Domain Events**:  
    - `ByteSplitEvent`: Triggered when a string is split by UTF-8 bytes.  
    - `SurrogatePairEvent`: Triggered when surrogate pairs are detected and treated as single