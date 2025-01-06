package main

import (
	"fmt"

	"github.com/CrymsonShadows/blog-aggregator/internal/config"
)

func main() {
	cfg := config.Read()
	cfg.SetUser("Crymson")
	cfg = config.Read()
	fmt.Printf("db_url: %s\ncurrent_user_name: %s\n", cfg.DbURL, cfg.CurrentUserName)
}
