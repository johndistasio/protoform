# Protoform

A really simple provisioning tool, intended to set basic properties of a system via templated files to faciliate a more powerful configuration management tool taking over. This is intended to replace, or at least supplement, shell scripts that might be baked into your images and handling this kind of task now.

Protoform uses Go's standard templating package and accepts JSON-formatted strings for more complex data like arrays and maps. The built-in template functions are supplemented with the [Sprig](https://masterminds.github.io/sprig/) library for maximum text-wrangling power.

Protoform will print the rendered template for standard out by default. With the `-inplace` flag, Protoform will overwrite the specified template file with the rendered version - useful if you don't want to leave things lying around post-provisioning!

## Examples

Given `examples/hello.tmpl`:
```
Hello there, {{.name}}! Good {{.time}}!
```

We can use Protoform to write a friendly message to admins that like to sleep in:
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

By using JSON-formatted strings, we can pass arrays and maps to Protoform:
```
$ protoform nameservers='["10.20.30.40", "8.8.8.8"]' domain=mydomain.com options='{"rotate": "", "timeout": "5"}' examples/resolv.conf.tmpl > /etc/resolv.conf
$ cat /etc/resolv.conf
nameserver 10.20.30.40
nameserver 8.8.8.8
domain mydomain.com

option rotate
option timeout:5
```
