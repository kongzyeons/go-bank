package message

import (
	"github.com/IBM/sarama"
	"github.com/jmoiron/sqlx"
	"github.com/kongzyeons/go-bank/internal/queues"
	accountbalance_repo "github.com/kongzyeons/go-bank/internal/repositories/account-balance"
	transaction_repo "github.com/kongzyeons/go-bank/internal/repositories/transaction"
	account_queue "github.com/kongzyeons/go-bank/internal/services/queue/account"
	"github.com/kongzyeons/go-bank/pkg/line"
)

func InitMessage(
	db *sqlx.DB,
) sarama.ConsumerGroupHandler {
	accountBalanceRepo := accountbalance_repo.NewaccountBalanceRepo(db)
	transectionRepo := transaction_repo.NewTransactionRepo(db)
	lineNoti := line.NewLineAPI()
	accountEventHandler := account_queue.NewAccountEventHandler(
		db,
		accountBalanceRepo, transectionRepo,
		lineNoti,
	)
	consumerHandler := queues.NewConsumerHandler(accountEventHandler)

	return consumerHandler
}
