# protoform

A really simple provisioning tool, intended to set basic properties of a system via templated files to faciliate a more powerful configuration management tool taking over. This is intended to replace, or at least supplement, shell scripts that might be baked into your images and handling this kind of task now.

This uses Go's built-in templating system.

Given `example.tmpl`:
```
Hello there, {{.name}}! Good {{.time}}!
```

We can use `protoform` to write a friendly message to admins that like to sleep in:
```
$ protoform name=sleepyhead time=morning example.tmpl > /etc/motd
$ cat /etc/motd
Hello there, sleepyhead! Good morning!
```

## Things to do
- [ ] Can the custom flag parsing logic be replaced by the standard flags package? Initial design required a custom flag parser; this is probably no longer true.
- [ ] Support more complex data than plain strings (arrays and maps, etc.)
- [ ] In-place file writing
- [ ] Better error messages
