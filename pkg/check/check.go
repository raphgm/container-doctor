package check

import (
	"context"

	"github.com/raphgm/container-doctor/pkg/report"
)

type Check interface {
	ID() string

	Name() string

	Description() string

	Run(context.Context) report.Result
}
