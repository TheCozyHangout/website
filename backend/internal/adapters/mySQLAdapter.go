package adapters

import (
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type MySQLAdapter struct {
	client *sqlx.DB
}

func NewMySQLAdapter() (*MySQLAdapter, error) {
	str := fmt.Sprintf("%s:%s@/cozywebsite", os.Getenv("MYSQL_USER"), os.Getenv("MYSQL_PASSWORD"))

	client, err := sqlx.Connect("mysql", str)

	if err != nil {
		return nil, err
	}

	pingErr := client.Ping()

	if pingErr != nil {
		return nil, err
	}

	return &MySQLAdapter{client: client}, nil
}
