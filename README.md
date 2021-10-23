# kurt
```
kurt: KUbernetes Restart Tracker

A restart tracker that gives context to what is restarting in your cluster

Usage:
  kurt [flags]
  kurt [command]

Available Commands:
  all         Print all groupings collected by kurt!
  completion  generate the autocompletion script for the specified shell
  help        Help about any command
  labels      Only print restart counts grouped by labels
  namespaces  Only print namespace-wide restart counts
  pods        Only print pod restart counts

Flags:
  -h, --help                help for kurt
  -l, --label strings       Specify multiple times for the label keys you want to see.
                            For example: -l app (default [*])
  -n, --namespace strings   Specify namespace for kurt to collect restart metrics.
                            Leave blank to collect in all namespaces.

Use "kurt [command] --help" for more information about a command.
```

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
