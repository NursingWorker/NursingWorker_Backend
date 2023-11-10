package model

type Comment struct {
	ID      uint   `gorm:"primarykey"`
	PostID  string `form:"postID" json:"postID"`
	UserID  string
	Content string `form:"content" json:"content"`
}

func Cmt(postID string) ([]Comment, error) {
	var cmts []Comment
	if err := DB.Model(&Comment{}).Where("post_id = ?", postID).Find(&cmts).Error; err != nil {
		return nil, err
	}
	return cmts, nil
}
func CmtCt(cmt Comment) error {
	if err := DB.Create(&cmt).Error; err != nil {
		return err
	}
	return nil
}
func IsMyComment(userID string, cmtID string) bool {
	var tmp Comment
	if err := DB.Model(&Comment{}).Select("id").Where("id = ? AND user_id= ?", cmtID, userID).Take(&tmp).Error; err != nil {
		return false
	}
	return true
}
func CmtDt(cmtID string) {
	DB.Delete(&Comment{}, "id = ?", cmtID)
}
