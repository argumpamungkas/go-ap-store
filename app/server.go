package app

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// server.go akan dipanggil setelah main.go
// digunakan untuk migrate db/connect db/
type Server struct {
	DB     *gorm.DB
	Router *mux.Router
}

type AppConfig struct {
	AppName string
	AppEnv  string
	AppPort string
}

type DBConfig struct {
	DBHost     string
	DBUser     string
	DBPassword string
	DBName     string
	DBPort     string
	DBDriver   string
}

// method default config
func getEnv(key, callback string) string {

	// cek apakah key yang dikirim ada nilainya ?
	// jika ada, kembalikan value nya
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	// jika tidak ada kembalikan nilai default
	return callback
}

// method init
func (server *Server) initialize(appConfig AppConfig, dbConfig DBConfig) {
	fmt.Println("Call initialize from " + appConfig.AppName + " in " + appConfig.AppEnv)

	server.initializeDB(dbConfig)
	server.initializeRoutes()
}

func (server *Server) initializeDB(dbConfig DBConfig) {
	var err error

	if dbConfig.DBDriver == "mysql" {
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbConfig.DBUser, dbConfig.DBPassword, dbConfig.DBHost, dbConfig.DBPort, dbConfig.DBName)
		server.DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	} else {
		// digunakan untuk driver lain
	}

	if err != nil {
		panic("failed to connect db " + err.Error())
	}

	for _, value := range RegisterModels() {
		// fmt.Println("VALUE =>", value.Model)
		err = server.DB.Debug().AutoMigrate(&value.Model)

		if err != nil {
			log.Fatal(err)
		}
	}

	fmt.Println("Migrate Success")
}

// db close saat sudah tidak digunakan
func (server *Server) CloseDB() {
	sqlDB, err := server.DB.DB()
	if err != nil {
		log.Println("Failed to get sql.DB:", err)
		return
	}

	err = sqlDB.Close()
	if err != nil {
		log.Println("Failed to close DB:", err)
	}

	fmt.Println("DB Close")
}

// method run server
func (server *Server) run(address string) {
	fmt.Printf("Run listening port %s", address)
	log.Fatal(http.ListenAndServe(address, server.Router))
}

// run yang akan dipanggil di app.go
func Run() {
	server := Server{}
	appConfig := AppConfig{}
	dbConfig := DBConfig{}

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error on loading get env file")
	}

	// ap
	appConfig.AppName = getEnv("APP_NAME", "GoAppStoreDev")
	appConfig.AppEnv = getEnv("APP_ENV", "Development")
	appConfig.AppPort = getEnv("APP_PORT", "8888")

	// db config
	dbConfig.DBHost = getEnv("DB_HOST", "localhost")
	dbConfig.DBPort = getEnv("DB_PORT", "go_ap_store_db")
	dbConfig.DBUser = getEnv("DB_USER", "root")
	dbConfig.DBPassword = getEnv("DB_PASSWORD", "")
	dbConfig.DBName = getEnv("DB_NAME", "go_ap_store_db")
	dbConfig.DBDriver = getEnv("DB_DRIVER", "mysql")

	server.initialize(appConfig, dbConfig)

	defer server.CloseDB()

	server.run(":" + appConfig.AppPort)
}
