# kurt
```
kurt: KUbernetes Restart Tracker

A restart tracker that gives context to what is restarting in your cluster

Usage:
  kurt [command]

Available Commands:
  all         Print all groupings collected by kurt!
  completion  generate the autocompletion script for the specified shell
  help        Help about any command
  labels      Only print restart counts grouped by labels
  namespaces  Only print namespace-wide restart counts
  nodes       Only print node restart counts
  pods        Only print pod restart counts
  version     Print the current version and exit

Flags:
  -h, --help                help for kurt
  -l, --label strings       Specify multiple times for the label keys you want to see.
                            For example: "kurt all -l app"
  -c, --limit int           Limit the number of resources you want to see. Set limit to 0 for no limits. Must be positive.
                            For example: "kurt all -c=10" (default 5)
  -n, --namespace strings   Specify namespace for kurt to collect restart metrics.
                            Leave blank to collect in all namespaces.

Use "kurt [command] --help" for more information about a command.
```

# Install
Head over to our [releases page](https://github.com/soraro/kurt/releases/latest) or run as a `kubectl` plugin with [krew](https://krew.sigs.k8s.io/)
```
kubectl krew install kurt
```

[Easy install for krew](https://krew.sh/)

# Examples
Show the top 5 highest restart counts grouped by `Namespace`, `Node`, `Label`, and `Pod`:
```
$ kurt all

kurt: KUbernetes Restart Tracker

==========

 Namespace      Restarts

 default        2
 test           1
 kube-system    0

==========

 Node           Restarts

 minikube-m02   2
 minikube-m03   1
 minikube       0

==========

 Label                                          Restarts

 run:nginx                                      3
 component:etcd                                 0
 k8s-app:kube-proxy                             0
 addonmanager.kubernetes.io/mode:Reconcile      0
 integration-test:storage-provisioner           0

==========

 Pod                            Namespace       Restarts

 nginx                          default         2
 nginx                          test            1
 kube-apiserver-minikube        kube-system     0
 storage-provisioner            kube-system     0
 etcd-minikube                  kube-system     0
```

Show more results:
```
kurt all -c 10

# use -c 0 if you want to show all results
```

Show which node has the most restarted pods:
```
kurt no
```

Show top 20 pod restart counts in the `default` namespace which also have the `app` label key:
```
kurt po -n default -l app -c 20
```

Get help:
```
kurt -h
```

# Permissions
As seen in the [`cmd/collect.go` file](https://github.com/soraro/kurt/blob/main/cmd/collect.go) the only permission required for kurt is `pods/list`.

# Requirements
Go Version 1.16

# Building
```
go build .
```
Outputs a `kurt` binary

# Testing
```
go test ./cmd -v
```
