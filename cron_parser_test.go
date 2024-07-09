package main

import (
	"testing"
)

func TestParseValidCronString(t *testing.T) {
	parser := &SimpleCronParser{}
	cronString := "*/15 0 1,15 * 1-5 /usr/bin/find"
	expected := &CronFields{
		Minute:     "minute         0 15 30 45",
		Hour:       "hour           0",
		DayOfMonth: "day of month   1 15",
		Month:      "month          1 2 3 4 5 6 7 8 9 10 11 12",
		DayOfWeek:  "day of week    1 2 3 4 5",
		Command:    "command        /usr/bin/find",
	}

	parsedFields, err := parser.Parse(cronString)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if parsedFields.Minute != expected.Minute ||
		parsedFields.Hour != expected.Hour ||
		parsedFields.DayOfMonth != expected.DayOfMonth ||
		parsedFields.Month != expected.Month ||
		parsedFields.DayOfWeek != expected.DayOfWeek ||
		parsedFields.Command != expected.Command {
		t.Errorf("parsed fields do not match expected fields.\nParsed: %+v\nExpected: %+v", parsedFields, expected)
	}
}

func TestParseInvalidCronString(t *testing.T) {
	parser := &SimpleCronParser{}
	cronString := "*/15 0 1,15 * 1-5"
	_, err := parser.Parse(cronString)
	if err == nil {
		t.Fatalf("expected an error for invalid cron string format, got nil")
	}
}

func TestParseComplexCronString(t *testing.T) {
	parser := &SimpleCronParser{}
	cronString := "5-10/2 8-18/2 * 1,6,12 0 /bin/true"
	expected := &CronFields{
		Minute:     "minute         5 7 9",
		Hour:       "hour           8 10 12 14 16 18",
		DayOfMonth: "day of month   1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18 19 20 21 22 23 24 25 26 27 28 29 30 31",
		Month:      "month          1 6 12",
		DayOfWeek:  "day of week    0",
		Command:    "command        /bin/true",
	}

	parsedFields, err := parser.Parse(cronString)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if parsedFields.Minute != expected.Minute ||
		parsedFields.Hour != expected.Hour ||
		parsedFields.DayOfMonth != expected.DayOfMonth ||
		parsedFields.Month != expected.Month ||
		parsedFields.DayOfWeek != expected.DayOfWeek ||
		parsedFields.Command != expected.Command {
		t.Errorf("parsed fields do not match expected fields.\nParsed: %+v\nExpected: %+v", parsedFields, expected)
	}
}
