# tsconvert [![Go](https://github.com/nezorflame/tsconvert/actions/workflows/go.yml/badge.svg)](https://github.com/nezorflame/tsconvert/actions/workflows/go.yml) [![Go Report Card](https://goreportcard.com/badge/github.com/nezorflame/tsconvert)](<https://goreportcard.com/report/github.com/nezorflame/tsconvert>) [![GolangCI](https://golangci.com/badges/github.com/nezorflame/tsconvert.svg)](https://golangci.com/r/github.com/nezorflame/tsconvert) [![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2Fnezorflame%2Ftsconvert.svg?type=shield)](https://app.fossa.io/projects/git%2Bgithub.com%2Fnezorflame%2Ftsconvert?ref=badge_shield)

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

## License

[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2Fnezorflame%2Ftsconvert.svg?type=large)](https://app.fossa.io/projects/git%2Bgithub.com%2Fnezorflame%2Ftsconvert?ref=badge_large)
