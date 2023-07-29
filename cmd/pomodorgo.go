package main

import (
	"fmt"

	"github.com/JoelSaleem/pomodorgo/internal/db"
	"github.com/kelseyhightower/envconfig"
)

type Specification struct {
	DBPath string `split_words:"true"`
}

func main() {
	var spec Specification
	envconfig.MustProcess("pomodorgo", &spec)

	if spec.DBPath == "" {
		panic("DBPath not set")
	}

	fmt.Println("\n\n x \n\n", spec.DBPath)

	repo := db.NewRepository(spec.DBPath)
}

// alt-screen toggle (full screen app)
// list default
// progress-animated (progress bar)
// (may need result -- using the result of selection)
// see timers -- how do they do progress updates -- most likely with goroutines
// how do I make it durable?
