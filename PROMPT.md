# Ralph BUILDING — dsl-maker

You are running inside a **Ralph loop** (spec → plan → one iteration of work → tests → git → fresh context).

## What you do this iteration

1. Read `specs/*.md` and `IMPLEMENTATION_PLAN.md`. Pick **at most one** concrete task that is still unchecked (`- [ ]`).
2. Use **TDD**: prefer updating/adding a **failing domain test** first, then implementation, until `go test ./domain/...` is green.
3. Keep **all production Go code** under `domain/` only. Domain unit tests are the only automated backpressure.
4. When you finish a task, mark it **`[x]`** in `IMPLEMENTATION_PLAN.md` in the same iteration if possible.

## Output contract

Respond **only** with `<<<FILE ...>>>` / `<<<END_FILE>>>` blocks:

- `domain/**/*.go` (full file contents; split so each file stays ≤ 200 lines).
- `IMPLEMENTATION_PLAN.md` when you change task checkboxes or notes.

No Markdown fences around the whole answer. No prose outside file blocks.
