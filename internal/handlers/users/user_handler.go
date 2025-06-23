package user_handlers

import (
	"errors"
	base_handlers "symphony-api/internal/handlers/base"
	request_model "symphony-api/internal/handlers/model"
	"symphony-api/internal/persistence/connectors/neo4j"
	"symphony-api/internal/persistence/connectors/postgres"
	"symphony-api/internal/persistence/repository"
	"symphony-api/internal/persistence/service"
	"symphony-api/internal/server"
)

type UserHandler struct {
	repository *repository.UserRepository
	communityService *service.CommunityService
}

func NewUserHandler(connection postgres.PostgreConnection, neo4jConnection neo4j.Neo4jConnection) *UserHandler {
	userRepository := repository.NewUserRepository(connection, neo4jConnection)
	return &UserHandler{
		repository: userRepository,
		communityService: service.NewCommunityService(
			repository.NewCommunityRepository(connection),
			userRepository,
		),
	}
}

func (handler *UserHandler) AddRoutes(server server.Server) {
	server.AddRoute(
		"/api/user/create", 
		base_handlers.CreatePostMethodHandler(handler.CreateUserHandler),
	)
	server.AddRoute(
		"/api/user/get_by_username", 
		base_handlers.CreateGetMethodHandler(handler.GetUserByUsername),
	)
	server.AddRoute(
		"/api/user/list_communities", 
		base_handlers.CreateGetMethodHandler(handler.ListUserCommunities),
	)

	server.AddRoute(
		"/api/user/create_friendship", 
		base_handlers.CreatePostMethodHandler(handler.CreateFriendship),
	)

	server.AddRoute(
		"/api/user/list_friends", 
		base_handlers.CreateGetMethodHandler(handler.GetUserFriends),
	)

	server.AddRoute(
		"/api/user/like_genre", 
		base_handlers.CreatePostMethodHandler(handler.LikeGenre),
	)

	server.AddRoute(
		"/api/user/list_liked_genres", 
		base_handlers.CreateGetMethodHandler(handler.ListLikedGenres),
	)

	server.AddRoute(
		"/api/user/get_friends_recommendations_on_genre", 
		base_handlers.CreateGetMethodHandler(handler.GetFriendRecommendationByGenre),
	)
}

// CreateUserHandler handles the creation of a new user.
//	@Summary		Create a new user
//	@Description	Creates a new user in the system.
//	@Tags			User
//	@Accept			json
//	@Produce		json
//	@Param			user	body		request_model.CreateUserRequest	true	"User data"
//	@Success		200		{object}	request_model.SuccessCreationResponse
//	@Failure		400		{object}	map[string]string	"Invalid Input"
//	@Failure		500		{object}	map[string]string	"Internal Server Error"
//	@Router			/api/user/create [post]
func (handler *UserHandler) CreateUserHandler(request request_model.CreateUserRequest) (*request_model.SuccessCreationResponse, error) {
	
	err := handler.repository.Put(request.ToUser())

	// Change this later if necessary. We should check why the creation failed and give a better
	// answer to the requester.
	if err != nil {
		return nil, errors.New("error creating user")
	}

	return request_model.NewSuccessCreationResponse("Successfully created user"), nil
}

// Return user data based on their username
//	@Summary		Get a user by its name
//	@Description	Return user data based on their username
//	@Tags			User
//	@Accept			json
//	@Produce		json
//	@Param			post	body		request_model.GetUserByUsernameRequest	true	"User data"
//	@Success		200		{object}	request_model.UserResponse
//	@Failure		400		{object}	map[string]string	"Invalid Input"
//	@Failure		500		{object}	map[string]string	"Internal Server Error"
//	@Router			/api/user/get_by_username [get]
func (handler *UserHandler) GetUserByUsername(request request_model.GetUserByUsernameRequest) (*request_model.UserResponse, error) {
	user, err := handler.repository.GetByUsername(request.Username)

	if err != nil {
        return nil, errors.New("error fetching user")
	}

	return request_model.NewUserResponse(user), nil
}

// Return all communities a user is part of
//	@Summary		Get all communities of a user
//	@Description	Return all communities a user is part of
//	@Tags			User
//	@Accept			json
//	@Produce		json
//	@Param			post	body 		request_model.ListUserCommunitiesRequest	true	"User data"
//	@Success		200		{object}	request_model.ListUserCommunitiesResponse
//	@Failure		400		{object}	map[string]string	"Invalid Input"
//	@Failure		500		{object}	map[string]string	"Internal Server Error"
//	@Router			/api/user/list_communities [get]
func (handler *UserHandler) ListUserCommunities(request request_model.ListUserCommunitiesRequest) (*request_model.ListUserCommunitiesResponse, error) {
	communities, err := handler.communityService.ListCommunitiesOfUser(request.Username)

	if err != nil {
		return nil, errors.New("error listing user communities")
	}

	communitiesResponseList := make([]*request_model.CommunityDataResponse, 0)

	for _, community := range communities {
		communitiesResponseList = append(communitiesResponseList, request_model.NewCommunityDataResponse(community))
	}

	return &request_model.ListUserCommunitiesResponse{
		Communities: communitiesResponseList,
	}, nil
}

