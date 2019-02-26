package calendar_test

import (
	"encoding/json"
	"fmt"
	"github.com/lpar/calendar"
	"testing"
	"time"
)

func checkValid(t *testing.T, y, m, d int) {
	da := calendar.NewDate(y, m, d)
	sda := da.String()
	xsda := fmt.Sprintf("%04d-%02d-%02d", y, m, d)
	if sda != xsda {
		t.Errorf("Date.String() failed, expected %s, got %s", xsda, sda)
	}
	json, err := da.MarshalJSON()
	if err != nil {
		t.Errorf("Date.MarshalJSON() failed to serialize %s", sda)
	}
	sjson := string(json)
	if len(sjson) != 12 {
		t.Errorf("Date.MarshalJSON() serialized %s to %s", sda, sjson)
	}
	uda := calendar.Date{}
	err = uda.UnmarshalJSON(json)
	if err != nil {
		t.Errorf("Date.UnmarshalJSON() failed to deserialize %s", sjson)
	}
	if time.Time(uda) != time.Time(da) {
		t.Errorf("Dates unequal after JSON round trip: %s <=> %s", da.String(), uda.String())
	}

}

func TestDate(t *testing.T) {
	for y := 1999; y <= 2001; y++ {
		for m := 1; m <= 12; m++ {
			maxd := []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}[m-1]
			checkValid(t, y, m, maxd)
			checkValid(t, y, m, 1)
		}
	}
}

func TestStruct(t *testing.T) {
	// Check that values still marshal/unmarshal correctly when inside structs
	tnow := time.Now()

	asst := struct {
		Date calendar.Date
		NullDate calendar.NullDate
		Time calendar.Time
		NullTime calendar.NullTime
	}{
		Date: calendar.DateFromTime(tnow),
		NullDate: calendar.NullDateFromTime(tnow),
		Time: calendar.Time(tnow),
		NullTime: calendar.NullTimeFromTime(tnow),
	}

	xd := tnow.Format("2006-01-02")
	xt := tnow.Format("15:04:05")

	djson, err := json.Marshal(asst)
	if err != nil {
		t.Errorf("failed to marshal struct of calendar values: %v", err)
	}

	type Check struct {
		Date string
		NullDate string
		Time string
		NullTime string
	}

	dasst := Check{}
	err = json.Unmarshal(djson, &dasst)
	if err != nil {
		t.Errorf("failed to unmarshal struct of calendar values: %v", err)
	}

	if dasst.Date != xd {
		t.Errorf("struct Date marshal/unmarshal failed, expected %s got %s", xd, dasst.Date)
	}
	if dasst.NullDate != xd {
		t.Errorf("struct NullDate marshal/unmarshal failed, expected %s got %s", xd, dasst.NullDate)
	}
	if dasst.Time != xt {
		t.Errorf("struct Time marshal/unmarshal failed, expected %s got %s", xd, dasst.Time)
	}
	if dasst.NullTime != xt {
		t.Errorf("struct NullTime marshal/unmarshal failed, expected %s got %s", xd, dasst.NullTime)
	}

}

func TestTimeIgnored(t *testing.T) {
	d1 := calendar.DateFromTime(time.Date(2016, 10, 4, 1, 0, 0, 0, time.UTC))
	d2 := calendar.DateFromTime(time.Date(2016, 10, 4, 21, 0, 0, 0, time.UTC))
	if d2.After(d1) {
		t.Errorf("%v after %v (when created from times with hour difference)", d2, d1)
	}
	if !d2.Equal(d1) {
		t.Errorf("%v not equal to %v (when created from times with hour difference)", d2, d1)
	}
	if d2.Before(d1) {
		t.Errorf("%v before %v (when created from times with hour difference)", d2, d1)
	}
}

func TestComparisons(t *testing.T) {
	d1 := calendar.NewDate(2016, 10, 15)
	d2 := calendar.NewDate(2016, 10, 16)
	if !d1.Before(d2) {
		t.Errorf("%v before %v returned false, expected true", d1, d2)
	}
	if d1.Equal(d2) {
		t.Errorf("%v equal %v returned true, expected false", d1, d2)
	}
	if d1.After(d2) {
		t.Errorf("%v after %v returned true, expected false", d1, d2)
	}
	if d1.After(d1) {
		t.Errorf("%v after itself returned true, expected false", d1)
	}
	if d1.Before(d1) {
		t.Errorf("%v before itself returned true, expected false", d1)
	}
}

func TestAdd(t *testing.T) {
	d1 := calendar.NewDate(2016, 2, 28)
	d2 := d1.AddDate(0, 0, 2)
	d3 := calendar.NewDate(2016, 3, 1)
	if !d2.Equal(d3) {
		t.Errorf("%v +2 days returned %v, expected %v", d1, d2, d3)
	}
}
