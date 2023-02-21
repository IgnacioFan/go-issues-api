package model

import "time"

type VoteIssue struct {
	ID        uint `json:"id"`
	IssueId   uint `json:"issue_id"`
	UserId    uint `json:"userid"`
	Vote      int  `json:"vote"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type VoteIssueRepository interface {
	Create(vi *VoteIssue) (*VoteIssue, error)
	Update(vi *VoteIssue) (*VoteIssue, error)
}
