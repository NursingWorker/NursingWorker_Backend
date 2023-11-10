package model

import "strconv"

func CarerSc(point string) ([]User, error) {
	var users []User
	err := DB.Model(&User{}).Where("identity = carer AND nickname like %?%", point).Find(&users).Error
	return users, err
}

func Type(tp string, number string) ([]User, error) {
	var users []User
	n, _ := strconv.Atoi(number)
	err := DB.Model(&User{}).Where("identity = ?", tp).Limit(n).Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

func FindViewByCarerID(carerID string) ([]View, error) {
	var views []View
	err := DB.Model(&View{}).Where("carer_id = ?", carerID).Find(views).Error
	if err != nil {
		return nil, err
	}
	return views, nil
}

func IsHire(carerID string, userID string) bool {
	err := DB.Model(&Hire{}).Where("carer_id = ? AND user_id = ?", carerID, userID).Take(&Hire{}).Error
	if err != nil {
		return false
	}
	return true
}
