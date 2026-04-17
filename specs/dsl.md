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
