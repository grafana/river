# River

[![Go Reference](https://pkg.go.dev/badge/github.com/grafana/river.svg)](https://pkg.go.dev/github.com/grafana/river)

River is an HCL-inspired configuration language originally written for
[Grafana Agent flow mode][flow] with the following goals:

* _Fast_: River is intended to be used in applications that may evaluate River
  expression multiple times a second.
* _Simple_: River must be easy to read and write to minimize the learning
  curve of yet another configuration language.
* _Debuggable_: River must give detailed information when there's a mistake in
  configuration.

```river
// Discover Kubernetes pods to collect metrics from.
discovery.kubernetes "pods" {
  role = "pod"
}

// Collect metrics from Kubernetes pods.
prometheus.scrape "default" {
  targets    = discovery.kubernetes.pods.targets
  forward_to = [prometheus.remote_write.default.receiver]
}

// Get an API key from disk.
local.file "apikey" {
  filename  = "/var/data/my-api-key.txt"
  is_secret = true
}

// Send metrics to a Prometheus remote_write endpoint.
prometheus.remote_write "default" {
  endpoint {
    url = "http://localhost:9009/api/prom/push"

    basic_auth {
      username = "MY_USERNAME"
      password = local.file.apikey.content
    }
  }
}
```


For more information on how to use River, see [our Go documentation][docs].

[flow]: https://grafana.com/docs/agent/latest/flow
[docs]: https://pkg.go.dev/github.com/grafana/river
