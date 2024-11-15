package main

import (
	"github.com/spikepanx/statmetrics"
)

func main() {
	statmetrics.ParseStats([]byte(validJSON))
}
