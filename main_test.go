package main

import "testing"

func TestCheckZip(t *testing.T) {
	got := checkZip("981")
	want := "00981"
	if got != want {
		t.Errorf("Zipcode is invalid, got: %s, want: %s.", got, want)
	}
}

func TestUpcaseFn(t *testing.T) {
	got := upcaseFn("Yeni Capote Diaz")
	want := "YENI CAPOTE DIAZ"
	if got != want {
		t.Errorf("Fullname was not upcased, got: %s, want: %s.", got, want)
	}
}

func TestChangeTime(t *testing.T) {
	got := changeTime("4/1/11 11:00:00 AM")
	want := "2011-04-01 14:00:00 -0400 EDT"
	if got != want {
		t.Errorf("Timestamp was not converted, got: %s, want: %s.", got, want)
	}
}

func TestPadTimestamp(t *testing.T) {
	got := padTimestamp("4/1/11 11:00:00 AM")
	want := "04/01/11 11:00:00 AM"
	if got != want {
		t.Errorf("Timestamp was not converted, got: %s, want: %s.", got, want)
	}
}

func TestPadZeros(t *testing.T) {
	got := padZeros("4/1/11", "/")
	want := "04/01/11"
	if got != want {
		t.Errorf("Timestamp was not converted, got: %s, want: %s.", got, want)
	}
}

func TestDurationSecs(t *testing.T) {
	got := durationSecs("1:23:32.123")
	want := 5012.123
	if got != want {
		t.Errorf("Timestamp was not converted, got: %v, want: %v.", got, want)
	}
}
