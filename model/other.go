package model

func UserID(openID string) string {
	var userID string
	DB.Model(&User{}).Select("id").Where("open_id = ?", openID).Find(&userID)
	return userID
}

func OpenID(userID string) string {
	var openID string
	DB.Model(&User{}).Select("open_id").Where("id = ?", userID).Find(&openID)
	return openID
}

func NickName(ID string) string {
	var nickname string
	DB.Model(&User{}).Select("nick_name").Where("id = ?", ID).Find(&nickname)
	return nickname
}

func UserIDCmt(commentID string) string {
	var id string
	DB.Model(&Comment{}).Select("user_id").Where("id = ?", commentID).Find(&id)
	return id
}
