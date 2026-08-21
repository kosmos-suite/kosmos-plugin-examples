//! Minimal reference plugin for Kosmos's WASM plugin host (`plugins.PluginHost`, Java side).
//! Not a real metadata source — this exists purely to exercise the ABI end to end:
//! host-to-guest string calls (`greet`), and a guest-to-host permission-scoped HTTP fetch
//! (`fetch_allowed`), so `PluginHostTest` has a real, non-stub `.wasm` binary to load and run.
//!
//! ABI details (pointer/length packed into i64, `alloc`/`dealloc`, `http_fetch`) live in
//! `kosmos-plugin-sdk` — see that crate for the full explanation.

kosmos_plugin_sdk::export_memory_abi!();

use kosmos_plugin_sdk::{http_fetch, read_string, write_string};

/// Pure guest-side round trip, no host import involved — proves the basic call+memory mechanism.
#[no_mangle]
pub extern "C" fn greet(name_ptr: i32, name_len: i32) -> i64 {
    let name = unsafe { read_string(name_ptr, name_len) };
    write_string(&format!("Hello, {name}!"))
}

/// Calls the host's permission-scoped `http_fetch` import and returns whatever it gave back —
/// either a real response body (host allowed the URL's host) or the host's own "blocked:<host>"
/// message (it didn't). The guest can't tell which case it got except by reading the text, same
/// as any real plugin would see.
#[no_mangle]
pub extern "C" fn fetch_allowed(url_ptr: i32, url_len: i32) -> i64 {
    let url = unsafe { read_string(url_ptr, url_len) };
    write_string(&http_fetch(&url))
}
