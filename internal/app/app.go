package app

import (
	"fmt"
	"log/slog"

	"github.com/koha90/podkrepizza-api-v1/config"
	"github.com/koha90/podkrepizza-api-v1/internal/storage/sqlite"
)

// Run - start aplication with config.
func Run(cfg config.Config, log *slog.Logger) error {
	const op = "app.Run"

	if _, err := fmt.Println("start application:", cfg.App.Name, "version:", cfg.App.Name); err != nil {
		return fmt.Errorf("%s: error: %w", op, err)
	}

	storage, err := sqlite.New(cfg.Storage.Path)
	if err != nil {
		return fmt.Errorf("%s: failed to initialize storage: %w", op, err)
	}

	if err = storage.AddPizzas(); err != nil {
		log.Error("error to add pizza ", err)
	}
	fmt.Println(storage.GetAllPizzas())
	return nil
}
