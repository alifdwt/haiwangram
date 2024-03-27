package photo

import "github.com/go-playground/validator/v10"

type CreatePhotoRequest struct {
	Title    string `json:"title" validate:"required"`
	Caption  string `json:"caption"`
	PhotoURL string `json:"photo_url" validate:"required,url"`
}

func (c *CreatePhotoRequest) Validate() error {
	validate := validator.New()

	err := validate.Struct(c)

	if err != nil {
		return err
	}

	return nil
}
