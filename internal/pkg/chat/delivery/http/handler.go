package http

import (
	"2020_1_Color_noise/internal/pkg/chat"
	"go.uber.org/zap"
)

type Handler struct {
	chatUsecase  chat.IUsecase
	logger         *zap.SugaredLogger
}

