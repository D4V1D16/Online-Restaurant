package Utilities

import (
	"errors"
	"fmt"
	"os"
	"log"
	"time"
	"userAuth/Users/internal/Database/Models"

	"github.com/golang-jwt/jwt/v5"
)


// Working
func GenerateToken(User Models.User) (string, error) {
	fmt.Println(User)
	key := []byte(os.Getenv("SECRET_KEY"))
	if len(key) == 0 {
		return "",errors.New("error loading the key")
		}

	t := jwt.New(jwt.SigningMethodHS256)

	claims := t.Claims.(jwt.MapClaims)

	claims["email"] = User.Email
	claims["id"] = User.IdUser
	//ttl means time to live, and is the time that the token is valid
	ttl := 900 * time.Second
	claims["exp"] = time.Now().Add(ttl).Unix()

	s, err := t.SignedString(key)

	if err != nil {
		return "", err
		}

	return s,nil
}


// Working
func VerifyToken(token string) (string error){

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

	if !t.Valid{
		return errors.New("token is not valid")
	}

	return nil
}

//Extract claim 
func ExtractClaim(tokenStr string, claimKey string) (interface{}, error) {
	// Secret key from environment variable
	key := []byte(os.Getenv("SECRET_KEY"))

	//Key didn't load or exist so return nil and false
	if len(key) == 0 {
		log.Println("Secret key is not set")
		return nil, errors.New("invalid secret key, try to look for it in the environment variable")
	}

	// JET Parsing, trying to centralize this function
	t, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
		// Verify that the signing method used matches the expected method
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return key, nil
	})

	// Error parsing
	if err != nil {
		log.Printf("Error parsing token: %v", err)
		return nil, errors.New("error parsing token")
	}

	// Extract the claims
	if claims, ok := t.Claims.(jwt.MapClaims); ok && t.Valid {
		// Verify if a specific claim exists
		if val, ok := claims[claimKey]; ok {
			return val, nil
		} else {
			log.Printf("Claim '%s' not found in the token", claimKey)
			return nil, errors.New("claim not found in the token")
		}
	}

	// Not JWT Valid
	log.Printf("Invalid JWT Token")
	return nil, errors.New("invalid JWT Token")
}


func TokenExpired(expTime int64) bool{
	currentTime := time.Now().Unix()

	return currentTime > expTime
}