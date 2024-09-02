package main

import (
	"fmt"
	"mako/internal/fetcher"
	"net/url"
)

func main() {
	params := url.Values{}
	params.Add("limit", "50")
	//params.Add("offset_date", "2023-06-04T15:46:11.636835%2B07:00")
	params.Add("rated", "true")
	params.Add("sort_by", "date_edited")
	params.Add("key", "37c04fe6-a560-4549-b459-02309cf643ad")
	params.Add("locale", "ru_KZ")

	baseURL := "https://public-api.reviews.2gis.com/2.0/branches/9429940000796245/reviews"

	dgisfetcher := fetcher.NewDGISReviewFetcher(baseURL)

	response, err := dgisfetcher.FetchReview(params)
	if err != nil {
		return
	}

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
