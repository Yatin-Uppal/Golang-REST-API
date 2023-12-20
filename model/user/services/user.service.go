package userServices

import (
	"REST-API/model/user/dto"
	userSchema "REST-API/model/user/schema"
	"REST-API/providers"
	"context"
	"errors"
	"os"
	"time"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// function to get the List of users
func GetUserList() ([]*userSchema.UserSchema, error) {
	var userList []*userSchema.UserSchema
	// creating database connectivity
	client, err := providers.GetMongoDBClient()
	if err != nil {
		return userList, err
	}
	defer client.Disconnect(context.Background())

	collection := client.Database((os.Getenv(("DATABASE_NAME")))).Collection(userSchema.USER_SCHEMA_NAME)
	cursor, err := collection.Find(context.TODO(), bson.D{})

	if err != nil {
		return userList, err
	}

	for cursor.Next(context.TODO()) {
		var user userSchema.UserSchema
		err := cursor.Decode(&user)
		if err != nil {
			return nil, err
		}
		userList = append(userList, &user)
	}
	return userList, nil
}

// function to fetch the details of a single user
func GetUserDetails(userId string) (*userSchema.UserSchema, error) {
	var userDetails *userSchema.UserSchema
	client, err := providers.GetMongoDBClient()
	if err != nil {
		return nil, err
	}
	defer client.Disconnect(context.Background())
	whereCondtion := bson.D{bson.E{Key: "userId", Value: userId}}
	collection := client.Database(os.Getenv("DATABASE_NAME")).Collection(userSchema.USER_SCHEMA_NAME)
	collection.FindOne(context.TODO(), whereCondtion).Decode(&userDetails)
	return userDetails, err
}

// function to create  a new user
func CreateUser(user dto.CreateUserDto) (*userSchema.UserSchema, error) {
	client, err := providers.GetMongoDBClient()
	if err != nil {
		return nil, err
	}
	defer client.Disconnect(context.Background())

	// Create a new user schema
	createdUser := &userSchema.UserSchema{
		Id:           primitive.NewObjectID(),
		UserId:       uuid.New().String(),
		FirstName:    user.FirstName,
		LastName:     user.LastName,
		Email:        user.Email,
		MobileNumber: user.MobileNumber,
		Password:     user.Password,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	collection := client.Database(os.Getenv("DATABASE_NAME")).Collection(userSchema.USER_SCHEMA_NAME)

	_, err = collection.InsertOne(context.TODO(), createdUser) // Assuming userService.context is valid

	if err != nil {
		return nil, err
	}
	return createdUser, nil
}

// function to update user
func UpdateUser(userId string, user *dto.UpdateUserDto) (*userSchema.UserSchema, error) {
	client, err := providers.GetMongoDBClient()
	if err != nil {
		return nil, err
	}
	defer client.Disconnect(context.Background())
	checkExistingUser, err := GetUserDetails(userId) // Added error handling here
	if err != nil {
		return nil, err
	}
	if checkExistingUser == nil {
		return nil, errors.New("user doesn't exist")
	}
	// Create a new updatedUser instance and populate it with the updated values
	updatedUser := &userSchema.UserSchema{
		Id:        checkExistingUser.Id,
		UserId:    checkExistingUser.UserId,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		UpdatedAt: time.Now(),
		CreatedAt: checkExistingUser.CreatedAt,
	}
	whereCondition := bson.D{bson.E{Key: "userId", Value: userId}}
	collection := client.Database(os.Getenv("DATABASE_NAME")).Collection(userSchema.USER_SCHEMA_NAME)
	_, updationError := collection.UpdateOne(context.TODO(), whereCondition, bson.D{bson.E{Key: "$set", Value: updatedUser}})
	if updationError != nil {
		return nil, updationError
	}
	return updatedUser, nil
}

// function to Delete user
func DeleteUser(userId string) error {
	client, err := providers.GetMongoDBClient()
	if err != nil {
		return err
	}
	defer client.Disconnect(context.Background())

	checkExistingUser, fetchingError := GetUserDetails(userId) // Added error handling here
	if fetchingError != nil {
		return fetchingError
	}
	if checkExistingUser == nil {
		return errors.New("user doesn't exist")
	}
	whereCondition := bson.D{bson.E{Key: "userId", Value: userId}}
	collection := client.Database(os.Getenv("DATABASE_NAME")).Collection(userSchema.USER_SCHEMA_NAME)
	_, deletionErr := collection.DeleteOne(context.TODO(), whereCondition)
	return deletionErr
}
