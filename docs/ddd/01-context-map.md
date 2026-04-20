# Context Map — String DSL

## Bounded Contexts
- **String DSL**: Core domain for string manipulation helpers (Split, Reverse, Join, etc.)
  - **Subcontexts**:
    - **Byte Splitting**: Handles UTF-8 byte splitting (e.g., Chinese characters)
    - **Unicode Handling**: Surrogate pairs, combining marks, and code point validation
    - **Pure Functions**: All helpers are stateless and side-effect-free

## Aggregate Boundaries
- Each string helper (Split, Reverse, Join, etc.) is a separate aggregate due to:
  - Pure function