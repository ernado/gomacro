package messages

// Server <-> Client (Sync)
//procm:use=derive_binary
type TogglePVP struct {
	PlayerID   byte
	PVPEnabled bool
}
