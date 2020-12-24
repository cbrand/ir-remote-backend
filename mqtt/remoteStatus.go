package mqtt

import (
	"time"

	"github.com/cbrand/ir-remote-backend/protocol"
)

// RemoteStatus returns the current ifnormation of a device status.
type RemoteStatus struct {
	protocol.RemoteStatus
	// LastTransferredLifeSignDate represents the time which was last time pushed
	LastTransferredLifeSignDate time.Time
	// TicksMs represents the last transferred tick configurations
	TicksMs int
}

// Update updates the current status and checks whether or not the underlying remote is online
func (remoteStatus *RemoteStatus) Update() *RemoteStatus {
	remoteStatus.Online = time.Now().Before(remoteStatus.LastTransferredLifeSignDate.Add(1*time.Minute + 30*time.Second))
	remoteStatus.Lifesign = Time{remoteStatus.LastTransferredLifeSignDate}.ToIso8601()
	return remoteStatus
}
