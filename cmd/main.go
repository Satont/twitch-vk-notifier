package main

import (
	i18nstore "github.com/satont/twitch-notifier/internal/i18n/store"
	"github.com/satont/twitch-notifier/internal/pgx"
	repositories "github.com/satont/twitch-notifier/internal/repository/fx"
	"github.com/satont/twitch-notifier/pkg/logger"
	"go.uber.org/fx"
)

func main() {
	fx.New(
		// fx.NopLogger,
		fx.Provide(
			fx.Annotate(
				logger.NewFx(),
				fx.As(new(logger.Logger)),
			),
			fx.Annotate(
				i18nstore.New,
				fx.As(new(i18nstore.I18nStore)),
			),
			pgx.New,
		),
		repositories.Module,
		fx.Invoke(pgx.New),
	).Run()
}
