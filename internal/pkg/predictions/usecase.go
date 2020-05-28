package predictions

type IUsecase interface {
	Predict(tags *[]string) (*[]string, error)
}
