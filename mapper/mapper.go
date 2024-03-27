package mapper

type Mapper struct {
	UserMapper         UserMapping
	PhotoMapper        PhotoMapping
	CommentMapper      CommentMapping
	CommentReplyMapper CommentReplyMapping
}

func NewMapper() *Mapper {
	return &Mapper{
		UserMapper:         NewUserMapper(),
		PhotoMapper:        NewPhotoMapper(),
		CommentMapper:      NewCommentMapper(),
		CommentReplyMapper: NewCommentReplyMapper(),
	}
}
