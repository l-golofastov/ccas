package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	path := os.Args[1]
	Run(path)
}

func Run(path string) string {
	data, err := parseFile(path)

	if err != "" {
		fmt.Println(err)
		return err
	}

	return makeReport(data)
}

func makeReport(data []string) string {
	res := ""

	tablesNum, _ := strconv.Atoi(data[0])

	workTimes := strings.Split(data[1], " ")
	opened, _ := parseTime(workTimes[0])
	closed, _ := parseTime(workTimes[1])

	price, _ := parsePositiveInt(data[2])

	clientsIn := make(map[string]bool)

	clientsSit := make(map[string]int)

	tables := make(map[int]string)

	queue := make([]string, 0)

	tablesIncome := make(map[int]int)
	for i := 1; i < tablesNum+1; i++ {
		tablesIncome[i] = 0
	}

	tablesBusiness := make(map[int]int)
	for i := 1; i < tablesNum+1; i++ {
		tablesBusiness[i] = 0
	}

	fmt.Println(workTimes[0])
	res += workTimes[0] + "\n"

	for _, eventStr := range data[3:] {
		fmt.Println(eventStr)
		res += eventStr + "\n"

		event, _ := parseEvent(eventStr, tablesNum)
		timeStr := event[0]
		eventId, _ := parsePositiveInt(event[1])
		name := event[2]

		if eventId == 1 {
			eventTime, _ := parseTime(timeStr)
			if visited := clientsIn[name]; visited {
				msg := generateError(timeStr, "YouShallNotPass")
				res += msg + "\n"
			} else if eventTime.Sub(opened).Minutes() < 0 || closed.Sub(eventTime).Minutes() < 0 {
				msg := generateError(timeStr, "NotOpenYet")
				res += msg + "\n"
			} else {
				clientsIn[name] = true
			}
		} else if eventId == 2 {
			tableId, _ := parsePositiveInt(event[3])

			if _, ok := tables[tableId]; ok {
				msg := generateError(timeStr, "PlaceIsBusy")
				res += msg + "\n"
			} else if visited := clientsIn[name]; !visited {
				msg := generateError(timeStr, "ClientUnknown")
				res += msg + "\n"
			} else {
				if id, ok := clientsSit[name]; ok {
					tablesIncome[id] += countIncome(tables[id], timeStr, price)
					tablesBusiness[id] += countBusiness(tables[id], timeStr)
					delete(tables, id)
				}

				clientsSit[name] = tableId
				tables[tableId] = timeStr
			}
		} else if eventId == 3 {
			if len(tables) < tablesNum {
				msg := generateError(timeStr, "ICanWaitNoLonger!")
				res += msg + "\n"
			} else if len(queue) > tablesNum {
				msg := generateLeave(timeStr, name)
				res += msg + "\n"
				delete(clientsIn, name)
			} else {
				alreadyWaiting := false
				for _, elem := range queue {
					if elem == name {
						alreadyWaiting = true
					}
				}
				if !alreadyWaiting {
					queue = append(queue, name)
				}
			}
		} else if eventId == 4 {
			if visited := clientsIn[name]; !visited {
				msg := generateError(timeStr, "ClientUnknown")
				res += msg + "\n"
			} else {
				if table, ok := clientsSit[name]; ok {
					tablesIncome[table] += countIncome(tables[table], timeStr, price)
					tablesBusiness[table] += countBusiness(tables[table], timeStr)
					delete(clientsSit, name)
					if len(queue) > 0 {
						newClient := queue[0]
						clientsSit[newClient] = table
						tables[table] = timeStr
						queue = queue[1:]
						msg := generateSit(timeStr, newClient, table)
						res += msg + "\n"
					} else {
						delete(tables, table)
					}
				}
				delete(clientsIn, name)
			}
		}
	}

	keys := make([]string, 0, len(clientsIn))

	for k := range clientsIn {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, client := range keys {
		msg := generateLeave(workTimes[1], client)
		res += msg + "\n"
		if table, ok := clientsSit[client]; ok {
			tablesIncome[table] += countIncome(tables[table], workTimes[1], price)
			tablesBusiness[table] += countBusiness(tables[table], workTimes[1])
			delete(tables, table)
		}
		delete(clientsIn, client)
	}

	fmt.Println(workTimes[1])
	res += workTimes[1] + "\n"

	for i := 1; i < tablesNum+1; i++ {
		msg := generateTableInfo(tablesIncome[i], tablesBusiness[i], i)
		res += msg + "\n"
	}

	return strings.TrimSuffix(res, "\n")
}

func countIncome(prev string, curr string, price int) int {
	start, _ := parseTime(prev)
	stop, _ := parseTime(curr)
	minutes := int(stop.Sub(start).Minutes())
	hours := 0

	if minutes%60 != 0 {
		hours = minutes/60 + 1
	} else {
		hours = minutes / 60
	}

	return hours * price
}

func countBusiness(start string, stop string) int {
	startTime, _ := parseTime(start)
	stopTime, _ := parseTime(stop)

	return int(stopTime.Sub(startTime).Minutes())
}
