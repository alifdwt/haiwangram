package mapper

type Mapper struct {
	UserMapper    UserMapping
	PhotoMapper   PhotoMapping
	CommentMapper CommentMapping
}

func NewMapper() *Mapper {
	return &Mapper{
		UserMapper:    NewUserMapper(),
		PhotoMapper:   NewPhotoMapper(),
		CommentMapper: NewCommentMapper(),
	}
}
