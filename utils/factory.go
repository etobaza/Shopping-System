package utils

import "shopping-system/models"

func ManufactureUser(username string, password string, email string, firstName string, lastName string, address string, phone string, usertype string) models.User {
	var user models.User
	user.Username = username
	user.Password = password
	user.Email = email
	user.FirstName = firstName
	user.LastName = lastName
	user.Address = address
	user.Phone = phone
	user.UserType = usertype
	return user
}
