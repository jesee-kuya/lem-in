package colony

import (
	"fmt"
	"strings"
)

type PathInfo struct {
	path             []int
	length, antCount int
}

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
	if len(routes) == 0 || numberOfAnts <= 0 {
		return
	}

	pathInfos := calculateOptimalDistribution(routes, numberOfAnts)

	type AntState struct {
		pathIndex, position int
		isFinished          bool
	}

	antStates := make([]AntState, numberOfAnts)
	currentAnt := 1
	var antsFinished int

	// Assign initial paths to ants
	for i := range pathInfos {
		for j := 0; j < pathInfos[i].antCount; j++ {
			if currentAnt <= numberOfAnts {
				antStates[currentAnt] = AntState{
					pathIndex:  i,
					position:   -1,
					isFinished: false,
				}
				currentAnt++
			}
		}
	}

	// Process moves turn by turn
	for antsFinished < numberOfAnts {
		moves := make([]string, 0)
		occupiedRooms := make(map[int]bool)

		for ant := 1; ant <= numberOfAnts; ant++ {
			if antStates[ant].isFinished {
				continue
			}

			state := &antStates[ant]
			path := pathInfos[state.pathIndex].path

			// Check if ant can start if it hasn't started yet
			if state.position == -1 {
				nextRoom := path[1]

				if !occupiedRooms[nextRoom] {
					state.position = 1
					occupiedRooms[nextRoom] = true
					moves = append(moves, fmt.Sprintf("L%d-%d", ant, nextRoom))
				}
				continue
			}

			// Try moving forward if ant is on the path
			if state.position < len(path)-1 {
				nextRoom := path[state.position+1]

				if !occupiedRooms[nextRoom] {
					state.position++
					occupiedRooms[nextRoom] = true
					moves = append(moves, fmt.Sprintf("L%d-%d", ant, nextRoom))

					// Check if ant has reached the end
					if state.position == len(path)-1 {
						state.isFinished = true
						antsFinished++
					}
				}
			}
		}

		if len(moves) > 0 {
			fmt.Println(strings.Join(moves, " "))
		}
	}
}

func assignRoutes(numberOfAnts int, routes [][]int) map[int][]int {
	assignedRoutes := make(map[int][]int)

	for i := 1; i <= numberOfAnts; i++ {
		assignedRoutes[i] = routes[(i-1)%len(routes)]
	}
	return assignedRoutes
}

func isRoomOccupied(antPositions []int, room int) bool {
	for _, position := range antPositions {
		if position == room {
			return true
		}
	}
	return false
}
