package services

import (
	"fmt"

	"github.com/DiTomasoUCC/arqui-software-I-tp-final-backend/clients"
	"github.com/DiTomasoUCC/arqui-software-I-tp-final-backend/dto"
)

func AddSubscription(subscriptionDto dto.SubscriptionDto) (dto.SubscriptionDto, error) {
	isSubscribed, err := isUserSubscribed(subscriptionDto.UserId, subscriptionDto.CourseId)

	if err != nil {
		return dto.SubscriptionDto{}, fmt.Errorf("error getting subscription: %w", err)
	}

	if isSubscribed {
		return dto.SubscriptionDto{}, fmt.Errorf("user is already subscribed to this course")
	}

	subscription, err := clients.CreateSubscription(subscriptionDto.CourseId, subscriptionDto.UserId)

	if err != nil {
		return dto.SubscriptionDto{}, fmt.Errorf("error creating subscription: %w", err)
	}

	return dto.SubscriptionDto{
		Id:           subscription.ID,
		CourseId:     subscription.CourseID,
		UserId:       subscription.UserID,
		CreationTime: subscription.CreationTime,
	}, nil
}

func GetSubscribedUsers(courseId int) ([]dto.UserDto, error) {
	users, err := clients.GetSubscribedUsers(courseId)

	if err != nil {
		return nil, fmt.Errorf("error getting subscribed users: %w", err)
	}

	result := make([]dto.UserDto, 0)

	for _, user := range users {
		result = append(result, dto.UserDto{
			UserId:       user.ID,
			Username:     user.UserName,
			Email:        user.Email,
			CreationTime: user.CreationTime,
		})
	}

	return result, nil
}

func GetUserCourses(userId int) ([]dto.CourseDto, error) {
	courses, err := clients.GetUserCourses(userId)

	if err != nil {
		return nil, fmt.Errorf("error getting user courses: %w", err)
	}

	result := make([]dto.CourseDto, 0)

	for _, course := range courses {
		result = append(result, dto.CourseDto{
			Id:           course.ID,
			Name:         course.Name,
			Description:  course.Description,
			InstructorId: course.InstructorID,
			Category:     course.Category,
			Requirements: course.Requirements,
			Length:       course.Length,
			ImageURL:     course.ImageURL,
			CreationTime: course.CreationTime,
			LastUpdated:  course.LastUpdated,
		})
	}

	return result, nil
}
