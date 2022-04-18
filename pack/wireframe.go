package pack

import (
	"github.com/markbates/pkger"

	"github.com/probr/probr-pack-ubuntu/internal/ubuntu"
	"github.com/probr/probr-sdk/probeengine"
)

// GetProbes returns a list of probe objects
func GetProbes() []probeengine.Probe {
	return []probeengine.Probe{
		ubuntu.Probe,
	}
}

func init() {
	// pkger.Include is a no-op that directs the pkger tool to include the desired file or folder.
	pkger.Include("/internal/ubuntu/ubuntu.feature")
}
