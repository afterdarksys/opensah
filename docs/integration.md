# opensah integration

opensah does not import adssh, systemapi, or computeapi. `ready` only calls `exec.LookPath`.

The role table is `internal/suite/suite.go`. Change a role there and in `suite_test.go` together. The suite-wide contract is `~/development/systemapi.io/docs/sah-suite.md`.

Leave `sahcore` as `not_built`. A binary would invent a product that has no assigned role.

## Next

`go test ./internal/suite/`
