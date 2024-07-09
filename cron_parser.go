package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// CronParser is an interface for parsing cron strings
type CronParser interface {
	Parse(cronString string) (*CronFields, error)
}

type CronFields struct {
	Minute     string
	Hour       string
	DayOfMonth string
	Month      string
	DayOfWeek  string
	Command    string
}

// SimpleCronParser is a struct that implements the CronParser interface
type SimpleCronParser struct{}

func expandField(field, fieldName string, min, max int) string {
	if field == "*" {
		values := make([]string, max-min+1)
		for i := min; i <= max; i++ {
			values[i-min] = strconv.Itoa(i)
		}
		return fmt.Sprintf("%-14s %s", fieldName, strings.Join(values, " "))
	}

	expanded := []string{}
	for _, part := range strings.Split(field, ",") {
		if strings.Contains(part, "/") {
			subParts := strings.Split(part, "/")
			rangePart, step := subParts[0], subParts[1]
			stepInt, _ := strconv.Atoi(step)

			if rangePart == "*" {
				for i := min; i <= max; i += stepInt {
					expanded = append(expanded, strconv.Itoa(i))
				}
			} else {
				rangeParts := strings.Split(rangePart, "-")
				start, _ := strconv.Atoi(rangeParts[0])
				end, _ := strconv.Atoi(rangeParts[1])
				for i := start; i <= end; i += stepInt {
					expanded = append(expanded, strconv.Itoa(i))
				}
			}
		} else if strings.Contains(part, "-") {
			rangeParts := strings.Split(part, "-")
			start, _ := strconv.Atoi(rangeParts[0])
			end, _ := strconv.Atoi(rangeParts[1])
			for i := start; i <= end; i++ {
				expanded = append(expanded, strconv.Itoa(i))
			}
		} else {
			expanded = append(expanded, part)
		}
	}

	return fmt.Sprintf("%-14s %s", fieldName, strings.Join(expanded, " "))
}

func (p *SimpleCronParser) Parse(cronString string) (*CronFields, error) {
	fields := strings.Fields(cronString)
	if len(fields) != 6 {
		return nil, fmt.Errorf("invalid cron string format")
	}

	minuteField := fields[0]
	hourField := fields[1]
	dayOfMonthField := fields[2]
	monthField := fields[3]
	dayOfWeekField := fields[4]
	commandField := fields[5]

	return &CronFields{
		Minute:     expandField(minuteField, "minute", 0, 59),
		Hour:       expandField(hourField, "hour", 0, 23),
		DayOfMonth: expandField(dayOfMonthField, "day of month", 1, 31),
		Month:      expandField(monthField, "month", 1, 12),
		DayOfWeek:  expandField(dayOfWeekField, "day of week", 0, 6),
		Command:    fmt.Sprintf("%-14s %s", "command", commandField),
	}, nil
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run main.go \"<cron string>\"")
		return
	}

	cronString := os.Args[1]
	parser := &SimpleCronParser{}
	parsedFields, err := parser.Parse(cronString)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(parsedFields.Minute)
	fmt.Println(parsedFields.Hour)
	fmt.Println(parsedFields.DayOfMonth)
	fmt.Println(parsedFields.Month)
	fmt.Println(parsedFields.DayOfWeek)
	fmt.Println(parsedFields.Command)
}
