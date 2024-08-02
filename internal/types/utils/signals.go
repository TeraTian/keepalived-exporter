package utils

import (
	"syscall"

	"github.com/hashicorp/go-version"
	"github.com/sirupsen/logrus"
)

var (
	sigNumSupportedVersion = version.Must(version.NewVersion("1.3.8"))

	defaultSignals = map[string]syscall.Signal{"DATA": syscall.SIGUSR1, "STATS": syscall.SIGUSR2}
)

// HasSigNumSupport checks if Keepalived supports --signum command.

func HasSigNumSupport(version *version.Version) bool {
	return version == nil || version.GreaterThanOrEqual(sigNumSupportedVersion)
}

// GetDefaultSignal returns default signals for Keepalived.

func GetDefaultSignal(sigString string) syscall.Signal {
	sig, ok := defaultSignals[sigString]

	if !ok {
		logrus.WithField("signal", sigString).Fatal("Unsupported signal for your keepalived")
	}

	return sig
}
