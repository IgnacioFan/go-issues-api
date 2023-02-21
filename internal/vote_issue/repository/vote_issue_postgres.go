package repository

import (
	"go-issues-api/internal/model"
	"time"

	"gorm.io/gorm"
)

type VoteIssue struct {
	ID        uint `gorm:"primaryKey"`
	IssueId   uint `gorm:"foreignKey:ID"`
	UserId    uint `gorm:"foreignKey:ID"`
	Vote      int  `gorm:"not null, default: 0"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (vi *VoteIssue) toModel() *model.VoteIssue {
	return &model.VoteIssue{
		ID:        vi.ID,
		IssueId:   vi.IssueId,
		UserId:    vi.UserId,
		Vote:      vi.Vote,
		CreatedAt: vi.CreatedAt,
		UpdatedAt: vi.UpdatedAt,
	}
}

func toGorm(vi *model.VoteIssue) *VoteIssue {
	return &VoteIssue{
		IssueId: vi.IssueId,
		UserId:  vi.UserId,
		Vote:    vi.Vote,
	}
}

type VoteIssueRepository struct {
	DB *gorm.DB
}

func NewVoteIssueRepository(conn *gorm.DB) *VoteIssueRepository {
	return &VoteIssueRepository{
		DB: conn,
	}
}

func (repo *VoteIssueRepository) Create(vi *model.VoteIssue) (*model.VoteIssue, error) {
	var dbVoteIssue = toGorm(vi)
	res := repo.DB.Create(dbVoteIssue)
	return dbVoteIssue.toModel(), res.Error
}

func (repo *VoteIssueRepository) Update(vi *model.VoteIssue) (*model.VoteIssue, error) {
	var dbVoteIssue VoteIssue
	res := repo.DB.Where(&VoteIssue{IssueId: vi.IssueId, UserId: vi.UserId}).First(&dbVoteIssue)

	if res.Error != nil {
		return nil, res.Error
	} else if dbVoteIssue.Vote == vi.Vote {
		return dbVoteIssue.toModel(), nil
	} else {
		res = repo.DB.Model(&dbVoteIssue).Update("vote", vi.Vote)
		return dbVoteIssue.toModel(), nil
	}
}
