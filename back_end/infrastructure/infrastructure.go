package infrastructure

import (
	"flag"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

const (
	APPPORT    = ""
	DBHOST     = ""
	DBPORT     = ""
	DBUSER     = ""
	DBPASSWORD = ""
	DBNAME     = ""

	HTTPSWAGGER = ""
	ROOTPATH    = ""
)

var (
	appPort    string
	dbHost     string
	dbPort     string
	dbUser     string
	dbPassword string
	dbName     string

	httpSwagger string
	rootPath    string

	InfoLog *log.Logger
	ErrLog  *log.Logger

	db *gorm.DB
)

func getStringEnvParameter(envParam string, defaultValue string) string {
	if value, ok := os.LookupEnv(envParam); ok {
		return value
	}
	return defaultValue
}

func goDotEnvVariable(key string) string {
	err := godotenv.Load(".env") 
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}

func loadEnvParameters() {
	root, _ := os.Getwd()
	appPort = getStringEnvParameter(APPPORT, goDotEnvVariable("APPORT"))
	dbHost = getStringEnvParameter(DBHOST, goDotEnvVariable("DBHOST"))
	dbUser = getStringEnvParameter(DBUSER, goDotEnvVariable("DBUSER"))
	dbPassword = getStringEnvParameter(DBPASSWORD, goDotEnvVariable("DBPASSWORD"))
	dbName = getStringEnvParameter(DBNAME, goDotEnvVariable("DBNAME"))

	httpSwagger = getStringEnvParameter(HTTPSWAGGER, goDotEnvVariable("HTTPSWAGGER"))
	rootPath = getStringEnvParameter(ROOTPATH, root)
}

func init() {
	InfoLog = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Llongfile)
	ErrLog = log.New(os.Stdout, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)

	var initDB bool
	flag.BoolVar(&initDB, "db", false, "allow recreate model database in postgres")
	flag.Parse()

	loadEnvParameters()

}

// GetDB get database instance
func GetDB() *gorm.DB {
	return db
}

// GetDBName get database name
func GetDBName() string {
	return dbName
}

// GetHTTPSwagger export link swagger
func GetHTTPSwagger() string {
	return httpSwagger
}

// GetAppPort export app port
func GetAppPort() string {
	return appPort
}

// GetRootPath export root path system
func GetRootPath() string {
	return rootPath
}