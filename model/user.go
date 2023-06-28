package model

type User struct {
	ID         uint   `gorm:"primarykey"`
	OpenID     string `json:"-"`
	NickName   string `form:"nickname"`
	Avatar     string `form:"avatar"`
	Intro      string `form:"intro"`
	Age        string `form:"age"`
	Gender     string `form:"gender"`
	Stature	   string `form:"stature"`
	Address    string `form:"address"`
	Experience string `form:"experience"`
	Identity   string `form:"identity"`
}

type Token struct {
	ID     uint `gorm:"primarykey"`
	OpenID string
	Token  string
}

func CreateUser(user User, token Token) error {
	if err := DB.Create(&user).Error; err != nil {
		return err
	}
	if err := DB.Create(&token).Error; err != nil {
		return err
	}
	return nil
}

func Exist(openID string) bool {
	var user User
	if err := DB.Model(&User{}).Select("id").Where("open_id = ?", openID).Take(&user).Error; err != nil {
		return false
	}
	return true
}

func TokenUpdate(token string, openID string) error {
	if err := DB.Model(&Token{}).Where("open_id = ?", openID).Update("token", token).Error; err != nil {
		return err
	}
	return nil
}

func TokenLatest(openID string) string {
	var token string
	DB.Model(&Token{}).Select("token").Where("open_id = ?", openID).Find(&token)
	return token
}

func Base(userID string) (User, error) {
	var user User
	err := DB.Model(&User{}).Where("id=?", userID).Take(&user).Error
	return user, err
}

func Avatar(openID string, path string) error {
	err := DB.Model(&User{}).Where("open_id = ?", openID).Update("avatar", path).Error
	return err
}

func Identity(openID string, identity string) error {
	err := DB.Model(&User{}).Where("open_id = ?", openID).Update("identity", identity).Error
	return err
}

func InfoUpt(userID string, user User) error {
	err := DB.Model(&User{}).Where("id = ?", userID).Updates(user).Error
	return err
}
