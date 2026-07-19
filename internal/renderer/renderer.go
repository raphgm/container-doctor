package renderer

import "github.com/raphgm/container-doctor/pkg/report"

type Renderer interface {
	Render(report.Report) error
}
