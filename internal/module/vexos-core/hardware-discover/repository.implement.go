package hardwarediscover

type repository struct {
}

func NewRepository() Repository {
	return &repository{}
}
