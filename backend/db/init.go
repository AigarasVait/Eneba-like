package db

func Init() {
	Connect()
	RunMigrations()
	SeedDatabase(DB)
}
