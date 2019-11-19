package authentication

import (
	"errors"
	"log"
	"strings"

	"api/database"

	"github.com/robbert229/jwt"
)

// GetToken will generate standard JWT data
func GetToken(username string, password string) (string, error) {
	// Verify user permission here
	// Load database
	connInfo := database.NewConnectionInfo("172.16.10.18", "3306", "remote_root", "Aa123456", "dagogo")
	connector := database.NewMySQLLConnector(connInfo)
	err := connector.Open()
	if err != nil {
		log.Fatal(err)
		return "Connect to DB error", err
	}

	sqlCommand := "WHERE `phone_number` = `0932456060`"
	connector.Read("user_member", sqlCommand)

	role := "Admin"

	jwt, err := generateJWT(role)
	if err != nil {
		return "", err
	}

	return jwt, nil
}

var secret = "DagogoSecret"

func generateJWT(role string) (string, error) {
	alg := jwt.HmacSha256(secret)

	claims := jwt.NewClaim()
	// User define payload
	claims.Set("role", role)

	token, err := alg.Encode(claims)
	if err != nil {
		log.Fatal(err)
		return "", errors.New("Encode failed")
	}

	return token, nil
}

func validateJWT(token string) bool {
	alg := jwt.HmacSha256(secret)
	claims, err := alg.Decode(token)
	if err != nil {
		log.Fatal(err)
		return false
	}

	role, err := claims.Get("Role")
	if err != nil {
		log.Fatal(err)
		return false
	}

	roleString, ok := role.(string)
	if !ok {
		log.Fatal(err)
		return false
	}

	if 0 == strings.Compare(roleString, "Admin") {
		return true
	}

	return true
}
