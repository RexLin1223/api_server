package authentication

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"strings"
	"time"

	"api/database"
	"api/server/protocol"

	"github.com/robbert229/jwt"
)

// Authenticator will generate and validate token from client
type Authenticator struct {
	dbInfo      database.ConnectionInfo
	expiredTime int64
}

// NewAuthenticator function will create an Authenticator
func NewAuthenticator() *Authenticator {
	return &Authenticator{
		dbInfo: database.ConnectionInfo{
			Domain:     "127.0.0.1",
			Port:       "3306",
			Username:   "root",
			Password:   "",
			TargetName: "",
		},
		expiredTime: 600, // Defualt value
	}
}

// SetExpiredTime will rewrite expired time into authenticator
func (auth *Authenticator) SetExpiredTime(expiredTime int64) *Authenticator {
	auth.expiredTime = expiredTime
	return auth
}

// SetHost will rewrite new domain and port into authenticator
func (auth *Authenticator) SetHost(domain string, port string) *Authenticator {
	auth.dbInfo.Domain = domain
	auth.dbInfo.Port = port
	return auth
}

// SetAuthentication will rewrite {username} and {password} into authenticator
func (auth *Authenticator) SetAuthentication(username string, password string) *Authenticator {
	auth.dbInfo.Username = username
	auth.dbInfo.Password = password
	return auth
}

// SetDatabase will rewrite {databaseName} into authenticator
func (auth *Authenticator) SetDatabase(databaseName string) *Authenticator {
	auth.dbInfo.TargetName = databaseName
	return auth
}

// GenerateToken function will generate standard JWT
func (auth *Authenticator) GenerateToken(url string, username string, password string) (string, error) {
	// Verify user permission here
	// Load database
	connector := database.NewMySQLLConnector(&auth.dbInfo)
	err := connector.Open()
	if err != nil {
		return "", err
	}

	sqlCommand := fmt.Sprintf(
		"WHERE (`phone_number` ='%s') AND (`password`='%s')", username, password)
	result, err := connector.Read("user_member", sqlCommand)
	if err != nil {
		return "", err
	}
	if len(result) == 0 {
		return "", errors.New("Empty query")
	}

	// Stuff user claims here
	params := make(map[string]interface{})
	params["role"] = result["role"].(string)
	params["exp"] = time.Now().Unix() + auth.expiredTime

	jwt, err := generateJWT(params)
	if err != nil {
		return "", err
	}

	return generateReponseJSON(jwt), nil
}

// VerifyToken function will valadiate token
func (auth *Authenticator) VerifyToken(authorization string) error {
	// Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImFkbWluIiwiaWF0IjoxNTM0OTI1MTQwLCJleHAiOjE1MzUwMTE1NDB9.MIcWFBzAr5WVhbaSa1kd1_hmEZsepo8fXqotqvAerKI
	// Skip 'Bearer '
	return validateJWT(authorization[7:])
}

func generateReponseJSON(token string) string {
	resp := protocol.TokenResponse{
		Result:  "Successful",
		Message: "JWT",
		Token:   token,
	}

	result, err := json.Marshal(resp)
	if err != nil {
		return ""
	}

	return string(result)
}

var secret = "DagogoSecret"

func generateJWT(params map[string]interface{}) (string, error) {
	alg := jwt.HmacSha256(secret)

	claims := jwt.NewClaim()
	// User define payload

	for key, value := range params {
		claims.Set(key, value)
	}

	token, err := alg.Encode(claims)
	if err != nil {
		log.Fatal(err)
		return "", errors.New("Encode failed")
	}

	return token, nil
}

func validateJWT(token string) error {
	alg := jwt.HmacSha256(secret)
	claims, err := alg.DecodeAndValidate(token)
	if err != nil {
		return err
	}

	role, err := claims.Get("role")
	if err != nil {
		return err
	}

	roleString, ok := role.(string)
	if !ok {
		return errors.New("Invalid token with incorrect payload")
	}

	if (0 == strings.Compare(roleString, "admin")) ||
		(0 == strings.Compare(roleString, "member")) ||
		(0 == strings.Compare(roleString, "assistant")) {
		return nil
	}

	return errors.New("Unkown error in validate JWT")
}
