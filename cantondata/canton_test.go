package cantondata_test

import (
	"testing"

	"github.com/bit-en5/swisstowngame/cantondata"
)

func TestGetName(t *testing.T) {
	has := cantondata.GetName("BE")
	want := "Bern"

	if has != want {
		t.Errorf("want %s / has %s", want, has)
	}
}
