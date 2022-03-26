package ports

type API interface {
	Start()
	ListenAndServe()
	initRoutes()
}
