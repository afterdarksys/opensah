# opensah usage

Module `github.com/afterdarksys/opensah`. Version 0.1.0.

## Install

1. `go install ./cmd/opensah` from this directory. Done when `opensah version` prints `opensah 0.1.0`.

## Commands

1. `opensah tools` prints TSV: `name`, `state`, `role`. `sahcore` is `not_built` and `unassigned`. `sahsync` is `optional`. `sahxfer` and `sah2grpc` are `early`.
2. `opensah ready` checks `PATH`. Built tools `sahlink`, `sahconnect`, and `sahapi` must resolve. `adssh` is reported. `sahsync` (`optional`) and `sahxfer` (`early`) are reported as missing without failing. `sahcore`, `sah2grpc`, `systemapi`, and `computeapi` have no binary in this check.
3. `opensah version` prints the version. An unknown command exits 2.

`ready` does not start the other tools and does not call their APIs.

## Next

`opensah tools`
