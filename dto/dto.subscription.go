package dto

import "time"

type SubscriptionDto struct {
	Id           int       `json:"subscription_id"`
	CourseId     int       `json:"course_id"`
	UserId       int       `json:"user_id"`
	CreationTime time.Time `json:"creation_time"`
}

type CourseSubscriptionsResponse struct {
	Results []UserDto `json:"results"`
}

type UserCoursesResponse struct {
	Results []CourseDto `json:"results"`
}
