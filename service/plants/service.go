package plants

type Service struct {
	plantsRepo Repository
}

func NewService(plantsRepo Repository) Service {
	return Service{
		plantsRepo: plantsRepo,
	}
}
