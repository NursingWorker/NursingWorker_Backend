package model

type Reply struct {
	ID      uint   `gorm:"primarykey"`
	CmtID   string `form:"commentID"`
	ObtID   string `form:"objectID"`
	UserID  string
	Content string `form:"content"`
}

func Rep(commentID string) ([]Reply, error) {
	var reps []Reply
	if err := DB.Model(&Reply{}).Where("cmt_id = ?", commentID).Find(&reps).Error; err != nil {
		return nil, err
	}
	authID := UserIDCmt(commentID)
	for i, _ := range reps {
		if reps[i].ObtID == authID {
			reps[i].ObtID = ""
			continue
		}
		reps[i].ObtID = NickName(reps[i].ObtID)
	}
	return reps, nil
}
func RepCt(rep Reply) error {
	err := DB.Create(&rep).Error
	return err
}

func IsMyReply(userID string, repID string) bool {
	var tmp Reply
	if err := DB.Model(&Reply{}).Select("id").Where("id = ? AND user_id= ?", repID, userID).Take(&tmp).Error; err != nil {
		return false
	}
	return true
}

func RepDt(repID string) {
	DB.Delete(&Reply{}, "id = ?", repID)
}
