package fsm

import (
	"context"
	ctrl "sigs.k8s.io/controller-runtime"
)

// to save the KymaPolicyConfig status at the beginning of the reconciliation
func sFnTakeSnapshot(ctx context.Context, m *fsm, s *systemState) (stateFn, *ctrl.Result, error) {
	m.log.Info("Take snapshot of the current KymaPolicyConfig status")
	s.savePolicyStatus()
	return switchState(sFnInitialize)
}
