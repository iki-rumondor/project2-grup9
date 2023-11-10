package repository

import (
	"github.com/iki-rumondor/project2-grup9/internal/domain"
	"gorm.io/gorm"
)

type CommentRepoImplementation struct {
	db *gorm.DB
}

func NewCommentRepository(db *gorm.DB) CommentRepository {
	return &CommentRepoImplementation{
		db: db,
	}
}

func (r *CommentRepoImplementation) CreateComment(comment *domain.Comment) (*domain.Comment, error) {
	if err := r.db.Save(comment).Error; err != nil {
		return nil, err
	}
	return comment, nil
}

func (r *CommentRepoImplementation) UpdateComment(comment *domain.Comment) (*domain.Comment, error) {
	var result domain.Comment
	if err := r.db.Model(&comment).Updates(&comment).First(&result).Error; err != nil {
		return nil, err
	}
	return &result, nil
}

func (r *CommentRepoImplementation) DeleteComment(comment *domain.Comment) error {
	if err := r.db.Delete(&comment).Error; err != nil {
		return err
	}
	return nil
}

func (r *CommentRepoImplementation) FindComments(userID uint) ([]*domain.Comment, error) {
	var comments []*domain.Comment
	if err := r.db.Preload("User").Preload("Photo").Find(&comments, "user_id = ?", userID).Error; err != nil {
		return nil, err
	}
	return comments, nil
}
