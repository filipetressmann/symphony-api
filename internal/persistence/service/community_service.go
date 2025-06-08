package service

import (
	"errors"
	"symphony-api/internal/persistence/model"
	"symphony-api/internal/persistence/repository"
)

type CommunityService struct {
	communityRepository *repository.CommunityRepository
	userRepository *repository.UserRepository
}

func NewCommunityService(
	communityRepository *repository.CommunityRepository,
	userRepository *repository.UserRepository,
	) *CommunityService {
	return &CommunityService{
		communityRepository: communityRepository,
		userRepository: userRepository,
	}
}

func (service *CommunityService) AddUserToCommunity(username string, communityName string) error {
	user, err := service.userRepository.GetByUsername(username)

	if err != nil {
		return errors.New("user does not exist")
	}

	community, err := service.communityRepository.GetByName(communityName)

	if err != nil {
		return errors.New("community does not exist")
	}

	err = service.communityRepository.AddUserToCommunity(user, community)

	return err
}

func (service *CommunityService) ListUsersFromCommunity(communityName string) ([]*model.User, error) {
	community, err := service.communityRepository.GetByName(communityName)

	if err != nil {
		return nil, errors.New("community does not exist")
	}

	return service.communityRepository.ListUsersFromCommunity(community)
}

func (service *CommunityService) ListCommunitiesOfUser(username string) ([]*model.Community, error) {
	user, err := service.userRepository.GetByUsername(username)

	if err != nil {
		return nil, errors.New("community does not exist")
	}

	return service.userRepository.ListUserCommunities(user)
}