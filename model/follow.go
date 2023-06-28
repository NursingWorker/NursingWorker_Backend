package model

type Follow struct {
	ID     uint `gorm:"primarykey"`
	IdolID string
	FanID  string
}

func FlCt(IdolID string, FanID string) error {
	var follow Follow
	follow.FanID = FanID
	follow.IdolID = IdolID
	err := DB.Create(&follow).Error
	return err
}

func FlDt(IdolID string, FanID string) error {
	err := DB.Delete(&Follow{}, "idol_id= ? AND fan_id =?", IdolID, FanID).Error
	return err
}

func Idols(fanID string) ([]User, error) {
	var IdolIDs []string
	if err := DB.Model(&Follow{}).Select("idol_id").Where("fan_id = ?", fanID).Find(&IdolIDs).Error; err != nil {
		return nil, err
	}
	var users []User
	err := DB.Where("id in (?)", IdolIDs).Find(&users).Error
	return users, err
}

func Fans(IdolID string) ([]User, error) {
	var FanIDs []string
	if err := DB.Model(&Follow{}).Select("fan_id").Where("idol_id = ?", IdolID).Find(&FanIDs).Error; err != nil {
		return nil, err
	}
	var users []User
	err := DB.Where("id in (?)", FanIDs).Find(&users).Error
	return users, err
}
