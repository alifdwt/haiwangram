package mapper

type Mapper struct {
	UserMapper  UserMapping
	PhotoMapper PhotoMapping
}

func NewMapper() *Mapper {
	return &Mapper{
		UserMapper:  NewUserMapper(),
		PhotoMapper: NewPhotoMapper(),
	}
}
