# ping-go — reference plugin (Go / TinyGo)

Same reference plugin as `ping` (Rust) and `ping-as` (AssemblyScript), written in Go and compiled
via TinyGo — proves `plugins.PluginHost`'s ABI generalizes across a third guest language. `greet`
proves host-to-guest string calls through guest memory; `fetch_allowed` proves the guest-to-host
permission-scoped HTTP fetch (`env.http_fetch`).

The [kosmos](https://github.com/kosmos-suite/kosmos) repo's `PluginHostGoTest`
(`src/test/java/de/oppahansi/kosmos/plugins/`) loads the compiled output — checked in there at
`src/test/resources/plugins/ping-go/ping-go.wasm` alongside its `manifest.json`. Not part of any
Gradle build; there's no Go/TinyGo toolchain wired into that repo.

Built on [kosmos-plugin-sdk-go](https://github.com/kosmos-suite/kosmos-plugin-sdk-go) for the ABI
scaffolding — `alloc`/`dealloc`, `HTTPFetch`, string pack/unpack.

## Rebuilding

```shell
mkdir -p build
tinygo build -o build/ping-go.wasm -target=wasm-unknown -gc=leaking -no-debug -opt=z .
cp build/ping-go.wasm ../../kosmos/src/test/resources/plugins/ping-go/ping-go.wasm
```
