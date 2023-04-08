package middlewares

import "shopping-system/models"

func ManufactureUser(username string, password string, email string, firstName string, lastName string, address string, phone string, usertype string) models.User {
	var user models.User
	AssignFieldsUser(&user, username, password, email, firstName, lastName, address, phone, usertype)
	return user
}

func AssignFieldsUser(user *models.User, Username string, Password string, Email string, FirstName string, LastName string, Address string, Phone string, UserType string) {
	user.Username = Username
	user.Password = Password
	user.Email = Email
	user.FirstName = FirstName
	user.LastName = LastName
	user.Address = Address
	user.Phone = Phone
	user.UserType = UserType
}
