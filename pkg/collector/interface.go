package collector

type MetricsCollector interface {
	Gather() error
}
