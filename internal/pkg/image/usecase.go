package image

type IUsecase interface {
	Analyze(pinId uint, userId uint, image string)
}
