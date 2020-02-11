package main

import (
	"os"

	"github.com/containers-ai/alameda/operator/pkg/probe"
)

func livenessProbe(cfg *probe.LivenessProbeConfig) {
	// probe.LivenessProbe(cfg)
	os.Exit(0)
}

func readinessProbe(cfg *probe.ReadinessProbeConfig) {
	// probe.ReadinessProbe(cfg)
	os.Exit(0)
}
