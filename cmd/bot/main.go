package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"time"
)

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

func main() {
	params := url.Values{}
	params.Add("limit", "50")
	params.Add("offset_date", "2024-06-09T00:16:46.728493%2B07:00")
	params.Add("is_advertiser", "false")
	params.Add("fields", "meta.providers,meta.branch_rating,meta.branch_reviews_count,meta.total_count,reviews.hiding_reason,reviews.is_verified")
	params.Add("without_my_first_review", "false")
	params.Add("rated", "true")
	params.Add("sort_by", "date_edited")
	params.Add("key", "37c04fe6-a560-4549-b459-02309cf643ad")
	params.Add("locale", "ru_KZ")

	baseURL := "https://public-api.reviews.2gis.com/2.0/branches/9429940000796245/reviews" + params.Encode()
	reviewClient := http.Client{
		Timeout: time.Second * 2,
	}

	req, err := http.NewRequest(http.MethodGet, baseURL, nil)
	if err != nil {
		log.Fatal(err)
	}

	res, getErr := reviewClient.Do(req)
	if getErr != nil {
		fmt.Println(getErr)
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	fmt.Println("HTTP Status Code:", res.StatusCode)

	if res.StatusCode != http.StatusOK {
		log.Fatalf("Request failed with status code: %d", res.StatusCode)
	}

	body, readErr := io.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	response := Response{}
	jsonErr := json.Unmarshal(body, &response)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	fmt.Println(res.StatusCode)
	fmt.Println(response.Meta.Code)

	for i, r := range response.Reviews {
		fmt.Println("Review", (i + 1), ": ", r.User.FirstName, r.User.LastName)
		fmt.Println("Total count of reviews: ", r.User.ReviewsCount)
		fmt.Println("----------------------------------------")
		fmt.Println("Text: ", r.Text)
		fmt.Println("Rating: ", r.Rating)
		fmt.Println("Likes_count: ", r.LikesCount)
		fmt.Println("Date_created: ", r.DateCreated)
		fmt.Println()
	}

}
