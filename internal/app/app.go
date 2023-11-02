package app

import (
	"fmt"
	"net/http"

	"github.com/koha90/podkrepizza-api-v1/config"
)

// Run - start aplication with config.
func Run(cfg *config.Config) error {
	const op = "App.Run"

	err := http.ListenAndServe(":"+cfg.HTTP.Port, nil)
	if err != nil {
		return fmt.Errorf("error %s: %w", op, err)
	}

	return nil
}
