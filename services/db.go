package services

import (
	"fmt"
	"database/sql"
	"github.com/joho/godotenv"
	"log"
	"os"

	_ "github.com/lib/pq" // Driver de conexión con Postgres
)

type DbData struct {
	Host   			string
	Port   			string
	DbName 			string
	RolName     string
	RolPassword string
}

var Db *sql.DB

func LoadEnv() (DbData, error) {
	var err error

	if err = godotenv.Load(".env"); err != nil {
		log.Fatalf("Error al cargar el archivo .env")
		return DbData{}, err
	}

	return DbData{
		Host:    		 os.Getenv("HOST"),
		Port:     	 os.Getenv("PORT"),
		DbName:   	 os.Getenv("DB_NAME"),
		RolName:  	 os.Getenv("ROL_NAME"),
		RolPassword: os.Getenv("ROL_PASSWORD"),
	}, nil
}

// Abre la conexión con la base de datos
func EstablishDbConnection() (error) {

	dbData, err := LoadEnv()

	if err != nil {
		log.Fatal(err)
	}

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", dbData.Host, dbData.Port, dbData.RolName, dbData.RolPassword, dbData.DbName)

	dbConn, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}

	Db = dbConn
	fmt.Println("Conexión exitosa a la base de datos:", dbConn)

	if err = Db.Ping(); err != nil {
		Db.Close()
		return err
	}

	return nil
}
