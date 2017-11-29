package main

import (
	"github.com/mndrix/tap-go"
)

func main() {
	t := tap.New()

	t.Ok(true, "ok")
	t.Diagnosticf("a demon to show it is ok")

	t.AutoPlan()
}
