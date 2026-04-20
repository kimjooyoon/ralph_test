# Context Map

## Bounded Contexts
- **String Manipulation**: Handles UTF-8 byte splitting, reverse, trim, join, etc. (domain/split.go, domain/reverse.go)
- **AI Code Generation**: Generates AI-style code snippets (domain/generate_code.go)
- **Unicode Handling**: Special focus on surrogate pairs, multi-byte characters (domain/split_test.go)
- **Pure Functions**: All helpers are stateless, no I/O (specs/dsl.md)

## Aggregate Boundaries