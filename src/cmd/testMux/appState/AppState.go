package appstate

var Games []*GameStateManager = make([]*GameStateManager, 0)

func FindGameById(id string) *GameStateManager {
	for _, v := range Games {
		if v.ID == id {
			return v
		}
	}

	return nil
}

func ListGameIds() *[]string {
	var gameIds = make([]string, len(Games))

	for i, gi := range Games {
		gameIds[i] = gi.ID
	}

	return &gameIds
}

func StartNewGame(ownerId string) *GameStateManager {
	var newGame = NewGameStateManager(ownerId)
	Games = append(Games, newGame)
	return newGame
}
