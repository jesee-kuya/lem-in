package colony

import (
	"fmt"
	"strings"
)

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

func Path(routes [][]int, numberOfAnts int) {
	antPositions := make([]int, numberOfAnts+1)
	antsFinished := 0

	// Iterate turn by turn
	for i := 0; antsFinished < numberOfAnts; i++ {
		var moves []string

		for ant := 1; ant <= numberOfAnts; ant++ {
			if antPositions[ant] == -1 {
				continue // Ant is already at ##end
			}

			for _, route := range routes {
				currentPosition := antPositions[ant]

				if currentPosition < len(route)-1 { // Ant has not finished the path
					nextRoom := route[currentPosition+1]

					if !isRoomOccupied(antPositions, nextRoom) || nextRoom == route[len(route)-1] {
						antPositions[ant]++
						moves = append(moves, fmt.Sprintf("L%d-%d", ant, nextRoom))

						if nextRoom == route[len(route)-1] {
							antPositions[ant] = -1 // Marks ant as finished
							antsFinished++
						}
						break
					}
				}
			}
		}

		if len(moves) > 0 {
			fmt.Println(strings.Join(moves, " "))
		}
	}
}

func isRoomOccupied(antPositions []int, room int) bool {
	for _, position := range antPositions {
		if position == room {
			return true
		}
	}
	return false
}
