# IMPLEMENTATION_PLAN — dsl-maker

> Ralph loop: complete unchecked tasks, then mark them `[x]`.

## Tasks

- [x] Bootstrap `Echo` with a domain unit test
- [x] Add `Join(a, b string) string` that returns `a + ":" + b` with tests
- [x] Add `Trim(s string) string` that returns s without leading/trailing whitespace with tests
- [x] Add `Split(s, sep string) []string` that splits strings with tests
- [x] Improve `Split` to handle multi-character separators
- [x] Add `Repeat(s string, n int) string` that returns s repeated n times with tests
- [x] Add `Replace(s, old, new string) string` that returns s with all occurrences of old replaced by new with tests
- [x] Add test for `Split` with empty separator and handle it
- [x] Add `Reverse(s string) string` that returns s reversed with tests
- [ ] Add `IsPalindrome(s string) bool` that returns true if s is a palindrome with tests


## Questions

- Should `Split` with empty separator return each character as a separate element (current behavior) or follow `strings.Split` semantics (which returns empty slice for empty separator)?
- Are there additional edge cases for `Reverse` that should be covered (e.g., Unicode characters, surrogate pairs)?

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

--- docs/ralph/README.md ---
# `docs/ralph/` — ideation fuel (dsl-maker)

Markdown files in this directory tree are **appended to every Ralph iteration** (sorted by path).

Use this space for **ideas that must still compile into pure `domain/` work**:

- candidate helpers and API shapes,
- edge cases worth encoding as tests,
- naming and ergonomics notes,
- “next experiment” bullets.

Do **not** use this to bypass the unit-only policy (no IO/network/time dependencies in `domain/`).

Suggested naming:

- `00-...` early ideation / backlog
- `10-...` more concrete experiments
- keep each file short (prefer many small files over one huge file)

## go test ./domain/... output
ok  	github.com/kimjooyoon/ralph-tdd/projects/dsl-maker/domain	0.176s


## domain Go sources (read-only context)
--- domain/dsl.go ---
package domain

// Placeholder domain logic for Ralph TDD bootstrap.
func Echo(s string) string {
	return s
}

--- domain/dsl_test.go ---
package domain

import "testing"

func TestEcho(t *testing.T) {
	if got := Echo("hi"); got != "hi" {
		t.Fatalf("Echo = %q, want %q", got, "hi")
	}
}

--- domain/join.go ---
package domain

// Join joins two strings with ":".
func Join(a, b string) string {
	return a + ":" + b
}

--- domain/join_test.go ---
package domain

import "testing"

func TestJoin(t *testing.T) {
	if got := Join("a", "b"); got != "a:b" {
		t.Fatalf("Join = %q, want a:b", got)
	}
}

--- domain/repeat.go ---
package domain

import "strings"

// Repeat returns s repeated n times.
func Repeat(s string, n int) string {
	if n <= 0 {
		return ""
	}
	return strings.Repeat(s, n)
}

--- domain/repeat_test.go ---
package domain

import "testing"

func TestRepeat(t *testing.T) {
	tests := []struct {
		input string
		n     int
		want  string
	}{
		{"a", 3, "aaa"},
		{"hello", 2, "hellohello"},
		{"", 5, ""},
		{"x", 0, ""},
		{"y", -1, ""},
	}
	for _, tt := range tests {
		if got := Repeat(tt.input, tt.n); got != tt.want {
			t.Errorf("Repeat(%q, %d) = %q, want %q", tt.input, tt.n, got, tt.want)
		}
	}
}

--- domain/replace.go ---
package domain

import "strings"

// Replace returns s with all occurrences of old replaced by new.
func Replace(s, old, new string) string {
	return strings.ReplaceAll(s, old, new)
}

--- domain/replace_test.go ---
package domain

import "testing"

func TestReplace(t *testing.T) {
	tests := []struct {
		input  string
		old    string
		new    string
		want   string
	}{
		{"hello", "l", "x", "hexxo"},
		{"aabbcc", "aa", "xx", "xxbbcc"},
		{"", "x", "y", ""},
		{"abc", "d", "x", "abc"},
		{"aaaa", "aa", "bb", "bbbb"},
	}
	for _, tt := range tests {
		if got := Replace(tt.input, tt.old, tt.new); got != tt.want {
			t.Errorf("Replace(%q, %q, %q) = %q, want %q", tt.input, tt.old, tt.new, got, tt.want)
		}
	}
}

--- domain/reverse.go ---
package domain

// Reverse returns s reversed.
func Reverse(s string) string {
	runes := []rune(s)
	for i := 0; i < len(runes)/2; i++ {
		runes[i], runes[len(runes)-1-i] = runes[len(runes)-1-i], runes[i]
	}
	return string(runes)
}

--- domain/reverse_test.go ---
package domain

import "testing"

func TestReverse(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{"abc", "cba"},
		{"", ""},
		{"a", "a"},
		{"hello world", "dlrow olleh"},
		{"12345", "54321"},
	}
	for _, tt := range tests {
		if got := Reverse(tt.input); got != tt.want {
			t.Errorf("Reverse(%q) = %q, want %q", tt.input, got, tt.want)
		}
	}
}

--- domain/split.go ---
package domain

import "strings"

// Split splits s by sep (mostly strings.Split semantics, except Split("", sep) returns an empty slice).
func Split(s, sep string) []string {
	if s == "" {
		return []string{}
	}
	if sep == "" {
		result := make([]string, len(s))
		for i := 0; i < len(s); i++ {
			result[i] = string(s[i])
		}
		return result
	}
	return strings.Split(s, sep)
}

--- domain/split_test.go ---
package domain

import "testing"

func TestSplit(t *testing.T) {
	tests := []struct {
		input  string
		sep    string
		want   []string
	}{
		{"a:b:c", ":", []string{"a", "b", "c"}},
		{"one,two,three", ",", []string{"one", "two", "three"}},
		{"", ":", []string{}},
		{"hello", "x", []string{"hello"}},
		{"a,,b", ",", []string{"a", "", "b"}},
		{"aabbccbbaa", "ab", []string{"a", "bccbbaa"}},
		{"abc", "", []string{"a", "b", "c"}},
	}
	for _, tt := range tests {
		if got := Split(tt.input, tt.sep); !equal(got, tt.want) {
			t.Errorf("Split(%q, %q) = %v, want %v", tt.input, tt.sep, got, tt.want)
		}
	}
}

func equal(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

--- domain/trim.go ---
package domain

import "strings"

func Trim(s string) string {
	return strings.TrimSpace(s)
}

--- domain/trim_test.go ---
package domain

import "testing"

func TestTrim(t *testing.T) {
	if got := Trim("  abc  "); got != "abc" {
		t.Errorf("Trim = %q, want %q", got, "abc")
	}
}
