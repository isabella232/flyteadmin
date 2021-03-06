package interfaces

import (
	"context"

	"github.com/lyft/flyteadmin/pkg/executioncluster"
)

// Interface for the Execution Cluster
type ClusterInterface interface {
	GetTarget(context.Context, *executioncluster.ExecutionTargetSpec) (*executioncluster.ExecutionTarget, error)
	GetAllValidTargets() []executioncluster.ExecutionTarget
}
