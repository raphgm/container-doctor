package report

type Status string

const (
	Pass Status = "PASS"
	Warn Status = "WARN"
	Fail Status = "FAIL"
	Skip Status = "SKIP"
)

type Severity string

const (
	Info     Severity = "INFO"
	Warning  Severity = "WARNING"
	Critical Severity = "CRITICAL"
)

type Result struct {
	ID string

	Name string

	Status Status

	Severity Severity

	Summary string

	Details string

	Recommendation string

	Documentation string
}
