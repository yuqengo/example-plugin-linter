This is an example linter that can be compiled into a plugin for `golangci-lint`.

To use this, download the code (I haven't found a way to enable `go get` functionality for plugins yet, if you know it let me know!).

From the root project directory, run `go build -buildmode=plugin -o plugin/example.go`. Then copy the generated `example.so` file into your project or to some other known location of your choosing. Alternately, you can use the `-o /path/to/location/example.so` output flag to have it put it there for you.

Next, download the `i473` branch of https://github.com/dbraley/golangci-lint/tree/i473 (will be easier once the PR created and merged, sorry).

From that projects root directy, run `make vendor_free_build`, which will create a local version of golangci-lint. Feel free to copy it to locations as well, however you want to access it.

In the project you want to lint, copy the golangci.yml file if you do not have one. Adjust the yaml to appropriate `linters-settings:custom` entries, and if `disable-all` is true add the linter name to `linters:enabled` (or manually enable them):

```
linters-settings:
 custom:
  example:
   path: /example.so #probably need to adjust this
```
