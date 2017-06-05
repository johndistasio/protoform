# Cauldron

[![Build Status](https://travis-ci.org/johndistasio/cauldron.svg?branch=master)](https://travis-ci.org/johndistasio/cauldron)

A really simple provisioning tool, intended to set basic properties of a system via templated files to faciliate a more powerful configuration management tool taking over. This is intended to replace, or at least supplement, shell scripts that might be baked into your images for handling this kind of task now. Cauldron uses Go's standard templating package and accepts JSON-formatted strings or files for more complex data like arrays and maps. The built-in template functions are supplemented with the [Sprig](https://masterminds.github.io/sprig/) library for maximum text-wrangling power.

Cauldron is inspired by [consul-template](https://github.com/hashicorp/consul-template).

## Usage

```
$ cauldron [arguments] [template parameters]
```

By default, Cauldron will print the rendered template to standard out.

### Arguments

Cauldron recognizes the following arguments:

`-exec <cmd>`

Run the specified command after successful template rendering. The command does not run in a shell so redirection, pipes, etc. will not work.

`-file <path>`

Write the rendered template to the specified path instead of standard output. Useful when executing Cauldron outside of a shell.

`-inplace`

Write the rendered template in-place instead of standard output, overwriting the template file.

`-json <path>`

Read template data from the specified path. Template parameters provided on the command line are ignored.

`-help`

Prints a help message then exits.

`-template <path>`

Path to the template to be rendered. This argument is required.

`-version`

 Prints version and build information, then exits.

### Template Parameters

Command-line template parameters are key-value pairs in the form of `key=value` that are munged into an object fed to the template. A parameter like `kittens=fuzzy` would be accessible in the template with `{{ .kittens }}`.

More complex data can be provided with JSON-formatted strings, e.g. `animals='["cow", "sheep", "duck"]'`.

## Examples

The files shown here are all in the `examples/` directory.

Given `examples/hello.tmpl`:

```
Hello there, {{.name}}! Good {{.time}}!
```

We can use Cauldron to write a friendly message to admins that like to sleep in:

```
$ cauldron -template examples/hello.tmpl name=sleepyhead time=morning > /etc/motd
$ cat /etc/motd
Hello there, sleepyhead! Good morning!
```

### JSON

We can use JSON-formatted data to configure something more complex like `resolv.conf`. Using `examples/resolv.conf.tmpl`:

```
{{ range .nameservers -}}
nameserver {{ . }}
{{ end -}}

domain {{ .domain }}

{{ range $key, $value := .options -}}
option {{ $key }}{{ with $value }}{{ printf ":%s" . }}{{end}}
{{ end -}}
```

By using JSON-formatted strings, we can pass arrays and maps to Cauldron:

```
$ cauldron -template examples/resolv.conf.tmpl nameservers='["10.20.30.40", "8.8.8.8"]' domain=mydomain.com options='{"rotate": "", "timeout": "5"}' > /etc/resolv.conf
$ cat /etc/resolv.conf
nameserver 10.20.30.40
nameserver 8.8.8.8
domain mydomain.com

option rotate
option timeout:5
```

Cauldron can also read data from a JSON file with the `-json` flag. Using `examples/treats.json`:

```
{
  "icecream": [ "chocolate", "vanilla", "strawberry" ],
  "slushes": [ "grape", "watermelon", "strawberry" ]
}
```

We can write a simple template like `examples/treats.tmpl` to list all of these things:

```
Summer Treats Menu:

Ice Cream:
{{ range $index, $flavor := .icecream -}}
    {{ add1 $index }}: {{ $flavor }}
{{ end }}
Slushes:
{{ range $index, $flavor := .slushes -}}
    {{ add1 $index }}: {{ $flavor }}
{{ end -}}
```

Finally, we render the template:

```
$ cauldron -template examples/treats.tmpl -json examples/treats.json
Summer Treats Menu:

Ice Cream:
1: chocolate
2: vanilla
3: strawberry

Slushes:
1: grape
2: watermelon
3: strawberry
```

## Building and Packaging

Cauldron can be built or prepared for packaging with `make`. A spec file is included for RPM builds. The default `make` target will run the tests and build a binary for the current platform in `build/`. Available targets are:

`archive`

Packs up the source tree into a tarball in `build/`.

`build`

Builds the current working tree into a binary in `build/`.

`clean`

Deletes all build artifacts.

`test`

Runs all Go tests.
