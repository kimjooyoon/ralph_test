# IMPLEMENTATION_PLAN — dsl-maker

> Ralph loop: complete unchecked tasks, then mark them `[x]`.

## Tasks

- [x] Bootstrap `Echo` with a domain unit test
- [x] Add `Join(a, b string) string` that returns `a + ":" + b` with tests
- [x] Add `Trim(s string) string` that returns s without leading/trailing whitespace with tests
- [ ] Implement `Split(s, sep string) []string` that splits by UTF-8 bytes (not runes) with tests
- [ ] Fix `split_test.go` to validate byte-based splitting for multi-byte characters (e.g., Chinese)
- [ ] Add `Reverse(s string) string` that handles surrogate pairs and combining marks with tests
- [ ] Add `EncodeBase64(data []byte) string` for base64 encoding with tests
- [ ] Add `EncodeHex(data []byte) string` for hexadecimal encoding with tests

## Questions

- Should `Split` handle empty separator (`sep == ""`) as UTF-8 byte splitting per specs/dsl.md?
- How to validate surrogate pair handling in `Reverse` without I/O (per Unicode Edge Cases)?
- Should `EncodeBase64` and `EncodeHex` be pure functions with no I/O, as per specs?

## specs
--- specs/dsl.md ---
# Spec — minimal DSL helpers (domain)

## Goal

Provide tiny, testable string helpers in `domain` for experiments.

## Acceptance

- `Echo` returns its input unchanged (baseline behavior).
- Helpers remain **pure** (no I/O, no clocks, no network) so they stay fast unit tests.

## Ideation + extension (within unit scope)

Exploration should happen **inside `domain/`** as new tests + small functions.

- Durable intent belongs in this `specs/` directory.
- Short-horizon tasks belong in `IMPLEMENTATION_PLAN.md`.
- “Maybe someday” ideas and brainstorming belong in `docs/ralph/**/*.md` (injected into the loop as extra context).

If an idea cannot be expressed without integration/e2e, it is **out of scope** for this harness loop.

## `Split` empty separator

For `sep == ""` and non-empty `s`, `domain.Split` returns one string per **UTF-8 byte** of `s` (ASCII is one byte per character). This differs from `strings.Split(s, "")`, which yields one segment per **decoded Unicode code point**. Empty `s` yields an empty slice.

## Unicode handling

- `Split` must correctly split UTF-8 bytes for multi-byte characters (e.g., Chinese, emoji).
- `Reverse` must handle surrogate pairs and combining marks correctly.
- Tests must use valid UTF-8 input to avoid invalid sequence errors.

## Extension notes (docs/ralph/)
--- docs/ralph/00-ideation-backlog-ideas.md ---
# Ideation backlog (examples)

These are **non-binding ideas** for future domain-only increments:

- **String ergonomics**: small helpers for trimming/joining/repeating with obvious semantics.
- **Table-driven edge cases**: prefer one behavior per test file when it helps readability.
- **Property-like checks**: where feasible, encode invariants as small pure functions + tests (still domain-only).

When you promote an idea, mirror it into:

1. `specs/*.md` if it becomes durable intent, and
2. `IMPLEMENTATION_PLAN.md` as a `- [ ]` task with a failing test sketch.

--- docs/ralph/10-hot-signals.md ---
# HN Hot Signals → Domain Experiments

## 1. Chinese Literacy Speedrun II: Character Cyclotron
**Signal**: Handling multi-byte Unicode (e.g., Chinese characters)
**Domain experiment**: Test `Split` with UTF-8 byte splitting for Chinese characters
**Next failing test**: Split("中文", "") = ["中", "文"], want ["中", "文"] (current test passes, but needs explicit validation)

## 2. Codex for almost everything
**Signal**: AI-generated code patterns
**Domain experiment**: Add `GenerateCode(pattern string) string` that returns AI-style code snippets
**Next failing test**: GenerateCode("for i in range(5): print(i)") = "for i in range(5): print(i)", want "for i in range(5): print(i)" (pure function, no I/O)

## 3. Average Is All You Need
**Signal**: Data analysis simplicity
**Domain experiment**: Add `Average(nums []int) int` that returns mean (floor)
**Next failing test**: Average([]int{1,2,3,4,5}) = 3, want 3 (current `Repeat`/`Replace` handle numbers, but not averages)

## 4. FIM – Linux framebuffer image viewer
**Signal**: Image processing via string manipulation
**Domain experiment**: Add `DecodeImage(data string) []byte` that simulates image decoding
**Next failing test**: DecodeImage("PNG") = [80, 78, 71, 13, 10, 26, 10], want [80, 78, 71, 13, 10, 26, 10] (pure function, no I/O)

## 5. A Python Interpreter Written in Python
**Signal**: Self-hosted language interpreters
**Domain experiment**: Add `Eval(code string) string` that returns evaluated code (mocked)
**Next failing test**: Eval("1+1") = "2", want "2" (pure function, no I/O)

--- docs/ralph/11-codex.md ---
# Codex for almost everything → Domain Experiments

## 2. Codex for almost everything
**Signal**: AI-generated code patterns  
**Domain experiment**: Add `GenerateCode(pattern string) string` that returns AI-style code snippets  
**Next failing test**: GenerateCode("for i in range(5): print(i)") = "for i in range(5): print(i)", want "for i in range(5): print(i)" (pure function, no I/O)

## 3. Average Is All You Need
**Signal**: Data analysis simplicity  
**Domain experiment**: Add `Average(nums []int) int` that returns mean (floor)  
**Next failing test**: Average([]int{1,2,3,4,5}) = 3, want 3 (current `Repeat`/`Replace` handle numbers, but not averages)

## 4. FIM – Linux framebuffer image viewer
**Signal**: Image processing via string manipulation  
**Domain experiment**: Add `DecodeImage(data string) []byte` that simulates image decoding  
**Next failing test**: DecodeImage("PNG") = [80, 78, 71, 13, 10, 26, 10], want [80, 78, 71, 13, 10, 26, 10] (pure function, no I/O)

## 5. A Python Interpreter Written in Python
**Signal**: