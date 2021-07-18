package server

import (
	"google.com/banking-app/transactions"
	"google.com/banking-app/useraccounts"
	"google.com/banking-app/users"
)

type Server struct{}

func (s *Server) LoginService(username, pass string) map[string]interface{} {

	return users.Login(username, pass)

}

func (s *Server) RegisterService(username, email, pass string) map[string]interface{} {

	return users.Register(username, email, pass)

}

func (s *Server) GetUserService(id, jwt string) map[string]interface{} {

	return users.GetUser(id, jwt)

}

func (s *Server) TransactService(userId, fromId, toId uint, amount int, jwt string) map[string]interface{} {

	return useraccounts.Transcation(userId, fromId, toId, amount, jwt)

}

func (s *Server) GetMyTransactionsService(userId, jwt string) map[string]interface{} {

	return transactions.GetMyTransactions(userId, jwt)

}
