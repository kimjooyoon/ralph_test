# Spec — minimal DSL helpers (domain)

## Goal

Provide tiny, testable string helpers in `domain` for experiments.

## Acceptance

- `Echo` returns its input unchanged (baseline behavior).
- Helpers remain **pure** (no I/O, no clocks, no network) so they stay fast unit tests.
