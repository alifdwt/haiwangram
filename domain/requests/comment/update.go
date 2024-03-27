package comment

import "github.com/go-playground/validator/v10"

type UpdateCommentRequest struct {
	Message string `json:"message" validate:"required"`
}

func (c *UpdateCommentRequest) Validate() error {
	validate := validator.New()

	err := validate.Struct(c)
	if err != nil {
		return err
	}

	return nil
}
