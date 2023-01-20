package main

import (
	"context"
	"max-inventory/internal/database"
	"max-inventory/internal/settings"

	"go.uber.org/fx"
)

func main() {
	app := fx.New(
		fx.Provide(
			context.Background,
			settings.New,
			database.New,
		),

		fx.Invoke(),
	)
	app.Run()

}
