# protoform

A really simple provisioning tool, intended to set basic properties of a system via templated files to faciliate a more powerful configuration management tool taking over. This is intended to replace, or at least supplement, shell scripts that might be baked into your images and handling this kind of task now.

This uses Go's standard templating package, for, well, templating, and accepts JSON-formatted strings for more complex data like arrays and maps.

## Examples

Given `examples/hello.tmpl`:
```
Hello there, {{.name}}! Good {{.time}}!
```

We can use `protoform` to write a friendly message to admins that like to sleep in:
```
$ protoform name=sleepyhead time=morning examples/hello.tmpl > /etc/motd
$ cat /etc/motd
Hello there, sleepyhead! Good morning!
```

Something more complex: configuring resolv.conf. Using `examples/resolv.conf.tmpl`:
```
{{ range .nameservers -}}
nameserver {{ . }}
{{ end -}}

domain {{ .domain }}

{{ range $key, $value := .options -}}
option {{ $key }}{{ with $value }}{{ printf ":%s" . }}{{end}}
{{ end -}}
```

By using JSON-formatted strings, we can pass arrays and maps to `protoform`:
```
$ protoform nameservers='["10.20.30.40", "8.8.8.8"]' domain=mydomain.com options='{"rotate": "", "timeout": "5"}' examples/resolv.conf.tmpl > /etc/resolv.conf
$ cat /etc/resolv.conf
nameserver 10.20.30.40
nameserver 8.8.8.8
domain mydomain.com

option rotate
option timeout:5
```

## Things to do
- [ ] Try replacing the custom parsing logic with the standard flags package. Initial design required a custom flag parser, but this is probably no longer true.
- [x] Support more complex data than plain strings (arrays and maps, etc.)
- [ ] Custom template functions for handy data manipulations (i.e. starting an array-based list at 1)
- [ ] In-place file writing
- [x] Better error messages
