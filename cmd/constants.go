package cmd

var namespaceTracker map[string]int32
var podTracker map[string]int32
var labelTracker map[string]int32

var printAll bool
var printNS bool
var printPods bool
var printLabel bool
