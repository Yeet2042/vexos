package resourcecollection

type usecase struct {
}

func NewUsecase() Usecase {
	return &usecase{}
}
