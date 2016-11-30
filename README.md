# lowger
Simple pluggable logger for Golang

### How to use?
Simply import `github.com/JesusIslam/lowger` and it will provide `Fatal`, `Error`, `Warning`, and `Info` methods.

### How to add my own plugin?
Just do `lowger.AddPlugin(YOUR_PLUGIN_HERE)` where your plugin must implement `Logger` interface.