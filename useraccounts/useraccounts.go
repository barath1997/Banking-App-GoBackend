package useraccounts

import (
	"fmt"

	"google.com/banking-app/database"
	"google.com/banking-app/helpers"
	"google.com/banking-app/models"
	"google.com/banking-app/transactions"
)

func UpdateAccount(id uint, amount int) *models.AccountsResponse {

	accResponse := new(models.AccountsResponse)
	account := new(models.Account)
	if err := database.Db.Where("id=?", id).First(account).Error; err != nil {
		er := &helpers.Error{ErrorMessage: "unable to get account details", Err: err}
		er.HandleErr()
		return nil
	}

	account.Balance = uint(amount)
	if err := database.Db.Save(account).Error; err != nil {
		er := &helpers.Error{ErrorMessage: "unable to update account", Err: err}
		er.HandleErr()
		return nil
	}

	accResponse.ID = int(account.ID)
	accResponse.Name = account.Name
	accResponse.Balance = account.Balance

	return accResponse
}

func GetAccount(id uint) *models.Account {

	account := new(models.Account)

	if err := database.Db.Where("id=?", id).First(&account).Error; err != nil {
		er := &helpers.Error{ErrorMessage: "cannot get account details", Err: err}
		er.HandleErr()
		return nil
	}
	return account
}

func Transcation(userId, from, to uint, amount int, jwt string) map[string]interface{} {

	userID := fmt.Sprint(userId)

	if isValid := helpers.ValidateToken(userID, jwt); isValid {
		fromAcc := GetAccount(from)
		toAcc := GetAccount(to)

		if fromAcc == nil || toAcc == nil {
			return map[string]interface{}{"message": "Account not found"}
		} else if fromAcc.UserId != userId {
			return map[string]interface{}{"message": "not owner of this account"}
		} else if int(fromAcc.Balance) < amount {
			return map[string]interface{}{"message": "insuffecient balance"}
		}

		updatedAccount := UpdateAccount(from, int(fromAcc.Balance)-amount)
		UpdateAccount(to, int(toAcc.Balance)+amount)

		transactions.CreateTransaction(from, to, amount)

		var response = map[string]interface{}{"message": "transaction complete"}
		response["details"] = updatedAccount
		return response
	} else {
		return map[string]interface{}{"message": "invalid token"}
	}
}
