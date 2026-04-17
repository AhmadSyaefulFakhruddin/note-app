package tags

import "context"

type Service interface {
	GetTags(c context.Context) ([]NewTagMinimalDto, error)
	CreateMultipleTags(c context.Context, tagNames []string) ([]NewTagMinimalDto, error)
}

type service struct {
	repo Repository
}

func NewService(r Repository) Service {
	return &service{
		repo: r,
	}
}

func (s *service) GetTags(c context.Context) ([]NewTagMinimalDto, error) {
	tagDbs, err := s.repo.GetTags(c)

	if err != nil {
		return []NewTagMinimalDto{}, err
	}

	return toNewTagMinimalDTOs(tagDbs), nil
}

func (s *service) CreateMultipleTags(c context.Context, tagNames []string) ([]NewTagMinimalDto, error) {
	tagDbs, err := s.repo.CreateMultipleTags(c, tagNames)

	if err != nil {
		return []NewTagMinimalDto{}, err
	}

	return toNewTagMinimalDTOs(tagDbs), nil
}
