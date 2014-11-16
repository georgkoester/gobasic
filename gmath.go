package gobasic

// half-up (neg: half-down) round seems to be missing from Go class library.
func Round(val float64) int64 {
	if val > 0 {
		return int64(val + 0.5)
	} else {
		return int64(val - 0.5)
	}
}
