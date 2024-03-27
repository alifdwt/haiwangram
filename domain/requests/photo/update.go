package photo

import "github.com/go-playground/validator/v10"

type UpdatePhotoRequest struct {
	Title    string `json:"title" validate:"required"`
	Caption  string `json:"caption"`
	PhotoURL string `json:"photo_url" validate:"required,url"`
}

func (l *UpdatePhotoRequest) Validate() error {
	validate := validator.New()

	err := validate.Struct(l)

	if err != nil {
		return err
	}

	return nil
}
