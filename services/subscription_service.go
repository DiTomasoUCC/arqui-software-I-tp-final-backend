package services

import (
	"fmt"

	"github.com/DiTomasoUCC/arqui-software-I-tp-final-backend/clients"
	"github.com/DiTomasoUCC/arqui-software-I-tp-final-backend/dto"
)

func AddSubscription(subscriptionDto dto.SubscriptionDto) (dto.SubscriptionDto, error) {
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
