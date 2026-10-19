package di

import (
	"game-service/internal/app"
	"game-service/internal/domain"
	"game-service/internal/infrastructure"
	"game-service/internal/web"

	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(infrastructure.NewGameStorage),
	fx.Provide(infrastructure.NewRepository),

	fx.Provide(domain.NewGameLogic),

	fx.Provide(app.NewService),

	fx.Provide(web.NewHandler),
)