// Creates a friendship between two users
//	@Summary		Create a friendship 
//	@Description	Creates a friendship between user with username1 and user with username2
//	@Tags			User
//	@Accept			json
//	@Produce		json
//	@Param			post	body		request_model.CreateFriendshipRequest	true	"User data"
//	@Success		200		{object}	request_model.SuccessCreationResponse
//	@Failure		400		{object}	map[string]string	"Invalid Input"
//	@Failure		500		{object}	map[string]string	"Internal Server Error"
//	@Router			/api/user/create_friendship [post]
func (handler *UserHandler) CreateFriendship(request request_model.CreateFriendshipRequest) (*request_model.SuccessCreationResponse, error) {
	err := handler.repository.AddFriendship(request.Username1, request.Username2)

	if err != nil {
		return nil, err
	}

	return request_model.NewSuccessCreationResponse("Successfully created friendship"), nil
}

// List all friendship a user has
//	@Summary		List friends of a user
//	@Description	List all friends of a user
//	@Tags			User
//	@Accept			json
//	@Produce		json
//	@Param			post	body		request_model.GetUserFriendsRequest	true	"User data"
//	@Success		200		{object}	request_model.GetUserFriendsResponse
//	@Failure		400		{object}	map[string]string	"Invalid Input"
//	@Failure		500		{object}	map[string]string	"Internal Server Error"
//	@Router			/api/user/list_friends [get]
func (handler *UserHandler) GetUserFriends(request request_model.GetUserFriendsRequest) (*request_model.GetUserFriendsResponse, error) {
	friends, err := handler.repository.ListFriendshipsByUsername(request.Username)

	if err != nil {
		return nil, err
	}

	friendsModel := make([]*request_model.UserResponse, 0)

	for _, friend := range friends {
		friendsModel = append(friendsModel, request_model.NewUserResponse(friend))
	}

	return &request_model.GetUserFriendsResponse{
		Friends: friendsModel,
	}, nil
}

// Marks a genre as liked by a user. Genre can be any string.
//	@Summary		Marks a genre as liked by a user.
//	@Description	Marks a genre as liked by a user. Genre can be any string.
//	@Tags			User
//	@Accept			json
//	@Produce		json
//	@Param			post	body		request_model.LikeGenreRequest	true	"User data"
//	@Success		200		{object}	request_model.SuccessCreationResponse
//	@Failure		400		{object}	map[string]string	"Invalid Input"
//	@Failure		500		{object}	map[string]string	"Internal Server Error"
//	@Router			/api/user/like_genre [post]
func (handler *UserHandler) LikeGenre(request request_model.LikeGenreRequest) (*request_model.SuccessCreationResponse, error) {
	err := handler.repository.LikeGenre(request.Username, request.GenreName)

	if err != nil {
		return nil, err
	}

	return request_model.NewSuccessCreationResponse("Successfully liked genre"), nil
}

// List all genres liked by a user
//	@Summary		List all genres liked by a user
//	@Description	List all genres liked by a user
//	@Tags			User
//	@Accept			json
//	@Produce		json
//	@Param			post	body		request_model.GetLikedGenresRequest	true	"User data"
//	@Success		200		{object}	request_model.GetLikedGenresResponse
//	@Failure		400		{object}	map[string]string	"Invalid Input"
//	@Failure		500		{object}	map[string]string	"Internal Server Error"
//	@Router			/api/user/list_liked_genres [get]
func (handler *UserHandler) ListLikedGenres(request request_model.GetLikedGenresRequest) (*request_model.GetLikedGenresResponse, error) {
	genres, err := handler.repository.ListLikedGenres(request.Username)

	if err != nil {
		return nil, err
	}

	return &request_model.GetLikedGenresResponse{
		Genres: genres,
	}, nil
}

// This API returns user that likes the same genre of the specified user
//	@Summary		Returns recommendations of user that likes the same genre
//	@Description	This API returns user that likes the same genre of the specified user
//	@Tags			User
//	@Accept			json
//	@Produce		json
//	@Param			post	body		request_model.GetFriendRecommendationByGenreRequest	true	"User data"
//	@Success		200		{object}	request_model.GetFriendRecommendationByGenreResponse
//	@Failure		400		{object}	map[string]string	"Invalid Input"
//	@Failure		500		{object}	map[string]string	"Internal Server Error"
//	@Router			/api/user/get_friends_recommendations_on_genre [get]
func (handler *UserHandler) GetFriendRecommendationByGenre(request request_model.GetFriendRecommendationByGenreRequest) (*request_model.GetFriendRecommendationByGenreResponse, error) {
	friends, err := handler.repository.GetRecommendationsOnGenre(request.Username)

	if err != nil {
		return nil, err
	}

	friendsModel := make([]*request_model.UserResponse, 0)

	for _, friend := range friends {
		friendsModel = append(friendsModel, request_model.NewUserResponse(friend))
	}

	return &request_model.GetFriendRecommendationByGenreResponse{
		Friends: friendsModel,
	}, nil
}