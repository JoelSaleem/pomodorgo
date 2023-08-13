package main

import (
	"fmt"
	"os"

	"github.com/JoelSaleem/pomodorgo/internal/db"
	"github.com/JoelSaleem/pomodorgo/internal/tabs"
	"github.com/JoelSaleem/pomodorgo/internal/tea_model"
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

	repo := db.NewRepository(spec.DBPath)
	if _, err := tea_model.NewProgram(repo, tabs.ConstructTabs()).Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}

// alt-screen toggle (full screen app)
// list default
// progress-animated (progress bar)
// (may need result -- using the result of selection)
// see timers -- how do they do progress updates -- most likely with goroutines
// how do I make it durable?
