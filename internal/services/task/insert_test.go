package task_svc

import (
	"errors"
	"testing"

	account_repo "github.com/kongzyeons/go-bank/internal/repositories/account"
	accountbalance_repo "github.com/kongzyeons/go-bank/internal/repositories/account-balance"
	accountdetail_repo "github.com/kongzyeons/go-bank/internal/repositories/account-detail"
	accountflag_repo "github.com/kongzyeons/go-bank/internal/repositories/account-flags"
	banner_repo "github.com/kongzyeons/go-bank/internal/repositories/banner"
	debitcard_repo "github.com/kongzyeons/go-bank/internal/repositories/debit-card"
	debitcarddesign_repo "github.com/kongzyeons/go-bank/internal/repositories/debit-card-design"
	debitcarddetails_repo "github.com/kongzyeons/go-bank/internal/repositories/debit-card-details"
	debitcardstatus_repo "github.com/kongzyeons/go-bank/internal/repositories/debit-card-status"
	transaction_repo "github.com/kongzyeons/go-bank/internal/repositories/transaction"
	user_repo "github.com/kongzyeons/go-bank/internal/repositories/user"
	usergreeting_repo "github.com/kongzyeons/go-bank/internal/repositories/user-greeting"
	"github.com/kongzyeons/go-bank/pkg/postgresql"
	"github.com/stretchr/testify/mock"
)

func TestInsertAdminData(t *testing.T) {
	testCases := []struct {
		nameTest           string
		errBeginTx         error
		errCommit          error
		erruserInsertMock  error
		erruser            error
		erruserGreeting    error
		errBanner          error
		errAccount         error
		errAccountBalance  error
		errAccountDetail   error
		errDebitCard       error
		errDebitCardStatus error
		errDebitCardDetail error
		errDebitCardDesign error
	}{
		{nameTest: "test", erruserInsertMock: errors.New("")},
		{nameTest: "test", erruserGreeting: errors.New("")},
		{nameTest: "test", errBanner: errors.New("")},
		{nameTest: "test", errAccount: errors.New("")},
		{nameTest: "test", errAccountBalance: errors.New("")},
		{nameTest: "test", errAccountDetail: errors.New("")},
		{nameTest: "test", errDebitCard: errors.New("")},
		{nameTest: "test", errDebitCardStatus: errors.New("")},
		{nameTest: "test", errDebitCardDetail: errors.New("")},
		{nameTest: "test", errDebitCardDesign: errors.New("")},
		{nameTest: "test", errCommit: errors.New("")},
		{nameTest: "test"},
	}
	for _, tc := range testCases {
		t.Run(tc.nameTest, func(t *testing.T) {
			db, mockDB, _ := postgresql.InitDatabaseMock()
			defer db.Close()
			if tc.errBeginTx == nil {
				mockDB.ExpectBegin()
			}
			if tc.errCommit == nil {
				mockDB.ExpectCommit()
			}

			userRepo := user_repo.NewUserRepoMock()
			userGreetingRepo := usergreeting_repo.NewUserGreetingRepoMock()
			bannerRepo := banner_repo.NewBannerRepoMock()
			accountRepo := account_repo.NewAccountRepoMock()
			accountBalanceRepo := accountbalance_repo.NewAccountBalanceRepoMock()
			accountDetailRepo := accountdetail_repo.NewAccountDetailRepoMock()
			accountFlagRepo := accountflag_repo.NewAccountFlagRepoMock()
			debitCardRepo := debitcard_repo.NewDebitCardRepoMock()
			debitcardstatuRepo := debitcardstatus_repo.NewDebitCardStatusRepoMock()
			debitCardSDetailRepo := debitcarddetails_repo.NewDebitCardSDetailRepoMock()
			debitCarddesignRepo := debitcarddesign_repo.NewDebitCarddesignRepoMock()
			transectionRepo := transaction_repo.NewTransactionRepoMock()

			userRepo.On("InsertMock", mock.Anything, mock.Anything).Return(
				"test", tc.erruserInsertMock,
			)
			userRepo.On("Insert", mock.Anything, mock.Anything).Return(
				"test", tc.erruser,
			)

			userGreetingRepo.On("Insert", mock.Anything, mock.Anything).Return(
				tc.erruserGreeting,
			)

			bannerRepo.On("Insert", mock.Anything, mock.Anything).Return(
				tc.errBanner,
			)

			accountRepo.On("Insert", mock.Anything, mock.Anything).Return(
				"test", tc.errAccount,
			)
			accountBalanceRepo.On("Insert", mock.Anything, mock.Anything).Return(
				tc.errAccountBalance,
			)
			accountDetailRepo.On("Insert", mock.Anything, mock.Anything).Return(
				tc.errAccountDetail,
			)

			debitCardRepo.On("Insert", mock.Anything, mock.Anything).Return(
				"test", tc.errDebitCard,
			)
			debitcardstatuRepo.On("Insert", mock.Anything, mock.Anything).Return(
				tc.errDebitCardStatus,
			)
			debitCardSDetailRepo.On("Insert", mock.Anything, mock.Anything).Return(
				tc.errDebitCardDetail,
			)
			debitCarddesignRepo.On("Insert", mock.Anything, mock.Anything).Return(
				tc.errDebitCardDesign,
			)

			task := NewTaskInsertSvc(
				db,
				userRepo, userGreetingRepo,
				bannerRepo,
				accountRepo, accountBalanceRepo, accountDetailRepo, accountFlagRepo,
				debitCardRepo, debitcardstatuRepo, debitCardSDetailRepo, debitCarddesignRepo,
				transectionRepo,
			)

			task.InsertAdminData()

		})
	}
}

