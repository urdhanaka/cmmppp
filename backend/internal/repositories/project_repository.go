package repositories

import (
	"cmmppp/internal/entities"
	"context"
	"errors"

	"gorm.io/gorm"
)

var ErrProjectNotFound = errors.New("project not found")

type (
	ProjectRepository interface{}

	projectRepository struct {
		db *gorm.DB
	}
)

func NewProjectRepository(db *gorm.DB) ProjectRepository {
	return &projectRepository{
		db: db,
	}
}

func (r *projectRepository) PostProject(ctx context.Context, project *entities.Project) error {
	if err := gorm.G[entities.Project](r.db).Create(ctx, project); err != nil {
		return errorMapper(err)
	}

	return nil
}

func (r *projectRepository) DeleteProject(ctx context.Context, project *entities.Project) error {
	rowsAffected, err := gorm.G[entities.Project](r.db).Where(project).Delete(ctx)
	if rowsAffected == 0 {
		return ErrProjectNotFound
	}
	if err != nil {
		return errorMapper(err)
	}

	return nil
}
