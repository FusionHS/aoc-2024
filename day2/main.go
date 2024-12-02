package main

import (
	"day2/answer"
	"day2/util"
	"fmt"
)

func main() {

	reports, _ := util.LoadReports()

	result := answer.GetAnswerA(reports)

	fmt.Println("Answer A: " + result)

	result = answer.GetAnswerB(reports)

	fmt.Println("Answer B: " + result)

}
