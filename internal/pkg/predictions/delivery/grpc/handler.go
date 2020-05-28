package grpc

import (
	. "2020_1_Color_noise/internal/pkg/error"
	"2020_1_Color_noise/internal/pkg/predictions"
	predictService "2020_1_Color_noise/internal/pkg/proto/predictions"
	"context"
	"log"
)

type PredictionsService struct {
	usecase predictions.IUsecase
}

func NewPredictionsService(usecase predictions.IUsecase) *PredictionsService {
	return &PredictionsService{
		usecase,
	}
}

func (ps *PredictionsService) Predict(ctx context.Context, in *predictService.Tags) (*predictService.Predictions, error) {
	if in == nil {
		return &predictService.Predictions{}, New("Bad Input for predictions")
	}

	p, err := ps.usecase.Predict(&in.Tags)
	if err != nil {
		log.Println(err)
		return &predictService.Predictions{}, Wrap(err, "Error in during predict")
	}

	return &predictService.Predictions{Predictions: *p}, nil
}