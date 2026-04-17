# Agent operations — dsl-maker

## Commands

- Run domain unit tests: `go test ./domain/... -count=1 -short`
- Module root: this directory (`go.mod` lives here).

## TDD Ralph expectations

- This project is driven by a **perpetual TDD Ralph loop** (no automatic stop when the plan is complete).
- Prefer **one behavioral change per iteration**, anchored by a domain test.
- If you are “done” with the plan, pull the next increment from `specs/*.md` into `IMPLEMENTATION_PLAN.md` and continue with RED tests.
- Optional ideation context is loaded from `docs/ralph/**/*.md` (see `docs/harness-extensions.md` at the **repository root**).
- When BUILDING stagnates (no effective FILE updates), the next iteration runs **PLANNING Ralph** using `PROMPT_PLANNING.md` (markdown-only writes: plan/specs/docs/ralph).
- The harness **reads files for you** each iteration and prepends a **Harness-loaded files** index (also saved to `.ralph/logs/last_file_index.txt`). The model does not get silent repo-wide file access.
- **PROPOSER Ralph** runs after **PLANNING** makes no disk changes: it pulls a small public “hot tech” snapshot (default: Hacker News titles/links) and writes proposals **only** under `docs/ralph/*.md`.

## Repo rules

- **No e2e / integration test packages** in this loop; only `domain` tests drive the harness.
- **File size**: each `domain/*.go` file must stay ≤ 200 lines; split packages/files instead of growing one file.
- **Plan file**: `IMPLEMENTATION_PLAN.md` lives at module root. Use Markdown task list syntax (`- [ ]` / `- [x]`).

## LLM output format

The harness parses **full-file blocks** only:

```text
<<<FILE path="domain/foo.go">>>
package domain
<<<END_FILE>>>
```

(Use `>>>` preferred after the path; `>>` is tolerated.)
