[![CircleCI](https://circleci.com/gh/giantswarm/net-exporter.svg?&style=shield)](https://circleci.com/gh/giantswarm/net-exporter) [![Go Report Card](https://goreportcard.com/badge/github.com/giantswarm/net-exporter)](https://goreportcard.com/report/github.com/giantswarm/net-exporter)

# net-exporter

net-exporter is a Prometheus exporter for exposing network information in Kubernetes clusters.
It is packaged as a Helm chart.

net-exporter runs as a Kubernetes Daemonset. This is to allow for intra-pod network calls,
to determine network latency.

## How to build

Build it using the standard `go build` command.

```bash
go build github.com/giantswarm/net-exporter
```

## Deployment

* Managed by [app-operator].
* Production releases are stored in the [default-catalog].
* WIP releases are stored in the [default-test-catalog].

## Installing the Chart

To install the chart locally:

```bash
$ git clone https://github.com/giantswarm/net-exporter.git
$ cd net-exporter
$ helm install helm/net-exporter
```

Provide a custom `values.yaml`:

```bash
$ helm install net-exporter -f values.yaml
```

## Release Process

* Ensure CHANGELOG.md is up to date.
* Create a new GitHub release with the version e.g. `v0.1.0` and link the
changelog entry.
* This will push a new git tag and trigger a new tarball to be pushed to the
[default-catalog].  
* Update [cluster-operator] with the new version.

[app-operator]: https://github.com/giantswarm/app-operator
[cluster-operator]: https://github.com/giantswarm/cluster-operator
[default-catalog]: https://github.com/giantswarm/default-catalog
[default-test-catalog]: https://github.com/giantswarm/default-test-catalog

## Collectors
All Collectors are enabled by default.

Name | Description
-----|-------------
dns | Exposes DNS latency statistics. Performs host lookups, exposing the time taken per host.
network | Exposes network latency statistics. Performs dials to the other net-exporter Pods, exposing the time taken per host.

## Metrics

Name | Description
-----|------------
`dns_latency_seconds_bucket` | A Prometheus Histogram of DNS resolution latency. See also `dns_latency_seconds_count` and `dns_latency_seconds_sum`.
`dns_resolve_error_total` | The total number of errors encountered resolving DNS.
`dns_error_total` | The total number of internal errors encountered testing DNS resolution.
`network_latency_seconds_bucket` | A Prometheus Histogram of network latency. See also `network_latency_seconds_count` and `network_latency_seconds_sum`.
`network_dial_error_total` | The total number of errors encountered dialing other hosts.
`network_error_total` | The total number of internal errors encountered testing network latency.

For example (some labels ommited for clarity):
```
dns_latency_seconds_bucket{instance="192.168.120.239:8000", host="kubernetes.default.svc.cluster.local", le="0.008"} | 7
```
Here, we expose the latency for the specific instance to resolve the dns host.

```
network_latency_seconds_bucket{instance="192.168.120.239:8000", host="192.168.120.239:8000", le="0.004"} | 28
```
Here, we expose the latency for the specific instance to resolve another instance (specifically, the net-exporter pod, labeled as host).

## Contact

- Mailing list: [giantswarm](https://groups.google.com/forum/!forum/giantswarm)
- IRC: #[giantswarm](irc://irc.freenode.org:6667/#giantswarm) on freenode.org
- Bugs: [issues](https://github.com/giantswarm/net-exporter/issues)

## Contributing & Reporting Bugs

See [CONTRIBUTING](CONTRIBUTING.md) for details on submitting patches, the
contribution workflow as well as reporting bugs.

## License

net-exporter is under the Apache 2.0 license. See the [LICENSE](LICENSE) file for
details.
