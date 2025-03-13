package partyrobot

import "fmt"

// Welcome greets a person by name.
func Welcome(name string) string {
	greeting := fmt.Sprintf("Welcome to my party, %s!", name)
	fmt.Println(greeting)
	return greeting
}

// HappyBirthday wishes happy birthday to the birthday person and exclaims their age.
func HappyBirthday(name string, age int) string {
	return fmt.Sprintf("Happy birthday %s! You are now %v years old!", name, age)
}

// AssignTable assigns a table to each guest.
func AssignTable(name string, table int, neighbor, direction string, distance float64) string {
	var greeting = Welcome(name) + "\n"
	var tableAss = fmt.Sprintf("You have been assigned to table %03d. Your table is %s, exactly %.1f meters from here.\n", table, direction, distance)
	var nextTo = fmt.Sprintf("You will be sitting next to %s.", neighbor)
	return greeting + tableAss + nextTo
}
