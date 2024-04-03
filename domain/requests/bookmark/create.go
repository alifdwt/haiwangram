package bookmark

import "github.com/go-playground/validator/v10"

type CreateBookmarkRequest struct {
	PhotoID int `json:"photo_id" validate:"required"`
}

func (b *CreateBookmarkRequest) Validate() error {
	validate := validator.New()

	err := validate.Struct(b)
	if err != nil {
		return err
	}

	return nil
}
