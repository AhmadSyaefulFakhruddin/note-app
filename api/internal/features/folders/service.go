package folders

import (
	"context"
)

type Service interface {
	GetFolders(ctx context.Context) ([]FolderResponse, error)
	CreateFolder(ctx context.Context, newFolderName string) (FolderResponse, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{
		repo: repo,
	}
}

func (s *service) GetFolders(ctx context.Context) ([]FolderResponse, error) {
	folderDbs, err := s.repo.GetFolders(ctx)

	if err != nil {
		return nil, err
	}

	folderResponses := make([]FolderResponse, len(folderDbs))

	for i := range folderDbs {
		folderResponses[i].ID = folderDbs[i].ID
		folderResponses[i].Name = folderDbs[i].Name
	}

	return folderResponses, nil
}

func (s *service) CreateFolder(ctx context.Context, newFolderName string) (FolderResponse, error) {
	newFolderResponse, err := s.repo.CreateFolder(ctx, FolderDb{
		Name: newFolderName,
	})

	if err != nil {
		return FolderResponse{}, err
	}

	return FolderResponse{Name: newFolderName, ID: newFolderResponse.ID}, nil
}
