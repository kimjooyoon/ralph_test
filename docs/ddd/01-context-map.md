# Context Map — dsl-maker

## Bounded Contexts
- **String Manipulation**: All domain functions (`Join`, `Split`, `Trim`, `Reverse`, etc.) belong to this context. They are pure, testable helpers with no external dependencies.

## Aggregate Boundaries
- No aggregates (all functions are value objects with no persistence).

## Domain Events
- None (all behavior is stateless).

## Ubiquitous Language
| Term         | Definition                                                                 | Example                     |
|--------------|----------------------------------------------------------------------------|-----------------------------|
| `Join`       | Concatenate two strings with `:`                                          | `Join("a", "b") = "a:b"`   |
| `Split`      | Split by `sep`; empty `sep` → one segment per **UTF-8 byte** (see `specs/dsl.md`) | `Split("abc", "") = ["a","b","c"]` |
| `Trim`       | Remove leading/trailing whitespace                                        | `Trim("  abc  ") = "abc"` |
| `Reverse`    | Reverse string with Unicode-aware handling (surrogate pairs, Unicode characters) | `Reverse("été") = "été"`   |
| `IsPalindrome` | Same forward/backward in **runes** (via `Reverse`) | `IsPalindrome("abcba")`, `IsPalindrome("été")` |

## Resolved
- **`Split` empty `sep`**: one element per UTF-8 byte; differs from `strings.Split` (rune grouping). Implementation uses byte slices `s[i:i+1]`.
- **`Reverse`**: supplementary plane and `été` covered in `domain/reverse_test.go`.

## Open questions
- **UTF-8 byte splitting**: Current `Split` implementation slices by byte index, but this breaks multi-byte sequences (e.g., "ÀÁ" → [� � � �]). Should we instead split by **runes** (Unicode code points) when `sep == ""`? This would align with `strings.Split` behavior and avoid invalid UTF-8 sequences. The test failure `Split("ÀÁ", "") = [� � � �], want [À Á]` demonstrates this issue.

## go test ./domain/... output
# github.com/kimjooyoon/ralph-tdd/projects/dsl-maker/domain [github.com/kimjooyoon/ralph-tdd/projects/dsl-maker/domain.test]
domain/split_test.go:17:13: undefined: Split
domain/split.go:3:8: "strings" imported and not used
FAIL	github.com/kimjooyoon/ralph-tdd/projects/dsl-maker/domain [build failed]
FAIL
