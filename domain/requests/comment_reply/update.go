package commentreply

import "github.com/go-playground/validator/v10"

type UpdateCommentReplyRequest struct {
	Message string `json:"message" validate:"required"`
}

func (c *UpdateCommentReplyRequest) Validate() error {
	validate := validator.New()

	err := validate.Struct(c)
	if err != nil {
		return err
	}

	return nil
}
