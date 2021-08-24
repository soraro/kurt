# kurt
```
kurt: KUbernetes Reboot Tracker
  -l value
        Specify multiple times for the label keys you want to see.
        For example: -l app
  -n string
        Specify namespace for kurt to collect reboot metrics.
        Leave blank to collect in all namespaces.
```

# Requirements
Go Version 1.16

# Building
```
go build ./cmd/kurt
```
Outputs a `kurt` binary

# Testing
```
go test ./cmd/kurt -v
```
