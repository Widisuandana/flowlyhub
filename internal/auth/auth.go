package auth

import (
	"context"
	"errors"
	"flowlyhub/internal/db/sqlc"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/jackc/pgx/v5"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	queries *sqlc.Queries
	config  *Config
}

type Config struct {
	JWTSecret string
}

func NewAuthService(queries *sqlc.Queries, config *Config) *AuthService {
	return &AuthService{queries: queries, config: config}
}

type RegisterInput struct {
	Email    string
	Password string
	Name     string
	Role     string
}

type LoginInput struct {
	Email    string
	Password string
}

type UpdateUserInput struct {
	ID       int32
	Email    string
	Password string
	Name     string
	Role     string
}

type Claims struct {
	UserID int32  `json:"user_id"`
	Email  string `json:"email"`
	Role   string `json:"role"`
	Name   string `json:"name"`
	jwt.StandardClaims
}

type User struct {
	ID    int32  `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Role  string `json:"role"`
}

func (s *AuthService) Register(ctx context.Context, input RegisterInput) (sqlc.User, error) {
	if input.Role != "owner" && input.Role != "staff" {
		return sqlc.User{}, errors.New("invalid role")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return sqlc.User{}, err
	}

	user, err := s.queries.CreateUser(ctx, sqlc.CreateUserParams{
		Email:    input.Email,
		Password: string(hashedPassword),
		Role:     input.Role,
		Name:     input.Name,
	})
	if err != nil {
		return sqlc.User{}, err
	}

	return sqlc.User{
		ID:        user.ID,
		Email:     user.Email,
		Password:  string(hashedPassword),
		Role:      user.Role,
		Name:      user.Name,
		CreatedAt: user.CreatedAt,
	}, nil
}

func (s *AuthService) Login(ctx context.Context, input LoginInput) (string, error) {
	user, err := s.queries.GetUserByEmail(ctx, input.Email)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return "", errors.New("invalid email or password")
		}
		return "", err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		return "", errors.New("invalid email or password")
	}

	claims := Claims{
		UserID: user.ID,
		Email:  user.Email,
		Role:   user.Role,
		Name:   user.Name,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(s.config.JWTSecret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (s *AuthService) UpdateUser(ctx context.Context, input UpdateUserInput) (sqlc.User, error) {
	if input.Role != "" && input.Role != "owner" && input.Role != "staff" {
		return sqlc.User{}, errors.New("invalid role")
	}

	var hashedPassword string
	if input.Password != "" {
		hashedPasswordBytes, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
		if err != nil {
			return sqlc.User{}, err
		}
		hashedPassword = string(hashedPasswordBytes)
	} else {
		existingUser, err := s.queries.GetUserByEmail(ctx, input.Email)
		if err != nil {
			return sqlc.User{}, err
		}
		hashedPassword = existingUser.Password
	}

	updatedUser, err := s.queries.UpdateUser(ctx, sqlc.UpdateUserParams{
		ID:       input.ID,
		Email:    input.Email,
		Password: hashedPassword,
		Role:     input.Role,
		Name:     input.Name,
	})
	if err != nil {
		return sqlc.User{}, err
	}

	return sqlc.User{
		ID:        updatedUser.ID,
		Email:     updatedUser.Email,
		Password:  hashedPassword,
		Role:      updatedUser.Role,
		Name:      updatedUser.Name,
		CreatedAt: updatedUser.CreatedAt,
		UpdatedAt: updatedUser.UpdatedAt,
	}, nil
}

func (s *AuthService) GetAllUsers(ctx context.Context) ([]User, error) {
	dbUsers, err := s.queries.ListUsers(ctx)
	if err != nil {
		return nil, err
	}

	users := make([]User, 0, len(dbUsers))
	for _, u := range dbUsers {
		users = append(users, User{
			ID:    u.ID,
			Name:  u.Name,
			Email: u.Email,
			Role:  u.Role,
		})
	}
	return users, nil
}

func (s *AuthService) DeleteUser(ctx context.Context, id int32) error {
	if err := s.queries.DeleteUser(ctx, id); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return errors.New("user not found")
		}
		return err
	}
	return nil
}
