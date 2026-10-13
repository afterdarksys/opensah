# opensah

Initial SAH tooling. `opensah` prints the role of each suite name and checks that the built tools are on `PATH`.

**sah** means Service Automation Hub. The hub encompasses the platform and the suite of utilities. This program does not provision machines and it does not open a shell.

```bash
go install ./cmd/opensah
opensah tools
opensah ready
```

`ready` exits 1 when `sahlink`, `sahconnect`, or `sahapi` is missing. A missing `adssh` is printed and does not fail the command. `sahcore` has no binary.

Usage: `docs/usage.md`. Join contract: `~/development/systemapi.io/docs/sah-suite.md`.

## Next

`go test ./...`
