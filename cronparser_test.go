package main

import "testing"

func TestCronParserParse(t *testing.T) {
	cron := "*/15 0 1,15 * 1-5 /usr/bin/find"
	min := "0 15 30 45"
	hour := "0"
	dom := "1 15"
	mon := "1 2 3 4 5 6 7 8 9 10 11 12"
	dow := "1 2 3 4 5"
	com := "/usr/bin/find"
	cp := NewCronParser()
	cp.Parse(cron)
	if cp.GetOutput(0, 59, cp.minute) != min {
		t.Errorf("Expected %s, got %s", min, cp.GetOutput(0, 59, cp.minute))
	}
	if cp.GetOutput(0, 23, cp.hour) != hour {
		t.Errorf("Expected %s, got %s", hour, cp.GetOutput(0, 23, cp.hour))
	}
	if cp.GetOutput(1, 31, cp.dayOfMonth) != dom {
		t.Errorf("Expected %s, got %s", dom, cp.GetOutput(1, 31, cp.dayOfMonth))
	}
	if cp.GetOutput(1, 12, cp.month) != mon {
		t.Errorf("Expected %s, got %s", mon, cp.GetOutput(1, 12, cp.month))
	}
	if cp.GetOutput(1, 7, cp.dayOfWeek) != dow {
		t.Errorf("Expected %s, got %s", dow, cp.GetOutput(1, 7, cp.dayOfWeek))
	}
	if cp.command != com {
		t.Errorf("Expected %s, got %s", com, cp.command)
	}
}
