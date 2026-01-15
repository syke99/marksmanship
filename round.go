package marksmanship

type round struct {
	impact func()
	// TODO: add the latency and
	// TODO: error counter metrics
	// TODO: here
}

type Round interface {
	// Impact ends the span that was started whenever
	// Fire was called
	Impact()
	MarkMiss(err error)
}

// Impact ends the span that was started whenever
// Fire was called
func (r *round) Impact() {
	r.impact()
}

func (r *round) MarkMiss(err error) {}
