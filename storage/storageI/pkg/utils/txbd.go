package utils

import (
	"database/sql"
	"fmt"
	"os"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/DATA-DOG/go-txdb"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load("./../../.env")
	if err != nil {
		panic("can't connect to database")
	}
	fmt.Println("Here init DB===========>>", fmt.Sprintf("%s:%s@/%s", os.Getenv("USERNAME"), os.Getenv("PASSWORD"), os.Getenv("DATABASE")))
	txdb.Register("txdb", "mysql", fmt.Sprintf("%s:%s@/%s", os.Getenv("USERNAME"), os.Getenv("PASSWORD"), os.Getenv("DATABASE")))

}

func InitTxDB() (*sql.DB, error) {

	db, err := sql.Open("txdb", uuid.New().String())

	if err == nil {
		return db, db.Ping()
	}

	return db, err
}

func NewDBMock(t *testing.T) (*sql.DB, sqlmock.Sqlmock, error) {
	db, mock, err := sqlmock.New()
	if err != nil {
		return nil, nil, err
	}

	return db, mock, nil
}
