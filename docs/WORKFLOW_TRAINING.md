# GitHub Workflow Training Script

Use this guide to practice a complete cycle with me.

## Goal

Learn all key steps:

- create issue
- create branch
- modify code
- add tests
- push and open PR
- review PR
- merge PR
- track issue status

## Practice Round 1: Feature

### Step A: Create issue

Create a feature issue: "Add subtract function and /subtract API".

Acceptance criteria:

- add `Subtract(a, b int) int` in `internal/calc`
- add unit tests for subtract
- add `/subtract?a=10&b=4` endpoint returning JSON
- all CI checks pass

### Step B: Create branch

```bash
git checkout -b feat/<issue-id>-add-subtract
```

### Step C: Implement

- update `internal/calc`
- update `cmd/server/main.go`
- add tests

### Step D: Verify locally

```bash
go test ./...
```

### Step E: Commit and push

```bash
git add .
git commit -m "feat: add subtract function and endpoint"
git push -u origin feat/<issue-id>-add-subtract
```

### Step F: Open PR

PR title example:

`feat: add subtract function and endpoint`

PR body:

- what changed
- why
- test evidence
- `Closes #<issue-id>`

### Step G: Review simulation (with me)

Ask me: "Please review this PR diff".
I will respond in code review mode with:

- critical findings first
- then medium/low risk findings
- missing tests or edge cases

### Step H: Merge

After approvals and green CI:

- squash and merge
- delete branch

### Step I: Sync local main

```bash
git checkout main
git pull --ff-only
```

## Practice Round 2: Bugfix

Create a bug issue for invalid query handling in `/sum` and repeat the same flow.

## Review checklist

- issue has clear acceptance criteria
- PR links issue
- CI is green
- tests cover new behavior
- reviewer comments are resolved
