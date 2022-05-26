package pack

import (
	"github.com/markbates/pkger"

	"github.com/probr/probr-pack-ubuntu/internal/access"
	"github.com/probr/probr-pack-ubuntu/internal/filesystem"
	"github.com/probr/probr-pack-ubuntu/internal/firewall"
	"github.com/probr/probr-pack-ubuntu/internal/group"
	"github.com/probr/probr-pack-ubuntu/internal/logging"
	"github.com/probr/probr-pack-ubuntu/internal/network"
	"github.com/probr/probr-pack-ubuntu/internal/services"
	"github.com/probr/probr-pack-ubuntu/internal/sshaccess"
	"github.com/probr/probr-pack-ubuntu/internal/useraccount"
	"github.com/probr/probr-sdk/probeengine"
)

// GetProbes returns a list of probe objects
func GetProbes() []probeengine.Probe {
	return []probeengine.Probe{
		access.Probe,
		filesystem.Probe,
		firewall.Probe,
		group.Probe,
		logging.Probe,
		network.Probe,
		services.Probe,
		sshaccess.Probe,
		useraccount.Probe,
	}
}

func init() {
	// pkger.Include is a no-op that directs the pkger tool to include the desired file or folder.
	pkger.Include("/internal/access/access.feature")
	pkger.Include("/internal/filesystem/filesystem.feature")
	pkger.Include("/internal/firewall/firewall.feature")
	pkger.Include("/internal/group/group.feature")
	pkger.Include("/internal/logging/logging.feature")
	pkger.Include("/internal/network/network.feature")
	pkger.Include("/internal/services/services.feature")
	pkger.Include("/internal/sshaccess/sshaccess.feature")
	pkger.Include("/internal/useraccount/useraccount.feature")
}
