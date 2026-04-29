# test

[![CI](https://github.com/clever-vpn/test/actions/workflows/ci.yml/badge.svg?branch=main)](https://github.com/clever-vpn/test/actions/workflows/ci.yml)

A small Go project for learning a GitHub-based development workflow.

## Project layout

- `cmd/server`: HTTP service entrypoint
- `internal/calc`: domain logic and unit tests
- `tests/integration`: integration tests
- `.github/workflows`: CI pipeline
- `.github/ISSUE_TEMPLATE`: issue forms

## Quick start

```bash
go test ./...
go run ./cmd/server
curl "http://localhost:8080/sum?numbers=1,2,3"
```

Expected response:

```json
{"total":6}
```

## Learning guide

Follow `CONTRIBUTING.md` and `docs/WORKFLOW_TRAINING.md` for full PR and issue practice.

test1
test2
fork pr training
