package main

import (
	"errors"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func parseFile(path string) ([]string, string) {
	data, errFile := os.ReadFile(path)
	if errFile != nil {
		panic(errFile)
	}

	dataSlice := strings.Split(string(data), "\n")

	if len(dataSlice) < 3 {
		return dataSlice, "file data length is less than 3 lines"
	}

	previousTime, _ := parseTime("00:00")

	for i, line := range dataSlice {
		if i == 0 {
			if num, errTables := parsePositiveInt(line); errTables != nil || num == -1 {
				return []string{}, line
			}
		} else if i == 1 {
			lineSplit := strings.Split(line, " ")
			if len(lineSplit) != 2 {
				return []string{}, line
			}

			opened, err1 := parseTime(lineSplit[0])
			closed, err2 := parseTime(lineSplit[1])

			if err1 != nil || err2 != nil || closed.Sub(opened).Minutes() < 0 {
				return []string{}, line
			}
		} else if i == 2 {
			if num, errPrice := parsePositiveInt(line); errPrice != nil || num == -1 {
				return []string{}, line
			}
		} else {
			tablesNum, _ := strconv.Atoi(dataSlice[0])
			eventSplit, errEvent := parseEvent(line, tablesNum)
			if errEvent != nil || len(eventSplit) == 0 {
				return []string{}, line
			}

			currentTime, _ := parseTime(eventSplit[0])

			if currentTime.Sub(previousTime).Minutes() < 0 {
				return []string{}, line
			}

			previousTime = currentTime
		}
	}
	return dataSlice, ""
}

func parseTime(s string) (time.Time, error) {
	value, err := time.Parse("15:04", s)

	return value, err
}

func parsePositiveInt(s string) (int, error) {
	num, err := strconv.Atoi(s)
	if err != nil || num <= 0 {
		return -1, err
	}
	return num, nil
}

func parseClient(s string) (string, error) {
	validRegex := regexp.MustCompile(`[a-z0-9_-]+`)

	emptyClient := validRegex.ReplaceAllString(s, "")
	if emptyClient != "" {
		return "", errors.New("client name format invalid")
	}

	return s, nil
}

func parseEvent(s string, tablesNum int) ([]string, error) {
	split := strings.Split(s, " ")
	if len(split) != 3 && len(split) != 4 {
		return []string{}, errors.New("event doesn't contain 3 or 4 elements")
	}

	_, errTime := parseTime(split[0])
	if errTime != nil {
		return []string{}, errTime
	}

	id, errEvent := parsePositiveInt(split[1])
	if errEvent != nil || id == -1 || id > 4 {
		return []string{}, errEvent
	}

	_, errClient := parseClient(split[2])
	if errClient != nil {
		return []string{}, errClient
	}

	if len(split) == 4 {
		if tableId, errTable := parsePositiveInt(split[3]); errTable != nil || tableId == -1 ||
			tableId > tablesNum || id != 2 {
			return []string{}, errors.New("table id format invalid")
		}
	}

	return split, nil
}
