package model

type Collection struct {
	ID     uint `gorm:"primarykey"`
	UserID string
	PostID string
}

func CltCt(userID string, postID string) error {
	var coll Collection
	coll.PostID = postID
	coll.UserID = userID
	err := DB.Create(&coll).Error
	return err
}

func CltDt(userID string, postID string) error {
	err := DB.Delete(&Collection{}, "user_id=? AND post_id = ?", userID, postID).Error
	return err
}

func IsMyClt(userID string, postID string) bool {
	var coll Collection
	err := DB.Model(&Collection{}).Where("user_id=? AND post_id = ?", userID, postID).Take(&coll).Error
	if err != nil {
		return false
	}
	return true
}
