package comment

import "github.com/go-playground/validator/v10"

type CreateCommentRequest struct {
	Message string `json:"message" validate:"required"`
	PhotoID int    `json:"photo_id" validate:"required"`
}

func (c *CreateCommentRequest) Validate() error {
	validate := validator.New()

	err := validate.Struct(c)
	if err != nil {
		return err
	}

	return nil
}
