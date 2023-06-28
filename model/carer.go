package model


func CarerSc(point string) ([]User, error) {
	var users []User
	err := DB.Model(&User{}).Where("identity = carer AND nickname like %?%", point).Find(&users).Error
	return users, err
}

//func Type(tp string , number string)([]User,error){
//	var users []User
//	result := DB.Raw("SELECT * FROM users WHERE  ORDER BY RAND() LIMIT ?", number).Scan(&users)
//	if result.Error != nil {
//		return nil, result.Error
//	}
//	return posts, nil
//}