package server

func Start() {

	// Connect to DB
	SetConnection(DBOptions())

	// Set Routers
	router := setRouter()

	// Start listening and serving requests
	router.Run(":8000")
}
