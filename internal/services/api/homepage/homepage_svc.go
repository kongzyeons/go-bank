package homepage_svc

import (
	"github.com/kongzyeons/go-bank/internal/models"
	usergreeting_repo "github.com/kongzyeons/go-bank/internal/repositories/user-greeting"
	"github.com/kongzyeons/go-bank/internal/utils/response"
)

type HomePageSvc interface {
	GetUserGreetings(req models.HomePageGetUserGreetingsReq) response.Response[*models.HomePageGetUserGreetingsRes]
}

type homePageSvc struct {
	userGreetingRepo usergreeting_repo.UserGreetingRepo
}

func NewHomePageSvc(userGreetingRepo usergreeting_repo.UserGreetingRepo) HomePageSvc {
	return &homePageSvc{
		userGreetingRepo: userGreetingRepo,
	}
}

func (svc *homePageSvc) GetUserGreetings(req models.HomePageGetUserGreetingsReq) response.Response[*models.HomePageGetUserGreetingsRes] {
	dataDB, err := svc.userGreetingRepo.GetByID(req.UserID)
	if err != nil {
		return response.InternalServerError[*models.HomePageGetUserGreetingsRes](err, err.Error())
	}
	if dataDB == nil {
		return response.Notfound[*models.HomePageGetUserGreetingsRes]("not found user id")
	}

	res := &models.HomePageGetUserGreetingsRes{
		Greeting: dataDB.Greeting,
	}
	return response.Ok(&res)
}
