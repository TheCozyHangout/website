package main

import (
	"fmt"

	"github.com/TheCozyHangout/website/backend/internal/adapters"
	"github.com/TheCozyHangout/website/backend/internal/core/application"
	"github.com/TheCozyHangout/website/backend/internal/env"
)

func main() {
	env.LoadEnv()
	adapter, err := adapters.NewMySQLAdapter()

	if err != nil {
		fmt.Println("idk new error")
	}

	application.NewApplication(adapter)
}
