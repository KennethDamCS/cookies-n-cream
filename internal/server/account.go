package server

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func createJWTToken(account string) string {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["account_email"] = account
	// Only valid for 15 mins
	claims["expire"] = time.Now().Add(time.Minute * 15).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := at.SignedString([]byte(os.Getenv("JWT_TOKEN")))
	if err != nil {
		return err.Error()
	}
	return token
}

func signUp(c *gin.Context) {

	account := new(Account)
	if err := c.Bind(account); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	// Attempt to add account
	if err := AddAccount(account); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
	}

	// Lets load env for any special tokens
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file")
	}

	// Add to list of accounts
	Accounts = append(Accounts, account)

	// Confirm signup
	c.JSON(http.StatusOK, gin.H{
		"message": "Signed up succesfully",
		"jwt":     createJWTToken(account.Email),
	})
}

func signIn(c *gin.Context) {

	account := new(Account)
	if err := c.Bind(account); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	account, err := Authenticate(account.Email, account.Password)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Sign in failed."})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Signed in successfully.",
		"jwt":     createJWTToken(account.Email),
	})
}
