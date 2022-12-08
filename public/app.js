const go = new Go()
WebAssembly.instantiateStreaming(fetch("app.wasm"), go.importObject)
  .then((result) => {
    go.run(result.instance)
  })

window.sqlite3InitModule()
  .then((sqlite3) => {
  self.sqlite3 = sqlite3
    const capi = sqlite3.capi
    const oo = sqlite3.oo1
    console.log("sqlite3 version", capi.sqlite3_libversion(), capi.sqlite3_sourceid())
})
