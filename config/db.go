package configs

import (
	"fmt"
	"os"
 

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	
	"github.com/nuchit2019/assessment-tax/model"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	 
)

var database *gorm.DB
var err error

var echoInstance *echo.Echo

func Echo() *echo.Echo {
    return echoInstance
}

func SetEcho(e *echo.Echo) {
    echoInstance = e
}

func InitDB() {

	err = godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}


	// host := os.Getenv("DB_HOST") //"localhost"
	// user := os.Getenv("DB_USER") // "postgres"
	// password := os.Getenv("DB_PASSWORD") // "postgres"
	// dbName := os.Getenv("DB_NAME") // "db"
	// port := os.Getenv("DB_PORT") // "5432"

	// dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai", host, user, password, dbName, port)

	// database, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	// if err != nil {
	// 	panic("Failed to initialize database ERROR: "+err.Error())
	// }

	databaseURL := os.Getenv("DATABASE_URL")
	database, err = gorm.Open(postgres.Open(databaseURL), &gorm.Config{})
	if err != nil {
		panic("Failed to initialize database ERROR: " + err.Error())
	}


	fmt.Println("----------------------------")
	fmt.Println("Database connected...")

	database.AutoMigrate(&model.Product{})
	fmt.Println("AutoMigrate Product...") 

	fmt.Println("----------------------------")
}

func DB() *gorm.DB {
	return database
}

func ApiPort() string {
	return os.Getenv("PORT")
}

func ApiAdmin() string {
	return os.Getenv("ADMIN_USERNAME")
}
func ApiAdminPass() string {
	return os.Getenv("ADMIN_PASSWORD")
}