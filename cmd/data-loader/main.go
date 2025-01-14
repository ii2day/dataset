package main

import (
	"github.com/BaizeAI/dataset/internal/cmd/dataloader"
	"github.com/BaizeAI/dataset/pkg/log"
)

func main() {
	log.SetDebug()

	cmd := dataloader.NewCommand()
	err := cmd.Execute()
	if err != nil {
		panic(err)
	}
}
