@help:
  just --list --unsorted

# Run tests.
@test:
  go test ./... -coverprofile=cover.out

# Generate documentation.
@doc PORT="6060":
  pkgsite -http=localhost:{{ PORT }} -open
