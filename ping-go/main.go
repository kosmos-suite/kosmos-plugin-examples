// Minimal reference plugin for Kosmos's WASM plugin host, in Go (TinyGo) — mirrors
// plugins/examples/ping (Rust) and plugins/examples/ping-as (AssemblyScript) to prove the host is
// genuinely language-agnostic. Not a real metadata source; exists purely to exercise the ABI end
// to end.
//
// ABI details (pointer/length packed into i64, alloc/dealloc, http_fetch) live in
// kosmos-plugin-sdk — see that module for the full explanation.
package main

import sdk "github.com/kosmos-suite/kosmos-plugin-sdk-go"

//go:wasmexport alloc
func alloc(size int32) int32 {
	return sdk.Alloc(size)
}

//go:wasmexport dealloc
func dealloc(ptr int32, length int32) {
	sdk.Dealloc(ptr, length)
}

// Pure guest-side round trip, no host import involved — proves the basic call+memory mechanism.
//
//go:wasmexport greet
func greet(namePtr int32, nameLen int32) int64 {
	name := sdk.ReadString(namePtr, nameLen)
	return sdk.WriteString("Hello, " + name + "!")
}

// Calls the host's permission-scoped http_fetch import and returns whatever it gave back — either
// a real response body (host allowed the URL's host) or the host's own "blocked:<host>" message
// (it didn't).
//
//go:wasmexport fetch_allowed
func fetch_allowed(urlPtr int32, urlLen int32) int64 {
	url := sdk.ReadString(urlPtr, urlLen)
	return sdk.WriteString(sdk.HTTPFetch(url))
}

func main() {}
