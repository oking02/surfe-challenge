package memory

import (
	"context"
	"github.com/oking02/surfe-challenge/internal/domain"
)

type userKey struct {
	clientID domain.ClientID
	userID   domain.UserID
}

type UserRepository struct {
	d map[userKey]domain.User
}

func NewUserRepository(userData []domain.User) *UserRepository {

	d := make(map[userKey]domain.User)
	for _, user := range userData {
		d[userKey{
			clientID: user.ClientID,
			userID:   user.ID,
		}] = user
	}
	return &UserRepository{
		d: d,
	}
}

func (u *UserRepository) GetUser(ctx context.Context, clientID domain.ClientID, id domain.UserID) (domain.User, error) {
	if u.d == nil {
		return domain.User{}, domain.ErrUserNotFound
	}

	user, ok := u.d[userKey{
		clientID: clientID,
		userID:   id,
	}]
	if !ok {
		return domain.User{}, domain.ErrUserNotFound
	}

	return user, nil
}

func (u *UserRepository) ListUsers(ctx context.Context, clientID domain.ClientID) ([]domain.User, error) {
	if u.d == nil {
		return []domain.User{}, nil
	}

	var results []domain.User
	for key, user := range u.d {
		if key.clientID == clientID {
			results = append(results, user)
		}
	}

	return results, nil
}
