package buildx

import (
	"github.com/raphgm/container-doctor/pkg/check"
)

type Provider struct {
	checks []check.Check
}

func New(checks ...check.Check) *Provider {
	return &Provider{
		checks: checks,
	}
}

func (p *Provider) Name() string {
	return "Buildx"
}

func (p *Provider) Checks() []check.Check {
	return p.checks
}
