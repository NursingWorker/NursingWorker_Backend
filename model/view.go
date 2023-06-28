package model

type View struct {
	ID      uint   `gorm:"primarykey"`
	CarerID string `form:"carerID"`
	UserID  string
	Content string `form:"content"`
}

func ViewCt(view View) error {
	err := DB.Create(&view).Error
	return err
}

func ViewDt(viewID string, userID string) error {
	err := DB.Delete(&View{}, "id = ? AND user_id = ?", viewID, userID).Error
	return err
}
