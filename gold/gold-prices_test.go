package main

import (
	"testing"
)

func TestGold_GetPrices(t *testing.T) {

	g := Gold{
		Prices: nil,
		Client: client,
	}

	p, err := g.GetPrices()

	if err != nil {
		t.Error(err)
	}

	if p.Price != 1969.7850 {
		t.Error("Wrong price returned:", p.Price)
	}
}
