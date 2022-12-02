package main

type roundResult int8

const (
	lose roundResult = iota
	draw
	win
)

type enemyMove byte

const (
	enemyMoveRock      enemyMove = 'A'
	enemyMovePaper     enemyMove = 'B'
	enemyMoveScrissors enemyMove = 'C'
)

type myMove byte

const (
	myMoveRock      myMove = 'X'
	myMovePaper     myMove = 'Y'
	myMoveScrissors myMove = 'Z'
)

type actionToTake byte

const (
	actionLose actionToTake = 'X'
	actionDraw              = 'Y'
	actionWin               = 'Z'
)

var moveScore = map[myMove]int{
	myMoveRock:      1,
	myMovePaper:     2,
	myMoveScrissors: 3,
}

var roundResultScore = map[roundResult]int{
	lose: 0,
	draw: 3,
	win:  6,
}

var actionScore = map[actionToTake]int{
	actionLose: roundResultScore[lose],
	actionDraw: roundResultScore[draw],
	actionWin:  roundResultScore[win],
}

func playRound(enemy enemyMove, me myMove) roundResult {
	if enemy == enemyMoveRock {
		switch me {
		case myMovePaper:
			return win
		case myMoveScrissors:
			return lose
		default:
			return draw
		}
	} else if enemy == enemyMovePaper {
		switch me {
		case myMoveRock:
			return lose
		case myMoveScrissors:
			return win
		default:
			return draw
		}
	} else {
		switch me {
		case myMoveRock:
			return win
		case myMovePaper:
			return lose
		default:
			return draw
		}
	}
}

func playRoundSmart(enemy enemyMove, action actionToTake) myMove {
	if enemy == enemyMoveRock {
		switch action {
		case actionWin:
			return myMovePaper
		case actionLose:
			return myMoveScrissors
		default:
			return myMoveRock
		}
	} else if enemy == enemyMovePaper {
		switch action {
		case actionWin:
			return myMoveScrissors
		case actionLose:
			return myMoveRock
		default:
			return myMovePaper
		}
	} else {
		switch action {
		case actionWin:
			return myMoveRock
		case actionLose:
			return myMovePaper
		default:
			return myMoveScrissors
		}
	}
}
