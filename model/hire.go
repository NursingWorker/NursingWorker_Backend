package model

type Hire struct {
	ID      uint `gorm:"primarykey"`
	CarerID string
	UserID  string
}

func HireCt(CarerID string, UserID string) error {
	var hire Hire
	hire.CarerID = CarerID
	hire.UserID = UserID
	err := DB.Create(&hire).Error
	return err
}

func HireDt(CarerID string, UserID string) error {
	err := DB.Delete(&Follow{}, "carer_id= ? AND user_id =?", CarerID, UserID).Error
	return err
}
