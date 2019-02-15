package calendar

import (
	"fmt"
	"testing"
)

func checkTimeValid(t *testing.T, h, m, s int) {
	ti := NewTime(h, m, s)
	sti := ti.String()
	xsti := fmt.Sprintf("%02d:%02d:%02d", h,m,s)
	if sti != xsti {
		t.Errorf("Time.String() failed, expected %s, got %s", xsti, sti)
	}
	json, err := ti.MarshalJSON()
	if err != nil {
		t.Errorf("Time.MarshalJSON() failed to serialize %s", sti)
	}
	sjson := string(json)
	if len(sjson) != 10 {
		t.Errorf("Time.MarshalJSON() serialized %s to %s", sti, sjson)
	}
	correctJson := `"` + xsti + `"`
	if sjson != correctJson {
		t.Errorf("Time.MarshalJSON() serialized %s to %s, expected %s", sti, sjson, correctJson)
	}
	uti := Time{}
	err = uti.UnmarshalJSON(json)
	if err != nil {
		t.Errorf("Time.UnmarshalJSON() failed to deserialize %s", sjson)
	}
	if !uti.Equal(ti) {
		t.Errorf("Times unequal after JSON round trip: %s <=> %s", ti.String(), uti.String())
	}
}

func TestTime(t *testing.T) {
  for h := 0; h < 24; h++ {
  	for m := 0; m < 60; m++ {
  		for s := 0; s < 60; s++ {
  			checkTimeValid(t, h,m,s)
			}
		}
	}
}