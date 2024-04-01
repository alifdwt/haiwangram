package like

import "github.com/go-playground/validator/v10"

type DeleteLikeRequest struct {
	PhotoID int `json:"photo_id" validate:"required"`
}

func (l *DeleteLikeRequest) Validate() error {
	validate := validator.New()

	err := validate.Struct(l)
	if err != nil {
		return err
	}

	return nil
}
