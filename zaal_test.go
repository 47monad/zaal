package zaal_test

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/47monad/zaal"
)

func TestZaal(t *testing.T) {
	res, err := zaal.Build(
		"./testdata/main.cue",
		"./testdata/main.env",
	)
	if err != nil {
		t.Fatal(err)
	}

	js, err := json.MarshalIndent(res, "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(js))
}
