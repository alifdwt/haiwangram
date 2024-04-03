package mapper

import (
	"github.com/alifdwt/haiwangram/domain/responses"
	"github.com/alifdwt/haiwangram/models"
)

type bookmarkMapper struct {
}

func NewBookmarkMapper() *bookmarkMapper {
	return &bookmarkMapper{}
}

func (m *bookmarkMapper) ToBookmarkResponse(request *models.Bookmark) *responses.BookmarkResponse {
	return &responses.BookmarkResponse{
		ID:      request.ID,
		PhotoID: request.PhotoID,
		UserID:  request.UserID,
	}
}

func (m *bookmarkMapper) ToBookmarkWithRelationResponse(request *models.Bookmark) *responses.BookmarkWithRelationResponse {
	return &responses.BookmarkWithRelationResponse{
		ID:      request.ID,
		PhotoID: request.PhotoID,
		UserID:  request.UserID,
		Photo:   *NewPhotoMapper().ToPhotoResponse(&request.Photo),
		User:    *NewUserMapper().ToUserResponse(&request.User),
	}
}

func (m *bookmarkMapper) ToBookmarkWithRelationResponses(requests *[]models.Bookmark) []responses.BookmarkWithRelationResponse {
	var bookmarkResponses []responses.BookmarkWithRelationResponse
	for _, request := range *requests {
		bookmarkResponses = append(bookmarkResponses, *m.ToBookmarkWithRelationResponse(&request))
	}

	if len(bookmarkResponses) > 0 {
		return bookmarkResponses
	}

	return []responses.BookmarkWithRelationResponse{}
}
