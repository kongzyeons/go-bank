package main

import (
	user_repo "github.com/kongzyeons/go-bank/internal/repositories/user"
	usergreeting_repo "github.com/kongzyeons/go-bank/internal/repositories/user-greeting"
	user_task "github.com/kongzyeons/go-bank/internal/services/task"
	"github.com/kongzyeons/go-bank/pkg/postgresql"
)

func main() {
	db := postgresql.InitPostgresql()
	defer db.Close()
	userRepo := user_repo.NewUserRepo(db)
	userGreetingRepo := usergreeting_repo.NewUserGreetingRepo(db)
	userTask := user_task.NewUserTask(
		db, userRepo, userGreetingRepo,
	)
	userTask.CreateTable()
}
