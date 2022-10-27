package Service

type UserCredentials struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

func InsertUser(username string, password string) (user UserCredentials, err error) {
	user = UserCredentials{
		Name:     username,
		Password: password,
	}
	if result := conn.Create(&user); result.Error != nil {
		return
	}
	return user, nil
}

func FindUserByUsername(username string) (user UserCredentials, err error) {
	var users UserCredentials
	tx := conn.Debug().Model(&UserCredentials{})
	if username != "" {
		tx = tx.Where("name= ?", username)
	}
	if err = tx.Find(&users).Error; err != nil {
		return
	}
	return users, err
}

func (UserCredentials) TableName() string {
	return "user"
}
