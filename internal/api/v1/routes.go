package v1

func (v *V1) registerRoutes() {
	v.router.POST("/game/start", v.gameStart)
	v.router.POST("/game/guess", v.gameGuess)
}
