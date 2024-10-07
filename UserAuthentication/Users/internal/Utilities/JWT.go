package Utilities

import (
	"errors"
	"fmt"
	"log"
	"os"
	"time"
	"userAuth/Users/internal/Database/Models"
	"userAuth/Users/internal/Database"

	"github.com/golang-jwt/jwt/v5"
)

// TTL Constant
const tokenTTL = 900 * time.Second

// Generate a JWT token
func GenerateToken(User Models.User, typeToken string) (string, error) {
	key := []byte(os.Getenv("SECRET_KEY"))
	if len(key) == 0 {
		return "", errors.New("error loading the key")
	}
	
	t := jwt.New(jwt.SigningMethodHS256)
	
	claims := t.Claims.(jwt.MapClaims)
	claims["email"] = User.Email
	claims["id"] = User.IdUser
	
	if typeToken == "access" {
		claims["exp"] = time.Now().Add(tokenTTL).Unix() 
		claims["token_type"] = "access"
	} else if typeToken == "refresh" {
		claims["exp"] = time.Now().Add(24 * time.Hour).Unix() 
		claims["token_type"] = "refresh"
	} else {
		return "", errors.New("invalid token type")
	}
	
	signedToken, err := t.SignedString(key)
	if err != nil {
		return "", err
	}
	
	return signedToken, nil
}

// Verify a JWT Token
func VerifyToken(token string, expectedType string) error {
	key := []byte(os.Getenv("SECRET_KEY"))

	if len(key) == 0 {
		return errors.New("error loading the key")
	}

	t, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return key, nil
	})

	if err != nil {
		return err
	}

	if !t.Valid {
		return errors.New("token is not valid")
	}

	claims := t.Claims.(jwt.MapClaims)
	if tokenType, ok := claims["token_type"]; ok {
		if tokenType != expectedType {
			return fmt.Errorf("invalid token type: expected %s, got %s", expectedType, tokenType)
		}
	} else {
		return errors.New("token_type claim not found in the token")
	}

	return nil
}

// Extract a specific claim from the token
func ExtractClaim(tokenStr string, claimKey string) (interface{}, error) {
	key := []byte(os.Getenv("SECRET_KEY"))

	if len(key) == 0 {
		log.Println("Secret key is not set")
		return nil, errors.New("invalid secret key, try to look for it in the environment variable")
	}

	t, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return key, nil
	})

	if err != nil {
		log.Printf("Error parsing token: %v", err)
		return nil, errors.New("error parsing token")
	}

	if claims, ok := t.Claims.(jwt.MapClaims); ok && t.Valid {
		if val, ok := claims[claimKey]; ok {
			return val, nil
		} else {
			log.Printf("Claim '%s' not found in the token", claimKey)
			return nil, errors.New("claim not found in the token")
		}
	}

	log.Printf("Invalid JWT Token")
	return nil, errors.New("invalid JWT Token")
}

// Verify if the token has expired
func TokenExpired(expTime int64) bool {
	return time.Now().After(time.Unix(expTime, 0))
}

func VerifyRefreshToken(refreshToken string) (string, error) {
	err := VerifyToken(refreshToken, "refresh")
	if err != nil {
		return "", fmt.Errorf("error: %v", err)
	}
	
	emailClaim, err := ExtractClaim(refreshToken, "email")
	if err != nil {
		return "", fmt.Errorf("error extracting claim: %v", err)
	}
	
	username, ok := emailClaim.(string)
	if !ok {
		return "", fmt.Errorf("invalid username claim type")
	}

	accessToken, err := GenerateToken(Models.User{Email: username}, "access")
	if err != nil {
		return "", fmt.Errorf("error generating access token: %v", err)
	}

	return accessToken, nil
}


func ValidateUserExistence(userID string) error {

		var user Models.User
		if err := Database.DB.First(&user, "id_user = ?", userID).Error; err != nil {
			return errors.New("user not found")
		}
		return nil
	
}


func IsInvalidated(token string) bool {

	var tokenModel Models.Token
	if err := Database.DB.First(&tokenModel, "token = ?", token).Error; err != nil {
		return false
	}
	return true
}