package main 

import (
	"fmt"
	"strconv"
	"strings"
)

type CronParser struct {
	minute, hour, dayOfMonth, month, dayOfWeek, command string
}

func NewCronParser() *CronParser {
	return &CronParser{}
}

func (c *CronParser) Parse(cron string) error {
	cronParts := strings.Fields(cron)
	if len(cronParts) != 6 {
		return fmt.Errorf("invalid cron string")
	}

	c.minute = cronParts[0]
	c.hour = cronParts[1]
	c.dayOfMonth = cronParts[2]
	c.month = cronParts[3]
	c.dayOfWeek = cronParts[4]
	c.command = cronParts[5]

	return nil
}

func (c *CronParser) Print() {
	fmt.Printf("minute        %s\n", c.GetOutput(0, 59, c.minute))
	fmt.Printf("hour          %s\n", c.GetOutput(0, 23, c.hour))
	fmt.Printf("day of month  %s\n", c.GetOutput(1, 31, c.dayOfMonth))
	fmt.Printf("month         %s\n", c.GetOutput(1, 12, c.month))
	fmt.Printf("day of week   %s\n", c.GetOutput(1, 7, c.dayOfWeek))
	fmt.Printf("command       %s\n", c.command)
}

func (c *CronParser) GetOutput(start int, end int, val string) string {
	ret := ""
	if val == "*" {
		return PrintRange(start, end)
	}

	if strings.Contains(val, "-") {
		parts := strings.Split(val, "-")
		start, _ = strconv.Atoi(parts[0])
		end, _ = strconv.Atoi(parts[1])
		return PrintRange(start, end)
	}

	if strings.Contains(val, ",") {
		parts := strings.Split(val, ",")
		for _, part := range parts {
			val, _ := strconv.Atoi(part)
			ret += fmt.Sprintf("%d ", val)
		}
		return ret[:len(ret)-1]
	}

	if strings.Contains(val, "/") {
		parts := strings.Split(val, "/")
		interval, _ := strconv.Atoi(parts[1])
		return PrintIntervals(start, end, interval)
	}

	return val
}
