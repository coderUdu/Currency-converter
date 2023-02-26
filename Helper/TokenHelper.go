package helper

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/currencyconverter/database"

	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type SignedDetails struct {
	Email    string
	UID      string
	UserType string
	jwt.StandardClaims
}

var userCollection *mongo.Collection = database.OpenCollection(database.Client, "User")

var SECRET_KEY string = os.Getenv("secretKey")

func GenerateAllTokens(email, firstname, lastname, uid, userType string) (signedToken, signedRefreshToken string, err error) {

	claims := &SignedDetails{
		Email:    email,
		UID:      uid,
		UserType: userType,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(24)).Unix(),
		},
	}

	refreshClaims := &SignedDetails{

		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(168)).Unix(),
		},
	}
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(SECRET_KEY))

	refreshToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims).SignedString([]byte(SECRET_KEY))

	if err != nil {
		log.Panic(err)
		return
	}

	return token, refreshToken, err

}

func UpdateAllTokens(signedToken, signedRefreshToken, userID string) {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

	var updatedObj primitive.D

	updatedObj = append(updatedObj, bson.E{"token", signedToken})
	updatedObj = append(updatedObj, bson.E{"refreshToken", signedRefreshToken})

	updatedAt, _ := time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	updatedObj = append(updatedObj, bson.E{"updatedAt", updatedAt})

	upsert := true
	filter := bson.M{"userID": UserID}
	opt := options.UpdateOptions{
		Upsert: &upsert,
	}

	_, err := userCollection.UpdateOne(ctx,
		filter,
		bson.D{
			{"$set", updatedObj},
		},
		&opt)

	defer cancel()
	if err != nil {
		log.Panic(err)
		return
	}
	return

}

func ValidateToken(signedToken string) (claims *SignedDetails, msg string) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&SignedDetails{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(SECRET_KEY), nil
		},
	)
	if err != nil {
		msg = err.Error()
		return
	}

	claims, ok := token.Claims.(*SignedDetails)
	if !ok {
		msg = fmt.Sprintf("the token is valid")
		msg = err.Error()
		return
	}

	if claims.ExpiresAt < time.Now().Local().Unix() {
		msg = fmt.Sprint("the token is invalid")
		msg = err.Error()
		return
	}
	return claims, msg
}
