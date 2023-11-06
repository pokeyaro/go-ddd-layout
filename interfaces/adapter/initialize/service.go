package initialize

import (
	"os"
	"strconv"

	userAppSrv "server/application/service/user"
	userDomainSrv "server/domain/user/service"
	"server/infrastructure/database"
	"server/infrastructure/persistence"
	"server/interfaces/controller"
)

func serviceRegister() {
	// Get environment variables
	dbHost := os.Getenv("DB_HOST")
	dbPort, _ := strconv.Atoi(os.Getenv("DB_PORT"))
	dbUser := os.Getenv("DB_USER")
	dbPasswd := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	// Initialize the database
	db := database.InitDB(dbHost, dbPort, dbUser, dbPasswd, dbName)

	// Repository implementation
	repos := persistence.NewRepositories(db)

	// Register services
	// Dependency Injection:
	// By injecting UserRepository into NewUserDomainImpl and then injecting UserDomain into NewServiceImpl,
	// the dependency relationship is passed and decoupled.
	appSrvUser := userAppSrv.NewServiceImpl(userDomainSrv.NewUserDomainImpl(repos.User))

	// Inject services into controllers
	controller.InitSrvInject(
		appSrvUser,
	)
}
