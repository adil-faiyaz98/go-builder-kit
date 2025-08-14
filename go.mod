module github.com/adil-faiyaz98/go-builder-kit/v2

go 1.23.0

toolchain go1.24.0

// This repository is tagged as v1.10.0 for GitHub display purposes,
// but the actual module is v2. Users should install the v2 module:
// go get github.com/adil-faiyaz98/go-builder-kit/v2@latest

// Retract older versions - users should use v2.1.3+
retract (
	v2.1.1 // Pre-release version with incomplete features
	v2.0.9 // Pre-release version with incomplete features
	v2.0.5 // Pre-release version with incomplete features
	v2.0.2 // Pre-release version with incomplete features
	v2.0.1 // Pre-release version with incomplete features
	v2.0.0 // Pre-release version with incomplete features
)

require (
	github.com/onsi/ginkgo/v2 v2.23.4
	github.com/onsi/gomega v1.38.0
)

require (
	github.com/go-logr/logr v1.4.3 // indirect
	github.com/go-task/slim-sprig/v3 v3.0.0 // indirect
	github.com/google/go-cmp v0.7.0 // indirect
	github.com/google/pprof v0.0.0-20250630185457-6e76a2b096b5 // indirect
	go.uber.org/automaxprocs v1.6.0 // indirect
	golang.org/x/net v0.43.0 // indirect
	golang.org/x/sys v0.35.0 // indirect
	golang.org/x/text v0.28.0 // indirect
	golang.org/x/tools v0.36.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
