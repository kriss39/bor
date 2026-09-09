package whitelist

import "testing"

func TestCheckpointInvalidMetricsDoNotDecrementValidMeters(t *testing.T) {
	chainBefore := CheckpointChainMeter.Snapshot().Count()
	peerBefore := CheckpointPeerMeter.Snapshot().Count()

	reportCheckpointMetrics(false, true, false)
	reportCheckpointMetrics(false, false, true)

	if got := CheckpointChainMeter.Snapshot().Count() - chainBefore; got != 0 {
		t.Fatalf("invalid checkpoint chain changed valid meter by %d", got)
	}
	if got := CheckpointPeerMeter.Snapshot().Count() - peerBefore; got != 0 {
		t.Fatalf("invalid checkpoint peer changed valid meter by %d", got)
	}
}
