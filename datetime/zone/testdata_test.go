package zone_test

import (
	_ "embed"
	"github.com/Jumpaku/tokiope/datetime/zone"
)

//go:embed testdata/tzot_test.json
var tzotTestJson []byte

func getTestProvider() zone.Provider {
	p, err := zone.LoadProvider(tzotTestJson, "test")
	if err != nil {
		panic(err)
	}
	return p
}
