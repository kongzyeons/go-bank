package user_task

import (
	"testing"

	user_repo "github.com/kongzyeons/go-bank/internal/repositories/user"
	usergreeting_repo "github.com/kongzyeons/go-bank/internal/repositories/user-greeting"
	"github.com/kongzyeons/go-bank/pkg/postgresql"
)

func TestCreateTable(t *testing.T) {
	db := postgresql.InitPostgresql()
	defer db.Close()
	userRepo := user_repo.NewUserRepo(db)
	userGreetingRepo := usergreeting_repo.NewUserGreetingRepo(db)
	userTask := NewUserTask(
		db, userRepo, userGreetingRepo,
	)
	userTask.CreateTable()
}
