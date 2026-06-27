package server

import (
	"context"
	"fmt"
	"time"

	"github.com/critical-dummy/flim/internal/auth"
	"github.com/critical-dummy/flim/internal/db"
	"github.com/critical-dummy/flim/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

type MessengerServer struct {
	db         *db.MongoDB
	jwtManager *auth.JWTManager
}

func NewMessengerServer(db *db.MongoDB, jwtManager *auth.JWTManager) *MessengerServer {
	return &MessengerServer{
		db:         db,
		jwtManager: jwtManager,
	}
}

// Register creates a new user
func (s *MessengerServer) Register(ctx context.Context, username, email, password string) (*models.User, string, error) {
	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, "", err
	}

	// Create user
	user := &models.User{
		Username:  username,
		Email:     email,
		Password:  string(hashedPassword),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	// Insert into DB
	usersCollection := s.db.GetUsersCollection()
	result, err := usersCollection.InsertOne(ctx, user)
	if err != nil {
		return nil, "", err
	}

	user.ID = result.InsertedID.(primitive.ObjectID)

	// Generate token
	token, err := s.jwtManager.GenerateToken(user.ID.Hex())
	if err != nil {
		return nil, "", err
	}

	return user, token, nil
}

// Login authenticates a user
func (s *MessengerServer) Login(ctx context.Context, username, password string) (*models.User, string, error) {
	usersCollection := s.db.GetUsersCollection()

	var user *models.User
	err := usersCollection.FindOne(ctx, bson.M{"username": username}).Decode(&user)
	if err != nil {
		return nil, "", fmt.Errorf("user not found")
	}

	// Verify password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, "", fmt.Errorf("invalid password")
	}

	// Generate token
	token, err := s.jwtManager.GenerateToken(user.ID.Hex())
	if err != nil {
		return nil, "", err
	}

	return user, token, nil
}

// SendMessage sends a message from one user to another
func (s *MessengerServer) SendMessage(ctx context.Context, fromUserID, toUserID, content string) (*models.Message, error) {
	message := &models.Message{
		FromUserID: fromUserID,
		ToUserID:   toUserID,
		Content:    content,
		Timestamp:  time.Now(),
		IsRead:     false,
	}

	messagesCollection := s.db.GetMessagesCollection()
	result, err := messagesCollection.InsertOne(ctx, message)
	if err != nil {
		return nil, err
	}

	message.ID = result.InsertedID.(primitive.ObjectID)
	return message, nil
}

// GetConversation retrieves messages between two users
func (s *MessengerServer) GetConversation(ctx context.Context, userID1, userID2 string, limit int64) ([]*models.Message, error) {
	messagesCollection := s.db.GetMessagesCollection()

	filter := bson.M{
		"$or": []bson.M{
			{
				"from_user_id": userID1,
				"to_user_id":   userID2,
			},
			{
				"from_user_id": userID2,
				"to_user_id":   userID1,
			},
		},
	}

	cursor, err := messagesCollection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var messages []*models.Message
	if err = cursor.All(ctx, &messages); err != nil {
		return nil, err
	}

	return messages, nil
}
