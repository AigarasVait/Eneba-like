package db

func Init() {
	Connect()
	RunMigrations()
}
