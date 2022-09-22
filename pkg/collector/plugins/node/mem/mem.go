package mem

type MemStats struct {
	UtilPercent float64
}

func (m *MemStats) Gather() error {
	m.UtilPercent = 0.2
	return nil
}
