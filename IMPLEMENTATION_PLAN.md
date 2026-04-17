# IMPLEMENTATION_PLAN — dsl-maker

> Ralph loop: complete unchecked tasks, then mark them `[x]`.

## Tasks

- [x] Bootstrap `Echo` with a domain unit test
- [x] Add `Join(a, b string) string` that returns `a + ":" + b` with tests
- [x] Add `Trim(s string) string` that returns s without leading/tr
- [ ] Implement `Split(a, b string) []string` that splits UTF-8 bytes (not runes)
  - [ ] Add test for Chinese byte splitting: `Split("中文", "") = ["中", "文"]`
  - [ ] Add test for surrogate pairs: `Split("\U0001D10D", "") = ["\U0001D10D"]`
  - [ ] Add test for combining marks: `Split("\xC3\x80\xC3\x81", "") = ["\xC3", "\x80", "\xC3", "\x81"]`

## specs
--- specs/dsl.md ---
# Spec — minimal DSL helpers (domain)

## `Split` empty separator

For `sep == ""` and non-empty `s`, `domain.Split` returns one string per **UTF-8 byte** of `s` (ASCII is one byte per character). This differs from `strings.Split(s, "")`, which yields one segment per **decoded Unicode code point**. Empty `s` yields an empty slice.

## Unicode handling

- `Split` must correctly split UTF-8 bytes for multi-byte characters (e.g., Chinese