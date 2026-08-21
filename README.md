# kosmos-plugin-examples

Reference plugins for [Kosmos](https://github.com/kosmos-suite/kosmos)'s WASM plugin host, one per
guest language/SDK — pure ABI checks, not real metadata sources. (The one real, production-used
plugin, `tmdb-search`, has its own repo:
[kosmos-plugin-tmdb-search](https://github.com/kosmos-suite/kosmos-plugin-tmdb-search).)

| Plugin | Language | SDK | What it proves |
|---|---|---|---|
| [`ping`](ping) | Rust | [kosmos-plugin-sdk-rust](https://github.com/kosmos-suite/kosmos-plugin-sdk-rust) | The base ABI — guest-memory string round trips, permission-scoped `http_fetch` |
| [`ping-as`](ping-as) | AssemblyScript | [kosmos-plugin-sdk-assemblyscript](https://github.com/kosmos-suite/kosmos-plugin-sdk-assemblyscript) | Same ABI, second language |
| [`ping-go`](ping-go) | Go (TinyGo) | [kosmos-plugin-sdk-go](https://github.com/kosmos-suite/kosmos-plugin-sdk-go) | Same ABI, third language |

Each plugin's own README has its exact rebuild steps. Every one of them depends on its SDK via a
local path/`replace`/`file:` reference assuming this repo is checked out as a sibling of the SDK
repos and of [kosmos](https://github.com/kosmos-suite/kosmos) itself — the layout the
`kosmos-suite` organization's repos are meant to be cloned into side by side.

Compiled `.wasm` output isn't published from here — each plugin's README says exactly where the
built binary needs to land in the `kosmos` repo (test fixtures for `ping`/`ping-as`/`ping-go`).
