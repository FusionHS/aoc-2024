package answer

import (
	"golang.org/x/exp/slices"
	"strconv"
)

type Direction int

const MinDiff = 1
const MaxDiff = 3

const (
	Ascending Direction = iota
	Descending
)

func GetAnswerA(reports [][]int) string {
	return getAnswer(reports, 0)
}
func GetAnswerB(reports [][]int) string {
	return getAnswer(reports, 1)
}

func getAnswer(reports [][]int, allowedFaults int) string {

	safeCount := 0
	for _, row := range reports {
		faults := getFaults(row)

		if faults <= allowedFaults {
			safeCount++
		}
	}

	return strconv.Itoa(safeCount)
}

func getFaults(row []int) int {

	length := len(row)
	if length < 2 {
		return 0
	}

	var direction Direction

	if row[1] > row[0] {
		direction = Ascending
	} else {
		direction = Descending

	}

	for col := 1; col < length; col++ {

		difference := absInt(row[col-1] - row[col])

		if difference < MinDiff || difference > MaxDiff ||
			direction == Ascending && row[col] < row[col-1] ||
			direction == Descending && row[col] > row[col-1] {

			faults := make([]int, 0, length)
			for i := 0; i < length; i++ {
				faults = append(faults, getFaults(removeElement(row, i)))
			}

			return 1 + slices.Min(faults)
		}

	}
	return 0
}

func absInt(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func removeElement(slice []int, index int) []int {
	copySlice := make([]int, len(slice))
	copy(copySlice, slice)

	if index < 0 || index >= len(slice) {
		return slice
	}

	return append(copySlice[:index], copySlice[index+1:]...)
}
