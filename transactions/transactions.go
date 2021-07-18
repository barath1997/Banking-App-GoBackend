package transactions

import (
	"fmt"

	"google.com/banking-app/database"
	"google.com/banking-app/helpers"
	"google.com/banking-app/models"
)

func CreateTransaction(from, to uint, amount int) {

	transaction := &models.Transaction{From: from, To: to, Amount: amount}
	if err := database.Db.Create(&transaction).Error; err != nil {
		er := &helpers.Error{ErrorMessage: "unable to add transaction details ", Err: err}
		er.HandleErr()
	}
}

func GetTranactionByAccount(id uint) []models.Transaction {

	transactions := []models.Transaction{}

	if err := database.Db.Model(&models.Transaction{}).Where("Transactions.From=?", id).Or("Transactions.To=?", id).Select("id,Transactions.from,Transactions.to,amount").Scan(&transactions).Error; err != nil {
		er := &helpers.Error{ErrorMessage: "unable to get transaction records ", Err: err}
		er.HandleErr()
	}

	fmt.Println(" transactions ", transactions)
	return transactions
}

func GetMyTransactions(id, jwt string) map[string]interface{} {

	if isValid := helpers.ValidateToken(id, jwt); isValid {
		accounts := []models.Account{}
		if err := database.Db.Model(&models.Account{}).Where("user_id=?", id).Scan(&accounts).Error; err != nil {
			er := &helpers.Error{ErrorMessage: "unable to get user's transaction records ", Err: err}
			er.HandleErr()
		}

		transactions := []models.Transaction{}

		for _, acc := range accounts {
			fmt.Println("id ", acc.ID)
			accTransaction := GetTranactionByAccount(acc.ID)
			transactions = append(transactions, accTransaction...)
		}

		response := make(map[string]interface{})
		response["data"] = transactions
		return response

	} else {
		return map[string]interface{}{"message": "not a valid token "}
	}
}
