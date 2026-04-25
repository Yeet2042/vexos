package resourcecollection

type repository struct {
}

func NewRepository() Repository {
	return &repository{}
}
