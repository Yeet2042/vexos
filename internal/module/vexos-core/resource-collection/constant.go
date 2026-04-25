package resourcecollection

type DRBDRole string
type DRBDConnectionState string

const (
	DRBDRolePrimary   DRBDRole = "primary"
	DRBDRoleSecondary DRBDRole = "secondary"
)

const (
	DRBDConnectionStateConnected  DRBDConnectionState = "connected"
	DRBDConnectionStateSyncTarget DRBDConnectionState = "sync_target"
	DRBDConnectionStateSyncSource DRBDConnectionState = "sync_source"
	DRBDConnectionStateStandAlone DRBDConnectionState = "stand_alone"
)
