package model

type Like struct {
	ID     uint `gorm:"primarykey"`
	UserID string
	PostID string
}

func LikeCt(userID string, postID string) error {
	var like Like
	like.PostID = postID
	like.UserID = userID
	err := DB.Create(&like).Error
	return err
}

func LikeDt(userID string, postID string) error {
	err := DB.Delete(&Like{}, "user_id=? AND post_id = ?", userID, postID).Error
	return err
}

func IsMyLike(userID string, postID string) bool {
	var like Like
	err := DB.Model(&Like{}).Where("user_id=? AND post_id = ?", userID, postID).Take(&like).Error
	if err != nil {
		return false
	}
	return true
}
