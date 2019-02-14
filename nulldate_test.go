package calendar

import (
	"fmt"
	"testing"
)

func checkNullDateValid(t *testing.T, y, m, d int) {
	da := NewNullDate(y, m, d)
	sda := da.String()
	xsda := fmt.Sprintf("%04d-%02d-%02d", y, m, d)
	if sda != xsda {
		t.Errorf("NullDate.String() failed, expected %s, got %s", xsda, sda)
	}
	json, err := da.MarshalJSON()
	if err != nil {
		t.Errorf("NullDate.MarshalJSON() failed to serialize %s", sda)
	}
	sjson := string(json)
	if len(sjson) != 12 {
		t.Errorf("NullDate.MarshalJSON() serialized %s to %s", sda, sjson)
	}
	uda := NullDate{}
	err = uda.UnmarshalJSON(json)
	if err != nil {
		t.Errorf("NullDate.UnmarshalJSON() failed to deserialize %s", sjson)
	}
	if !uda.Date.Equal(da.Date) {
		t.Errorf("NullDates unequal after JSON round trip: %s <=> %s", da.String(), uda.String())
	}
	sqlval, err := da.Value()
	if err != nil {
		t.Errorf("NullDate Value() failed: %v", err)
	}
	uda2 := NullDate{}
	err = uda2.Scan(sqlval)
	if err != nil {
		t.Errorf("NullDate Scan() failed: %v", err)
	}
	if !uda2.Equal(da) {
		t.Errorf("NullDates unequal after SQL round trip: %s <=> %s", da.String(), uda2.String())
	}
}

func checkNullDateNull(t *testing.T, y, m, d int) {
	da := NewNullDate(y, m, d)
	da.Valid = false
	sda := da.String()
	if sda != "null" {
		t.Errorf("NullDate.String() failed, expected null, got %s", sda)
	}
	json, err := da.MarshalJSON()
	if err != nil {
		t.Errorf("NullDate.MarshalJSON() failed to serialize %s", sda)
	}
	sjson := string(json)
	if sjson != "null" {
		t.Errorf("NullDate.MarshalJSON() serialized %s to %s", sda, sjson)
	}
	uda := NullDate{}
	err = uda.UnmarshalJSON(json)
	if err != nil {
		t.Errorf("NullDate.UnmarshalJSON() failed to deserialize %s", sjson)
	}
	if !uda.Equal(da) {
		t.Errorf("NullDates unequal after JSON round trip: %s <=> %s", da.String(), uda.String())
	}
	sqlval, err := da.Value()
	if err != nil {
		t.Errorf("NullDate Value() failed: %v", err)
	}
	uda2 := NullDate{}
	err = uda2.Scan(sqlval)
	if err != nil {
		t.Errorf("NullDate Scan() failed: %v", err)
	}
	if !uda2.Equal(da) {
		t.Errorf("NullDates unequal after SQL round trip: %s <=> %s", da.String(), uda2.String())
	}
}

func TestNullDate(t *testing.T) {
	for y := 1999; y <= 2001; y++ {
		for m := 1; m <= 12; m++ {
			checkNullDateValid(t, y, m, 1)
			checkNullDateNull(t, y, m, 1)
		}
	}
}
