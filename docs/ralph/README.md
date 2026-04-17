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
