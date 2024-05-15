package main

import (
	"fmt"
	"strconv"
)

func generateError(errorTime string, msg string) string {
	eventError := errorTime + " 13 " + msg
	fmt.Println(eventError)

	return eventError
}

func generateLeave(leavingTime string, client string) string {
	leave := leavingTime + " 11 " + client
	fmt.Println(leave)

	return leave
}

func generateSit(sitTime string, name string, table int) string {
	tableStr := strconv.Itoa(table)
	sit := sitTime + " 12 " + name + " " + tableStr
	fmt.Println(sit)

	return sit
}

func generateTableInfo(income int, business int, id int) string {
	idStr := strconv.Itoa(id)
	incomeStr := strconv.Itoa(income)

	hoursStr := ""
	hours := business / 60
	if hours >= 10 {
		hoursStr = strconv.Itoa(hours)
	} else {
		hoursStr = "0" + strconv.Itoa(hours)
	}

	minutesStr := ""
	minutes := business % 60
	if minutes >= 10 {
		minutesStr = strconv.Itoa(minutes)
	} else {
		minutesStr = "0" + strconv.Itoa(minutes)
	}

	time := hoursStr + ":" + minutesStr

	info := idStr + " " + incomeStr + " " + time

	fmt.Println(info)

	return info
}
