# AGENTS.md

opensah is the local index of the SAH groundwork. It is not a platform.

Commands: `docs/usage.md`. Cross-repo contract: `~/development/systemapi.io/docs/sah-suite.md`.

## Do this

1. Read `docs/usage.md` before changing a command.
2. Change a role only in `internal/suite/suite.go`, and update `suite_test.go` in the same edit. Done when `go test ./...` reports ok and the sahcore line is still `not_built`.

## Rules

- Keep sahcore `not_built` / `unassigned` until a role is assigned. Shipping a sahcore program invents a product.
- Do not shell out to adssh, systemapi, or computeapi from this module. Execution belongs to those tools. `ready` is a `LookPath` check.
- Do not rename `adssh`, module `github.com/afterdarksys/adssh`, or repo `afterdarksys/adssh`. The brand word is sah (Service Automation Hum). The binary name stays.
- When a join becomes real, edit the Wired join list in the suite contract in the same change. A sentence here does not make it real.

## Next

`go test ./...`
