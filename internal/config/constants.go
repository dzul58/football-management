package config

// Player positions
const (
	PositionPenyerang     = "Penyerang"
	PositionGelandang     = "Gelandang"
	PositionBertahan      = "Bertahan"
	PositionPenjagaGawang = "Penjaga Gawang"
)

// Match statuses
const (
	MatchStatusScheduled = "Scheduled"
	MatchStatusCompleted = "Completed"
	MatchStatusCancelled = "Cancelled"
)

// Match results
const (
	MatchResultHomeWin = "Tim Home Menang"
	MatchResultAwayWin = "Tim Away Menang"
	MatchResultDraw    = "Draw"
)

// ValidPlayerPositions returns all valid player positions
func ValidPlayerPositions() []string {
	return []string{
		PositionPenyerang,
		PositionGelandang,
		PositionBertahan,
		PositionPenjagaGawang,
	}
}

// ValidMatchStatuses returns all valid match statuses
func ValidMatchStatuses() []string {
	return []string{
		MatchStatusScheduled,
		MatchStatusCompleted,
		MatchStatusCancelled,
	}
}
