# opensah

Initial SAH tooling. `opensah` prints the role of each suite name and checks that the built tools are on `PATH`.

**sah** means Service Automation Hub. The hub encompasses the platform and the suite of utilities. This program does not provision machines and it does not open a shell.

```bash
go install ./cmd/opensah
opensah tools
opensah ready
```

`ready` exits 1 when `sahlink`, `sahconnect`, or `sahapi` is missing. A missing `adssh` is printed and does not fail the command. `sahsync` is `optional` (wired, not required) and `sahxfer` is `early` (not wired yet); a missing binary of either is reported and does not fail the check. `sahcore` and `sah2grpc` have no binary yet.

Usage: `docs/usage.md`. Join contract: `~/development/systemapi.io/docs/sah-suite.md`.

## Next

`go test ./...`
