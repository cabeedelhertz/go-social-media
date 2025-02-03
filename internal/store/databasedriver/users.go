package databasedriver

import (
	"context"
	"errors"
	"fmt"
	"social/internal/models"
	"social/pkg/common/logging"
	socialv1 "social/proto/gen/social/v1"
	"strings"
	"time"

	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

func (s *Store) CreateOrUpdateUser(ctx context.Context, usr *socialv1.User) (*socialv1.User, error) {
	if usr.ExternalAuthId == "" {
		return nil, fmt.Errorf("external auth id is required")
	}
	var dbUser *models.User
	err := s.db.WithContext(ctx).Where("external_auth_id = ?", usr.ExternalAuthId).First(&dbUser).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		dbUser = models.UserFromProto(usr)
		dbUser.ID = uuid.Must(uuid.NewV4())
		if dbUser.Username == "" {
			dbUser.Username = fmt.Sprintf("%s-%v", strings.Split(usr.Email, "@")[0], time.Now().Unix())
		}
		dbUser.Username = strings.ToLower(dbUser.Username)
		if err := s.db.WithContext(ctx).Create(&dbUser).Error; err != nil {
			if IsDuplicateKeyError(err) {
				logging.FromContext(ctx).Error("duplicate key error, retrying with new username")
				dbUser.Username = fmt.Sprintf("%s-%v", dbUser.Username, time.Now().Unix())
				dbUser.Username = strings.ToLower(dbUser.Username)
				if err := s.db.WithContext(ctx).Create(&dbUser).Error; err != nil {
					return nil, err
				}
			}
			return nil, err
		}
	}

	dbUser.LastLoginTime = time.Now()
	if err := s.db.WithContext(ctx).Save(&dbUser).Error; err != nil {
		return nil, err
	}
	return dbUser.ToProto(), nil
}

func (s *Store) GetUser(ctx context.Context, id string) (*socialv1.User, error) {
	var user models.User
	if err := s.db.WithContext(ctx).Where("id = ?", id).First(&user).Error; err != nil {
		return nil, err
	}
	return user.ToProto(), nil
}

func buildSearchUsersQuery(db *gorm.DB, searchTerm string) (*gorm.DB, error) {
	if searchTerm == "" {
		return nil, fmt.Errorf("search term is required")
	}
	splitSearchTerm := strings.Split(searchTerm, " ")
	firstNameQuery := splitSearchTerm[0]
	lastNameQuery := ""
	if len(splitSearchTerm) > 1 {
		lastNameQuery = splitSearchTerm[1]
	}
	query := db.Where("username LIKE ?", "%"+strings.ToLower(firstNameQuery)+"%").Or("email LIKE ?", "%"+strings.ToLower(firstNameQuery)+"%")
	query = query.Or("first_name ILIKE ?", "%"+strings.ToLower(firstNameQuery)+"%")
	if lastNameQuery != "" {
		query = query.Or("last_name ILIKE ?", "%"+strings.ToLower(lastNameQuery)+"%")
	}
	return query, nil
}

func (s *Store) ListUsers(ctx context.Context, req *socialv1.ListUsersRequest) ([]*socialv1.User, error) {
	if req.GetSearchTerm() == "" {
		return nil, fmt.Errorf("search term is required")
	}

	var users []models.User
	query := s.db.WithContext(ctx)
	query, err := buildSearchUsersQuery(query, req.GetSearchTerm())
	if err != nil {
		return nil, err
	}
	if err := query.Find(&users).Limit(15).Error; err != nil {
		return nil, err
	}
	var protoUsers []*socialv1.User
	for _, user := range users {
		protoUsers = append(protoUsers, user.ToProto())
	}
	return protoUsers, nil
}
