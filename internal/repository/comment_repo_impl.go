package repository

import (
	"github.com/iki-rumondor/project2-grup9/internal/domain"
	"gorm.io/gorm"
)

type RepoImplementation struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) CommentRepository {
	return &RepoImplementation{
		db: db,
	}
}

func (r *RepoImplementation) CreateComment(comment *domain.Comment) (*domain.Comment, error) {
	if err := r.db.Create(&comment).Error; err != nil {
		return nil, err
	}
	return comment, nil
}

func (r *RepoImplementation) FindComment(uint) ([]domain.Comment, error) {
	var comments []domain.Comment

	if err := r.db.Find(&comments).Error; err != nil {
		return nil, err
	}

	return comments, nil
}

func (r *RepoImplementation) UpdateComment(comment *domain.UpdateComment) (*domain.Comment, error) {
	var result domain.Comment
	if err := r.db.Model(&domain.Comment{}).Where("id = ?", comment.ID).Updates(&comment).First(&result).Error; err != nil {
		return nil, err
	}
	return &result, nil
}

func (r *RepoImplementation) DeleteComment(comment *domain.Comment) error {
	if err := r.db.Delete(&comment).Error; err != nil {
		return err
	}
	return nil
}

func (r *RepoImplementation) FindUser(id uint) (*domain.User, error) {
	var user domain.User
	if err := r.db.First(&user, "id = ?", id).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *RepoImplementation) FindPhoto(id uint) (*domain.Photo, error) {
	var photo domain.Photo
	if err := r.db.First(&photo, "id = ?", id).Error; err != nil {
		return nil, err
	}

	return &photo, nil
}

func (r *RepoImplementation) FindComments(userID uint) ([]*domain.Comment, error) {
	var comments []*domain.Comment
	if err := r.db.Where("user_id = ?", userID).Find(&comments).Error; err != nil {
		return nil, err
	}
	return comments, nil
}
