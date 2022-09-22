package cpu

type CPUStats struct {
	cpuUtilPercent float64
}

func (c *CPUStats) Gather() error {
	c.cpuUtilPercent = 0.1
	return nil
}
