package app

func Run() {
	dbConnectAndMigrate()
	serverInitAndStart()
}
