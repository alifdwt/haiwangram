package commentreply

import "github.com/go-playground/validator/v10"

type CreateCommentReplyRequest struct {
	Message   string `json:"message" validate:"required"`
	CommentID int    `json:"comment_id" validate:"required"`
}

func (c *CreateCommentReplyRequest) Validate() error {
	validate := validator.New()

	err := validate.Struct(c)
	if err != nil {
		return err
	}

	return nil
}
