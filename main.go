package main

import (
	"fmt"

	"github.com/CrymsonShadows/blog-aggregator/internal/config"
)

func main() {
	cfg := config.Read()
	err := cfg.SetUser("Crymson")
	if err != nil {
		fmt.Println(err.Error())
	}
	cfg = config.Read()
	fmt.Printf("db_url: %s\ncurrent_user_name: %s\n", cfg.DbURL, cfg.CurrentUserName)
}
