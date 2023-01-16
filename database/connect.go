package database

import (
	"fmt"
	"go-issues-api/config"
	"go-issues-api/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConnectDatabase(mode string) {
	DB, err = gorm.Open(postgres.Open(dsn(mode)), &gorm.Config{})

	if err != nil {
		fmt.Println(err.Error())
		panic("failed to connect database")
	}

	fmt.Println("connect to database")

	err := DB.AutoMigrate(&model.Issue{}) // TODO: refactor later

	if err != nil {
		panic("failed to run data migration")
	}
}

// TODO: refactor later
func SeedIssues() {
	var seedIssues = []model.Issue{
		{
			Title:       "issue 1",
			Description: "This is issue 1",
		},
		{
			Title:       "issue 2",
			Description: "This is issue 2",
		},
	}
	DB.Create(&seedIssues)
}

func DisconnectDatabase() {
	db, _ := DB.DB()
	defer db.Close()
	DB.Migrator().DropTable(&model.Issue{})
}

func dsn(mode string) string {
	conn := config.PostgresConn(mode)

	return fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		conn.Host,
		conn.User,
		conn.Password,
		conn.Dbname,
		conn.Port,
		conn.Ssl,
		conn.Timezone,
	)
}
