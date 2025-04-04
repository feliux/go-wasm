### Understanding WebAssembly

WebAssembly (Wasm) has established itself as a pivotal technology, enabling quick and efficient code execution in web browsers and forming a robust bridge between web applications and the high-performance typically associated with native applications.

Wasm acts as a low-level virtual machine, executing a compact binary code that’s translated from high-level languages.

- Primary Advantages:

1. Universal Browser Support: Thanks to its support from all major browsers, Wasm delivers consistent performance across diverse platforms.
2. Near-Native Performance: Intentionally designed to execute binary code at a speed akin to native applications, Wasm enhances the responsiveness of web applications considerably.

### Why WebAssembly?

1. Execute In-Browser: apps could operate straight within the browser, sidestepping server maintenance overheads and repetitive API calls, and notably, making ongoing maintenance a breeze in comparison to older server-based approaches.
2. Achieve Peak Performance: Employing WebAssembly ensures that our application operates with a level of performance that competes with native applications, enhancing user interactions and bolstering response times.

### How WebAssembly Works with the Browser

When a Wasm module is loaded in a browser, it is executed by a virtual machine called the WebAssembly Runtime, which translates the Wasm code into machine code that the browser’s JavaScript engine can execute.

The WebAssembly Runtime is implemented in the browser as a JavaScript library and provides a set of APIs for loading, validating and executing Wasm modules. When a Wasm module is loaded, the Runtime validates the module’s bytecode and creates an instance of the module, which can be used to call its functions and access its data.

Wasm modules can interact with the browser’s Document Object Model (DOM) and other web APIs using JavaScript. For example, a Wasm module can modify the contents of a webpage, listen for user events, and make network requests using the browser’s web APIs.

One of the key benefits of using Wasm with the browser is that it provides a way to run code that is more performant than JavaScript. JavaScript is an interpreted language, which means that it can be slower than compiled languages like C++ or Go. However, by compiling code into Wasm format, it can be executed at near-native speeds, making it ideal for computationally intensive tasks such as machine learning or 3D graphics rendering.

## References

https://thenewstack.io/webassembly-and-go-a-guide-to-getting-started-part-1/

https://www.fermyon.com/blog/tinygo-webassembly-favicon-server

https://www.youtube.com/watch?v=4kBvvk2Bzis

**containers**

https://medium.com/@shivraj.jadhav82/webassembly-wasm-docker-vs-wasm-275e317324a1

https://kodekloud.com/blog/webassembly-vs-docker/

https://wasmedge.org/wasm_linux_container/

**ml**

https://github.com/PacktPublishing/Hands-On-Deep-Learning-with-Go

https://gorgonia.org/

https://syslog.ravelin.com/go-tensorflow-74d1101fab3f
https://pkg.go.dev/github.com/tensorflow/tensorflow/tensorflow/go
https://github.com/tensorflow/build/tree/master/golang_install_guide
https://www.tensorflow.org/tutorials/keras/save_and_load?hl=es-419

https://github.com/galeone/tfgo
https://github.com/unravelin/gotf
 