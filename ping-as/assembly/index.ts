// Minimal reference plugin for Kosmos's WASM plugin host, in AssemblyScript — mirrors
// plugins/examples/ping (Rust) to prove the host is genuinely language-agnostic. Not a real
// metadata source; exists purely to exercise the ABI end to end.
//
// ABI details (pointer/length packed into i64, alloc/dealloc, http_fetch) live in
// kosmos-plugin-sdk — see that package for the full explanation.

export { alloc, dealloc } from "@kosmos-suite/plugin-sdk-assemblyscript/assembly/index";
import { httpFetch, readString, writeString } from "@kosmos-suite/plugin-sdk-assemblyscript/assembly/index";

// Pure guest-side round trip, no host import involved — proves the basic call+memory mechanism.
export function greet(namePtr: i32, nameLen: i32): i64 {
  const name = readString(namePtr, nameLen);
  return writeString(`Hello, ${name}!`);
}

// Calls the host's permission-scoped http_fetch import and returns whatever it gave back —
// either a real response body (host allowed the URL's host) or the host's own "blocked:<host>"
// message (it didn't).
export function fetch_allowed(urlPtr: i32, urlLen: i32): i64 {
  const url = readString(urlPtr, urlLen);
  return writeString(httpFetch(url));
}
