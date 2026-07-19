package provider

import "github.com/raphgm/container-doctor/pkg/check"

type Provider interface {
	Name() string

	Checks() []check.Check
}