func TestInsertSimpleData(t *testing.T) {
	testCases := []struct {
		nameTest           string
		errBeginTx         error
		errCommit          error
		erruserInsertMock  error
		erruser            error
		erruserGreeting    error
		errBanner          error
		errAccount         error
		errAccountBalance  error
		errAccountDetail   error
		errDebitCard       error
		errDebitCardStatus error
		errDebitCardDetail error
		errDebitCardDesign error
	}{
		{nameTest: "test", erruserInsertMock: errors.New("")},
		{nameTest: "test", erruserGreeting: errors.New("")},
		{nameTest: "test", errBanner: errors.New("")},
		{nameTest: "test", errAccount: errors.New("")},
		{nameTest: "test", errAccountBalance: errors.New("")},
		{nameTest: "test", errAccountDetail: errors.New("")},
		{nameTest: "test", errDebitCard: errors.New("")},
		{nameTest: "test", errDebitCardStatus: errors.New("")},
		{nameTest: "test", errDebitCardDetail: errors.New("")},
		{nameTest: "test", errDebitCardDesign: errors.New("")},
		{nameTest: "test", errCommit: errors.New("")},
		{nameTest: "test"},
	}
	for _, tc := range testCases {
		t.Run(tc.nameTest, func(t *testing.T) {
			db, mockDB, _ := postgresql.InitDatabaseMock()
			defer db.Close()
			if tc.errBeginTx == nil {
				mockDB.ExpectBegin()
			}
			if tc.errCommit == nil {
				mockDB.ExpectCommit()
			}

			userRepo := user_repo.NewUserRepoMock()
			userGreetingRepo := usergreeting_repo.NewUserGreetingRepoMock()
			bannerRepo := banner_repo.NewBannerRepoMock()
			accountRepo := account_repo.NewAccountRepoMock()
			accountBalanceRepo := accountbalance_repo.NewAccountBalanceRepoMock()
			accountDetailRepo := accountdetail_repo.NewAccountDetailRepoMock()
			accountFlagRepo := accountflag_repo.NewAccountFlagRepoMock()
			debitCardRepo := debitcard_repo.NewDebitCardRepoMock()
			debitcardstatuRepo := debitcardstatus_repo.NewDebitCardStatusRepoMock()
			debitCardSDetailRepo := debitcarddetails_repo.NewDebitCardSDetailRepoMock()
			debitCarddesignRepo := debitcarddesign_repo.NewDebitCarddesignRepoMock()
			transectionRepo := transaction_repo.NewTransactionRepoMock()

			userRepo.On("InsertMock", mock.Anything, mock.Anything).Return(
				"test", tc.erruserInsertMock,
			)
			userRepo.On("Insert", mock.Anything, mock.Anything).Return(
				"test", tc.erruser,
			)

			userGreetingRepo.On("Insert", mock.Anything, mock.Anything).Return(
				tc.erruserGreeting,
			)

			bannerRepo.On("Insert", mock.Anything, mock.Anything).Return(
				tc.errBanner,
			)

			accountRepo.On("Insert", mock.Anything, mock.Anything).Return(
				"test", tc.errAccount,
			)
			accountBalanceRepo.On("Insert", mock.Anything, mock.Anything).Return(
				tc.errAccountBalance,
			)
			accountDetailRepo.On("Insert", mock.Anything, mock.Anything).Return(
				tc.errAccountDetail,
			)

			debitCardRepo.On("Insert", mock.Anything, mock.Anything).Return(
				"test", tc.errDebitCard,
			)
			debitcardstatuRepo.On("Insert", mock.Anything, mock.Anything).Return(
				tc.errDebitCardStatus,
			)
			debitCardSDetailRepo.On("Insert", mock.Anything, mock.Anything).Return(
				tc.errDebitCardDetail,
			)
			debitCarddesignRepo.On("Insert", mock.Anything, mock.Anything).Return(
				tc.errDebitCardDesign,
			)

			task := NewTaskInsertSvc(
				db,
				userRepo, userGreetingRepo,
				bannerRepo,
				accountRepo, accountBalanceRepo, accountDetailRepo, accountFlagRepo,
				debitCardRepo, debitcardstatuRepo, debitCardSDetailRepo, debitCarddesignRepo,
				transectionRepo,
			)

			task.InsertSimpleData()

		})
	}
}
