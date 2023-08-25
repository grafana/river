# River

River is an HCL-inspired configuration language originally written for
[Grafana Agent flow mode][flow] with the following goals:

* _Fast_: River is intended to be used in applications that may evaluate River
  expression multiple times a secret.
* _Simple_: River must be easy to read and write to minimize the learning
  curve of yet another configuration language.
* _Debuggable_: River must give detailed information when there's a mistake in
  configuration.

For more information on how to use River, see [our Go documentation][docs].

[flow]: https://grafana.com/docs/agent/latest/flow
[docs]: https://pkg.go.dev/github.com/grafana/river
