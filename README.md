This is an example linter that can be compiled into a plugin for `golangci-lint`.

To use this:

### Create the Plugin From This Linter

1. Download the code \*
2. From the root project directory, run `go build -buildmode=plugin -o plugin/example.go`.
3. Copy the generated `example.so` file into your project or to some other known location of your choosing. \**


### Create a Copy of `golangci-lint` that Can Run with Plugins
(After the PR this will be either the production code or simpler to get hopefully)

1. Download the `i473` branch of https://github.com/dbraley/golangci-lint/tree/i473
2. From that projects root directy, run `make vendor_free_build`
3. Copy the `golangci-lint` executable that was created to your path, project, or other location

### Configure Your Project for Linting

1. If the project you want to lint does not have one already, copy the https://github.com/dbraley/golangci-lint/blob/i473/.golangci.yml to the root directory.
2. Adjust the yaml to appropriate `linters-settings:custom` entries as so:
```
linters-settings:
 custom:
  example: # If this doesn't match the linters name definition, it will warn you but still run
   path: /example.so # Adjust this the location of the plugin
   enabled: true # This determines if the linter is run by default
   original-url: github.com/dbraley/example-linter # This is just documentation for custom linters
   slow: false # Set this to true to observe `--fast` option
```
3. If your `.golangci.yml` file has `linters:disable-all` set to true, you'll also need to add your linter to the `linters:enable` seciont:
```
linters:
 enable:
  # ...
  - varcheck
  - whitespace
  # Custom linters
  - example
```
4. Alternately, you can leave it disabled and turn it on via command line only: `golangci-lint run -Eexample`


\* Sorry, I haven't found a way to enable `go get` functionality for plugins yet. If you know how, let me know!
\** Alternately, you can use the `-o /path/to/location/example.so` output flag to have it put it there for you.
