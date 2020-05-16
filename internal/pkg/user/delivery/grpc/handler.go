package grpc

import (
	"2020_1_Color_noise/internal/models"
	. "2020_1_Color_noise/internal/pkg/error"
	userService "2020_1_Color_noise/internal/pkg/proto/user"
	"2020_1_Color_noise/internal/pkg/user"
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserService struct {
	usecase user.IUsecase
}

func NewUserService(usecase user.IUsecase) *UserService {
	return &UserService{
		usecase,
	}
}

func (us *UserService) Create(ctx context.Context, in *userService.SignUp) (*userService.User, error) {
	input := &models.SignUpInput{
		Email: in.Email,
		Login: in.Login,
		Password: in.Password,
		}

	u, err := us.usecase.Create(input)
	if err != nil {
		return nil, status.Error(codes.Code(uint(GetType(err))), Wrap(err, "GRPC create: Creating user error").Error())
	}

	userProto := &userService.User{
		Id:     int64(u.Id),
		Email:  u.Email,
		Login:  u.Login,
		Avatar: u.Avatar,
	}

	return userProto, nil
}

func (us *UserService) GetById(ctx context.Context, in *userService.UserID) (*userService.User, error) {
	u, err := us.usecase.GetById(uint(in.Id))
	if err != nil {
		return nil, status.Error(codes.Code(uint(GetType(err))), Wrapf(err, "GRPC GetById error, id: %d", in.Id).Error())
	}

	userProto := &userService.User{
		Id:     int64(u.Id),
		Email:  u.Email,
		Login:  u.Login,
		Avatar: u.Avatar,
		About:  u.About,
		Subscribers: int64(u.Subscribers),
		Subscriptions: int64(u.Subscriptions),
	}

	return userProto, nil
}

func (us *UserService) GetByLogin(ctx context.Context, in *userService.Login) (*userService.User, error) {
	u, err := us.usecase.GetByLogin(in.Login)
	if err != nil {
		return nil, status.Error(codes.Code(uint(GetType(err))), Wrap(err, "GRPC GetById error").Error())
	}


	userProto := &userService.User{
		Id:     int64(u.Id),
		Email:  u.Email,
		EncryptedPassword: u.EncryptedPassword,
		Login:  u.Login,
		Avatar: u.Avatar,
		About:  u.About,
		Subscribers: int64(u.Subscribers),
		Subscriptions: int64(u.Subscriptions),
	}

	return userProto, nil
}

func (us *UserService) UpdateProfile(ctx context.Context, in *userService.Profile) (*userService.Nothing, error) {
	input := &models.UpdateProfileInput {
		Email: in.Input.Email,
		Login: in.Input.Login,
	}
	err := us.usecase.UpdateProfile(uint(in.Id.Id), input)
	if err != nil {
		return nil, status.Error(codes.Code(uint(GetType(err))), Wrap(err, "GRPC UpdateProfile error").Error())
	}

	return &userService.Nothing{}, nil
}

func (us *UserService) UpdateDescription(ctx context.Context, in *userService.Description) (*userService.Nothing, error) {
	input := &models.UpdateDescriptionInput{
		Description: in.Description,
	}

	err := us.usecase.UpdateDescription(uint(in.Id.Id), input)
	if err != nil {
		return nil, status.Error(codes.Code(uint(GetType(err))), Wrap(err, "GRPC UpdateDescription error").Error())
	}

	return &userService.Nothing{}, nil
}

func (us *UserService) UpdatePassword(ctx context.Context, in *userService.Password) (*userService.Nothing, error) {
	input := &models.UpdatePasswordInput{
		Password: in.Password,
	}

	err := us.usecase.UpdatePassword(uint(in.Id.Id), input)
	if err != nil {
		return nil, status.Error(codes.Code(uint(GetType(err))), Wrap(err, "GRPC UpdateDPassword error").Error())
	}

	return &userService.Nothing{}, nil
}

func (us *UserService) UpdateAvatar(ctx context.Context, in *userService.Avatar) (*userService.Address, error) {
	image, err := us.usecase.UpdateAvatar(uint(in.Id.Id), in.Avatar)
	if err != nil {
		return nil, status.Error(codes.Code(uint(GetType(err))), Wrap(err, "GRPC UpdateAvatar error").Error())
	}

	return &userService.Address{Avatar: image}, nil
}

func (us *UserService) Follow(ctx context.Context, in *userService.Following) (*userService.Nothing, error) {
	err := us.usecase.Follow(uint(in.Id.Id), uint(in.SubId.Id))
	if err != nil {
		return nil, status.Error(codes.Code(uint(GetType(err))), Wrap(err, "GRPC Follow error").Error())
	}

	return &userService.Nothing{}, nil
}

func (us *UserService) Unfollow(ctx context.Context, in *userService.Following) (*userService.Nothing, error) {
	err := us.usecase.Unfollow(uint(in.Id.Id), uint(in.SubId.Id))
	if err != nil {
		return nil, status.Error(codes.Code(uint(GetType(err))), Wrap(err, "GRPC Unfollow error").Error())
	}

	return &userService.Nothing{}, nil
}

func (us *UserService) GetSubscribers(ctx context.Context, in *userService.Sub) (*userService.Users, error) {
	users, err := us.usecase.GetSubscribers(uint(in.Id.Id), int(in.Start), int(in.Limit))
	if err != nil {
		return nil, status.Error(codes.Code(uint(GetType(err))), Wrap(err, "GRPC GetSubscribers error").Error())
	}

	u := &userService.Users{}
	for _, user := range users {
		u.Users = append(u.Users,
			&userService.User{
				Id:     int64(user.Id),
				Email:  user.Email,
				Login:  user.Login,
				Avatar: user.Avatar,
				About:  user.About,
				Subscribers: int64(user.Subscribers),
				Subscriptions: int64(user.Subscriptions),
			},
			)
	}

	return u, nil
}

func (us *UserService) GetSubscriptions(ctx context.Context, in *userService.Sub) (*userService.Users, error) {
	users, err := us.usecase.GetSubscriptions(uint(in.Id.Id), int(in.Start), int(in.Limit))
	if err != nil {
		return nil, status.Error(codes.Code(uint(GetType(err))), Wrap(err, "GRPC GetSubscriptions error").Error())
	}

	u := &userService.Users{}
	for _, user := range users {
		u.Users = append(u.Users,
			&userService.User{
				Id:     int64(user.Id),
				Email:  user.Email,
				Login:  user.Login,
				Avatar: user.Avatar,
				About:  user.About,
				Subscribers: int64(user.Subscribers),
				Subscriptions: int64(user.Subscriptions),
			},
		)
	}

	return u, nil
}

func (us *UserService) Search(ctx context.Context, in *userService.Searching) (*userService.Users, error) {
	users, err := us.usecase.Search(in.Login.Login, int(in.Start), int(in.Limit))
	if err != nil {
		return nil, status.Error(codes.Code(uint(GetType(err))), Wrap(err, "GRPC Search error").Error())
	}

	u := &userService.Users{}
	for _, user := range users {
		u.Users = append(u.Users,
			&userService.User{
				Id:     int64(user.Id),
				Email:  user.Email,
				Login:  user.Login,
				Avatar: user.Avatar,
				About:  user.About,
				Subscribers: int64(user.Subscribers),
				Subscriptions: int64(user.Subscriptions),
			},
		)
	}

	return u, nil
}

func (us *UserService) UpdatePreferences(ctx context.Context, in *userService.Pref) (*userService.Nothing, error) {
	err := us.usecase.UpdatePreferences(uint(in.UserId), in.Preferences)
	if err != nil {
		return nil, status.Error(codes.Code(uint(GetType(err))), Wrap(err, "GRPC UpdatePreferences error").Error())
	}

	return &userService.Nothing{}, nil
}