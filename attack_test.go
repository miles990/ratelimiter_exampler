package main

import (
	"testing"
	"time"

	vegeta "github.com/tsenart/vegeta/v12/lib"
)

func TestAttack(t *testing.T) {
	rate := vegeta.Rate{Freq: 61, Per: 60 * time.Second}
	duration := 60 * time.Second
	targeter := vegeta.NewStaticTargeter(vegeta.Target{
		Method: "GET",
		URL:    "http://127.0.0.1/",
	})
	attacker := vegeta.NewAttacker()

	var metrics vegeta.Metrics
	for res := range attacker.Attack(targeter, rate, duration, "Big Bang!") {
		metrics.Add(res)
	}
	metrics.Close()

	t.Logf("Requests:%v\n Rate:%v StatusCodes[200]:%v, StatusCodes[403]:%v", metrics.Requests, metrics.Rate, metrics.StatusCodes["200"], metrics.StatusCodes["403"])
}
