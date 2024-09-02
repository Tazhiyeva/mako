package models

import "time"

type Response struct {
	Meta    Meta     `json:"meta"`
	Reviews []Review `json:"reviews"`
}

type Meta struct {
	BranchRating       float64 `json:"branch_rating"`
	BranchReviewsCount int     `json:"branch_reviews_count"`
	Code               int     `json:"code"`
	TotalCount         int     `json:"total_count"`
}

type Review struct {
	Text        string    `json:"text"`
	Rating      int       `json:"rating"`
	LikesCount  int       `json:"likes_count"`
	DateCreated time.Time `json:"date_created"`
	User        User      `json:"user"`
}

type User struct {
	ReviewsCount int    `json:"reviews_count"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	Name         string `json:"name"`
}
