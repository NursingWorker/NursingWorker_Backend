package model

type Post struct {
	ID      uint   `gorm:"primarykey;column:id"`
	OpenID  string `json:"-"`
	Title   string `form:"title" gorm:"column:title"`
	Content string `gorm:"column:content" form:"content"`
}

type Videos struct {
	ID     uint `gorm:"primarykey"`
	PostID string
	Path   string
}

type Images struct {
	ID     uint `gorm:"primarykey"`
	PostID string
	Path   string
}
type PostFl struct {
	ID      uint   `gorm:"primarykey;column:id"`
	OpenID  string `json:"-"`
	Title   string `form:"title" gorm:"column:title"`
	Content string `gorm:"column:content" form:"content"`
	Images  []string
	Videos  []string
}

func Create(post Post) (Post, error) {
	if err := DB.Create(&post).Error; err != nil {
		return Post{}, err
	}
	return post, nil
}

func CreateImage(postID string, path string) error {
	var image Images
	image.Path = path
	image.PostID = postID
	if err := DB.Create(&image).Error; err != nil {
		return err
	}
	return nil
}

func CreateVideo(postID string, path string) error {
	var video Videos
	video.PostID = postID
	video.Path = path
	if err := DB.Create(&video).Error; err != nil {
		return err
	}
	return nil
}

func IsMyPost(postID string, openID string) bool {
	var tmp Post
	if err := DB.Model(&Post{}).Select("id").Where("id = ? AND open_id= ?", postID, openID).Take(&tmp).Error; err != nil {
		return false
	}
	return true
}

func PtDt(postID string) {
	DB.Delete(&Post{}, "id = ?", postID)
}

func HisPost(openID string) ([]PostFl, error) {
	var posts []PostFl
	err := DB.Model(&Post{}).Where("open_id = ?", openID).Find(&posts).Error
	if err != nil {
		return nil, err
	}
	for i, _ := range posts {
		postID := posts[i].ID
		if err = DB.Model(&Images{}).Select("path").Where(" post_id = ? ", postID).Find(&posts[i].Images).Error; err != nil {
			return nil, err
		}
		if err = DB.Model(&Videos{}).Select("path").Where(" post_id = ? ", postID).Find(&posts[i].Videos).Error; err != nil {
			return nil, err
		}
	}
	return posts, err
}

func HisLike(userID string) ([]Post, error) {
	var postIDs []string
	DB.Model(&Like{}).Select("post_id").Where("user_id= ?", userID).Find(&postIDs)
	var posts []Post
	err := DB.Model(&Post{}).Where("id in (?)", postIDs).Find(&posts).Error
	return posts, err
}

func HisColl(userID string) ([]Post, error) {
	var postIDs []string
	DB.Model(&Collection{}).Select("post_id").Where("user_id= ?", userID).Find(&postIDs)
	var posts []Post
	err := DB.Model(&Post{}).Where("id in (?)", postIDs).Find(&posts).Error
	return posts, err
}

func ReCmt(old []string, number string) ([]PostFl, error) {
	var posts []PostFl
	result := DB.Raw("SELECT * FROM posts ORDER BY RAND() LIMIT ?", number).Scan(&posts)
	if result.Error != nil {
		return nil, result.Error
	}
	for i, _ := range posts {
		postID := posts[i].ID
		if err := DB.Model(&Images{}).Select("path").Where(" post_id = ? ", postID).Find(&posts[i].Images).Error; err != nil {
			return nil, err
		}
		if err := DB.Model(&Videos{}).Select("path").Where(" post_id = ? ", postID).Find(&posts[i].Videos).Error; err != nil {
			return nil, err
		}
	}
	return posts, nil
}
