package colony

func isRoomOccupied(antPositions []int, room int) bool {
	for _, position := range antPositions {
		if position == room {
			return true
		}
	}
	return false
}
