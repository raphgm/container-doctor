package report

import "time"

type ProviderReport struct {
	Name string

	Results []Result
}

type Report struct {
	Timestamp time.Time

	Duration time.Duration

	Providers []ProviderReport
}
