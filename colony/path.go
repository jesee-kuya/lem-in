package colony

import "fmt"

func ParseAnts(lines []string) int {
	if len(lines) == 0 {
		return 0
	}

	var ants int
	_, err := fmt.Sscanf(lines[0], "%d", &ants)
	if err != nil {
		fmt.Println("Error parsing number of ants:", err)
		return 0
	}
	return ants
}

func isRoomOccupied(antPositions []int, room int) bool {
	for _, position := range antPositions {
		if position == room {
			return true
		}
	}
	return false
}
