# Contributing Workflow (Practice Version)

This repository is intentionally designed for practicing GitHub workflow skills.

## 1) Create an issue first

- Open a bug or feature issue from templates.
- Write clear acceptance criteria.
- Assign labels (`bug`, `enhancement`, `good first issue`).

## 2) Create a branch

Recommended branch naming:

- `feat/<issue-id>-short-name`
- `fix/<issue-id>-short-name`
- `chore/<issue-id>-short-name`

Example:

```bash
git checkout -b feat/12-add-subtract-endpoint
```

## 3) Develop with tests

- Add or update tests first when possible.
- Run:

```bash
go test ./...
```

## 4) Commit with clear messages

Examples:

```bash
git add .
git commit -m "feat: add subtract API"
git push -u origin feat/12-add-subtract-endpoint
```

## 5) Open a Pull Request

- Use the PR template.
- Link issue with `Closes #<id>`.
- Explain what changed and how you tested.

## 6) Review and update

- Reviewer leaves comments.
- Author pushes follow-up commits.
- Re-run CI and resolve all comments.

## 7) Merge strategy

For learning, use **Squash and merge** to keep history clean.

- PR title becomes commit title on `main`.
- Delete branch after merge.

## 8) Post-merge checks

- Pull latest `main` locally.
- Verify CI on `main` is green.
- Close/confirm linked issue is resolved.
