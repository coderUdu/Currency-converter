package controllers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/currencyconverter/database"
	"github.com/currencyconverter/helper"
	"github.com/currencyconverter/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

var userCollection *mongo.Collection = database.OpenCollection(database.Client, "user")

var validate = validator.New()

var Balance models.User.Account == 100

var AppBalance models.User.Account == 0


func Login() gin.HandlerFunc {

	return func(c *gin.Context) {

		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

		var user models.User
		var foundUser models.User

		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}

		err := userCollection.FindOne(ctx, bson.M{"email": user.Email}).Decode(&foundUser)
		defer cancel()

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"email": "Email or password is incorrect"})
			return
		}

		passwordIsValid, msg := VerifyPassword(user.Password, foundUser.Password)
		if passwordIsValid != true {
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
		}

		defer cancel()

		if foundUser.Email == nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "user not found"})
		}

		token, refreshtoken, err := helper.GenerateAllTokens(foundUser.Email, foundUser.User_ID, *foundUser.User_Type)
		helper.UpdateAllTokens(token, refreshtoken, foundUser.User_ID)
		err = userCollection.FindOne(ctx, bson.M{"user_id": foundUser.User_ID}).Decode(&foundUser)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, foundUser)

	}

}

func SignUp() gin.HandlerFunc {

	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

		var user models.User

		if err := c.BindJSON(&user); err != nil {

			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if validationErr := validate.Struct(user); validationErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
			return
		}

		count, err := userCollection.CountDocuments(ctx, bson.M{"email": user.Email})

		defer cancel()

		if err != nil {
			log.Panic(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error occured while checking for the user e:mail"})
		}

		password := HashPassword(*user.Password)
		user.Password = &password

		if count > 0 {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "This email or phone number already exists"})
		}

		user.CreatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		user.UpdatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		user.ID = primitive.NewObjectID()
		user.User_ID = user.ID.Hex()
		token, refreshToken, _ := helper.GenerateAllTokens(*user.Email, *&user.User_ID, *user.UserType)
		user.Token = &token
		user.RefreshToken = &refreshToken

		resultInsertionNumber, insertErr := userCollection.InsertOne(ctx, user)
		if insertErr != nil {
			msg := fmt.Sprintf("user item was not created")

			c.JSON(http.StatusBadRequest, gin.H{"error": msg})
			return
		}

		defer cancel()

		c.JSON(http.StatusOK, resultInsertionNumber)
	}

}

func GetBalance() {

}

func HashPassword(password string) string {

	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)

	if err != nil {
		log.Panic(err)
	}
	return string(bytes)
}

func VerifyPassword(userPassword, providedPassword string) (bool, string) {

	err := bcrypt.CompareHashAndPassword([]byte(userPassword), []byte(providedPassword))

	check := true

	msg := " "

	if err != nil {
		msg = fmt.Sprint("email of password is incorrect")
		check = false
	}
	return check, msg
}

func CreditApp() {

}

func Conversion() gin.HandlerFunc {
	return func(c *gin.Context) {

		var response struct{
			From  string
			To    string
			Amount float64
		}
    
		Models.Currency.Currencies := GetAllCurrency()


		var fromValue float64
		fromValue = getCurrencyValue(from, data)

		var toValue float64
		toValue = getCurrencyValue(to, data)

		var convertedAmount float64

		floatAmount, _ := strconv.ParseFloat(amount, 8)
		convertedAmount = (floatAmount * toValue) / fromValue
	}

}

func GetCurrencyValue(moneyCurrency string, data models.Currency) float64 {

	var value float64

	if moneyCurrency == "USD" {
		value = float64(data.Currencies.NGN)
	} else if moneyCurrency == "USD" {
		value = float64(data.Currencies.USD)
	}
	return value
}

func GetAllCurrency() models.Currency.Currencies{

	if models.Currency.Currencies == NGN {
		return NGN
	}else if models.Currency.Currencies == USD {
		return USD
	}
      return models.Currency.Currencies
}
