# ping-as — reference plugin (AssemblyScript)

Same reference plugin as `ping` (Rust), written in AssemblyScript instead — proves
`plugins.PluginHost`'s ABI is genuinely language-agnostic, not accidentally Rust-shaped. `greet`
proves host-to-guest string calls through guest memory; `fetch_allowed` proves the guest-to-host
permission-scoped HTTP fetch (`env.http_fetch`).

The [kosmos](https://github.com/kosmos-suite/kosmos) repo's `PluginHostAssemblyScriptTest`
(`src/test/java/de/oppahansi/kosmos/plugins/`) loads the compiled output — checked in there at
`src/test/resources/plugins/ping-as/ping-as.wasm` alongside its `manifest.json`. Not part of any
Gradle build; there's no Node/AssemblyScript toolchain wired into that repo.

Built on
[kosmos-plugin-sdk-assemblyscript](https://github.com/kosmos-suite/kosmos-plugin-sdk-assemblyscript)
for the ABI scaffolding — `alloc`/`dealloc`, `httpFetch`, string pack/unpack.

## Rebuilding

```shell
npm install
npm run asbuild
cp build/ping-as.wasm ../../kosmos/src/test/resources/plugins/ping-as/ping-as.wasm
```
