package main

import (
	"github.com/mndrix/tap-go"
)

func main() {
	t := tap.New()

	t.Ok(false, "fail")
	t.Diagnosticf("a demon to show it is fail")

	t.AutoPlan()
}
