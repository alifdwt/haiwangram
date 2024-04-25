package follow

import "github.com/go-playground/validator/v10"

type CreateFollowRequest struct {
	UserId int `json:"user_id" validate:"required"`
}

func (l *CreateFollowRequest) Validate() error {
	validate := validator.New()

	err := validate.Struct(l)
	if err != nil {
		return err
	}

	return nil
}
