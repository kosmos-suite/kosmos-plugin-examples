# ping — reference plugin

Not a real metadata source. Exists to exercise `plugins.PluginHost`'s ABI end to end —
`greet` proves host-to-guest string calls through guest memory; `fetch_allowed` proves the
guest-to-host permission-scoped HTTP fetch (`env.http_fetch`, allowed/blocked by the loading
manifest's `permissions.allowedHosts`). See `src/lib.rs` for the ABI itself.

The [kosmos](https://github.com/kosmos-suite/kosmos) repo's `PluginHostTest`
(`src/test/java/de/oppahansi/kosmos/plugins/`) loads the compiled output — checked in there at
`src/test/resources/plugins/ping/ping.wasm` alongside its `manifest.json`. Not part of any Gradle
build; there's no Rust toolchain wired into that repo.

Built on [kosmos-plugin-sdk-rust](https://github.com/kosmos-suite/kosmos-plugin-sdk-rust) for the
ABI scaffolding — `alloc`/`dealloc`, `http_fetch`, string pack/unpack.

## Rebuilding

```shell
rustup target add wasm32-unknown-unknown   # once
cargo build --release --target wasm32-unknown-unknown
cp target/wasm32-unknown-unknown/release/ping.wasm \
  ../../kosmos/src/test/resources/plugins/ping/ping.wasm
```
