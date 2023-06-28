package service

import (
	"fmt"
	"nursing_work/model"
)

func CreatePost(post model.Post, images []string, videos []string) error {
	post, err := model.Create(post)
	if err != nil {
		return err
	}
	postID := fmt.Sprint(post.ID)
	//images
	for _, image := range images {
		model.CreateImage(postID, image)
	}
	//videos
	for _, video := range videos {
		model.CreateVideo(postID, video)
	}
	return nil
}
