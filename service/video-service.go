package service

import "go-gin/entity"

type VideoService interface {
	Save(entity.Video) entity.Video
	FindAll() []entity.Video
}

type videoServ struct {
	videos []entity.Video
}

func New() VideoService {
	return &videoServ{}
}

func (service *videoServ) Save(video entity.Video) entity.Video {
	service.videos = append(service.videos, video)
	return video
}
func (service *videoServ) FindAll() []entity.Video {
	return service.videos
}
