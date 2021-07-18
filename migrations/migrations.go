package migrations

import (
	"google.com/banking-app/database"
	"google.com/banking-app/helpers"
	"google.com/banking-app/models"
)

// create mock data and insert into tables.
func CreateAccounts() {
	// connect to db

	// generate mock data
	users := &[2]models.User{
		{Username: "steve smith", Email: "steve.smith@gmail.com"},
		{Username: "virat kohli", Email: "virat.kohli@gmail.com"},
	}

	// generate mock passowrd and insert into table
	for _, u := range users {
		pass := helpers.HashandSalt([]byte(u.Username))
		user := &models.User{Username: u.Username, Email: u.Email, Password: pass}
		database.Db.Create(&user)

		account := &models.Account{Type: "Savings Account", Name: user.Username, Balance: uint(4000), UserId: user.ID}
		database.Db.Create(&account)
	}

}

func Migrate() {

	database.Db.AutoMigrate(&models.User{}, &models.Account{})

	CreateAccounts()
}

func MigrateTr() {

	database.Db.AutoMigrate(&models.Transaction{})
}
