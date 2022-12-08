const go = new Go()
WebAssembly.instantiateStreaming(fetch("app.wasm"), go.importObject)
  .then((result) => {
    go.run(result.instance)
  })
