package api

func (a *API) registerRoutes() {
	a.router.POST("/game", a.Start)
	a.router.POST("/game/play", a.Play)
}
