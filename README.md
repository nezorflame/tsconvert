# tsconvert [![go](https://github.com/nezorflame/tsconvert/actions/workflows/go.yml/badge.svg)](<https://github.com/nezorflame/tsconvert/actions/workflows/go.yml>) [![golangci-lint](https://github.com/nezorflame/tsconvert/actions/workflows/golangci-lint.yml/badge.svg)](<https://github.com/nezorflame/tsconvert/actions/workflows/go.yml>) [![Go Report Card](https://goreportcard.com/badge/github.com/nezorflame/tsconvert)](<https://goreportcard.com/report/github.com/nezorflame/tsconvert>)

Converts Unix timestamp to human-readable form and vice versa.

## Installation

This project uses Go modules.
To install it, just use `go install`:

`go install github.com/nezorflame/tsconvert@latest`

## Usage

```text
-f string
    Format of the time (default "time.RFC3339")
-l string
    Location of the time (default "UTC")
-local
    Use current user location
-utc
    Use UTC location
```
