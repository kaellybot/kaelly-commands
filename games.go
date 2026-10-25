package commands

// game holds what differs between the command sets of each game.
type game struct {
	// name fills the {{ .game }} placeholder of the descriptions.
	name              string
	alignmentMaxLevel int
	jobMaxLevel       int
}

func getDofusGame() game {
	return game{name: "DOFUS", alignmentMaxLevel: AlignmentMaxLevel, jobMaxLevel: JobMaxLevel}
}

func getTouchGame() game {
	return game{name: "DOFUS Touch", alignmentMaxLevel: AlignmentTouchMaxLevel, jobMaxLevel: JobTouchMaxLevel}
}

func getRetroGame() game {
	return game{name: "DOFUS Retro", alignmentMaxLevel: AlignmentRetroMaxLevel, jobMaxLevel: JobRetroMaxLevel}
}
