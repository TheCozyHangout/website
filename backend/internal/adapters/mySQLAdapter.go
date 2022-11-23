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
		fmt.Println("Failed to connect to MySQL")
		return nil, err
	}

	pingErr := client.Ping()

	if pingErr != nil {
		fmt.Println("cancer")
	}

	return &MySQLAdapter{client: client}, nil
}
