package users

import (
	"strconv"

	"golang.org/x/crypto/bcrypt"
	"google.com/banking-app/database"
	"google.com/banking-app/helpers"
	"google.com/banking-app/models"
)

func Login(username, pass string) map[string]interface{} {

	message, valid := helpers.Validation([]models.Validation{{Valid: "username", Value: username}, {Valid: "password", Value: pass}})

	if valid {

		user := &models.User{}

		if err := database.Db.Where("username=?", username).Find(user).Error; err != nil {
			er := &helpers.Error{ErrorMessage: "cannot get user details", Err: err}
			er.HandleErr()
			return map[string]interface{}{"cannot get user details": er}
		}

		if passErr := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(pass)); passErr != nil {
			er := &helpers.Error{ErrorMessage: "incorrect password", Err: passErr}
			er.HandleErr()
			return map[string]interface{}{"incorrect password": er}
		}

		accounts := []models.Account{}
		if err := database.Db.Model(&models.Account{}).Select("id,name,balance").Where("user_id=?", user.ID).Find(&accounts).Error; err != nil {
			er := &helpers.Error{ErrorMessage: "cannot get account details", Err: err}
			er.HandleErr()
			return map[string]interface{}{"cannot get account details": er}
		}
		accResponse := []models.AccountsResponse{}
		for _, acc := range accounts {
			accResponse = append(accResponse, models.AccountsResponse{ID: int(acc.ID), Name: acc.Name, Balance: acc.Balance})
		}

		userResponse := &models.UserResponse{
			ID:       int(user.ID),
			UserName: user.Username,
			Email:    user.Email,
			Accounts: accResponse,
		}

		response := make(map[string]interface{})
		response["details"] = userResponse
		response["token"] = helpers.GetJwtToken(strconv.Itoa(int(user.ID)))

		return response

	} else {
		return map[string]interface{}{"message": message}
	}

}

func Register(username, email, pass string) map[string]interface{} {

	if message, valid := helpers.Validation([]models.Validation{{Valid: "username", Value: username}, {Valid: "email", Value: email}, {Valid: "password", Value: pass}}); valid {

		// generate mock passowrd and insert into table

		password := helpers.HashandSalt([]byte(pass))
		user := &models.User{Username: username, Email: email, Password: password}
		database.Db.Create(&user)

		account := &models.Account{Type: "Savings Account", Name: username, Balance: 0, UserId: user.ID}
		database.Db.Create(&account)

		resp := []*models.AccountsResponse{}
		resp = append(resp, &models.AccountsResponse{ID: int(user.ID), Name: user.Username, Balance: 0})

		return map[string]interface{}{"details": resp}
	} else {
		return map[string]interface{}{"message": message}
	}
}

func GetUser(id, jwt string) map[string]interface{} {
	user := new(models.User)

	if err := database.Db.Where("id=?", id).Find(user).Error; err != nil {
		er := &helpers.Error{ErrorMessage: "cannot get user details", Err: err}
		er.HandleErr()
		return map[string]interface{}{"cannot get user details": er}
	}

	accounts := []models.AccountsResponse{}
	if err := database.Db.Model(&models.Account{}).Select("id,name,balance").Where("user_id=?", user.ID).Find(&accounts).Error; err != nil {
		er := &helpers.Error{ErrorMessage: "cannot get account details", Err: err}
		er.HandleErr()
		return map[string]interface{}{"cannot get account details": er}
	}

	userResponse := &models.UserResponse{
		ID:       int(user.ID),
		UserName: user.Username,
		Email:    user.Email,
		Accounts: accounts,
	}

	resp := make(map[string]interface{})
	resp["details"] = userResponse

	return resp
}
