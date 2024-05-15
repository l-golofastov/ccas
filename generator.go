package main

import (
	"fmt"
	"strconv"
)

func GenerateError(errorTime string, msg string) string {
	eventError := errorTime + " 13 " + msg
	fmt.Println(eventError)

	return eventError
}

func GenerateLeave(leavingTime string, client string) string {
	leave := leavingTime + " 11 " + client
	fmt.Println(leave)

	return leave
}

func GenerateSit(sitTime string, name string, table int) string {
	tableStr := strconv.Itoa(table)
	sit := sitTime + " 12 " + name + " " + tableStr
	fmt.Println(sit)

	return sit
}

func GenerateTableInfo(income int, business int, id int) string {
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
