package model

type View struct {
	ID      uint   `gorm:"primarykey"`
	CarerID string `json:"carer_id"`
	UserID  string
	Content string `json:"content"`
}

func ViewCt(view View) error {
	err := DB.Create(&view).Error
	return err
}

func ViewDt(viewID string, userID string) error {
	err := DB.Delete(&View{}, "id = ? AND user_id = ?", viewID, userID).Error
	return err
}

func FindView(userID string) ([]View, error) {
	var views []View
	err := DB.Model(&View{}).Where("user_id = ?", userID).Find(&views).Error
	if err != nil {
		return nil, err
	}
	return views, nil
}
