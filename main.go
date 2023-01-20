package main

import (
	"context"
	"max-inventory/database"
	"max-inventory/settings"

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
