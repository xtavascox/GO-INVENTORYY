package main

import (
	"context"
	"github.com/go-inventory/database"
	"github.com/go-inventory/internal/repository"
	"github.com/go-inventory/internal/services"
	"github.com/go-inventory/settings"
	"go.uber.org/fx"
)

func main() {
	app := fx.New(
		fx.Provide(
			context.Background,
			settings.New,
			database.New,
			repository.New,
			services.New,
		),
		fx.Invoke(
			func(ctx context.Context, serv services.Service) {

				//err := serv.RegisterUser(ctx, "my@email2.com", "1234", "gustavo 2")
				//if err != nil {
				//	panic(err)
				//}

				u, err := serv.LoginUser(ctx, "my@email2.com", "1234")
				if err != nil {
					panic(err)
				}

				println(u.Name, u.Email, u.ID)

			},
		),
	)
	app.Run()
}
