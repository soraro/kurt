package cmd

var namespaceTracker = make(map[string]int32)
var nodeTracker = make(map[string]int32)
var podTracker = make(map[string]int32)
var labelTracker = make(map[string]int32)
var containerTracker = make(map[string]map[string]int32)

var printAll bool
var printNS bool
var printNode bool
var printPods bool
var printLabel bool
