package verification

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"

	"api/database"
	"api/server/config"
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
	serverConfig := config.Load()
	return &Authenticator{
		dbInfo: database.ConnectionInfo{
			Domain:     serverConfig.DBInfo.Domain,
			Port:       serverConfig.DBInfo.Port,
			Username:   serverConfig.DBInfo.User,
			Password:   serverConfig.DBInfo.Password,
			TargetName: serverConfig.DBInfo.TargetName,
		},
		expiredTime: 600, // Defualt value
	}
}

// SetExpiredTime will rewrite expired time into authenticator
func (auth *Authenticator) SetExpiredTime(expiredTime int64) *Authenticator {
	auth.expiredTime = expiredTime
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
	} else if len(result) == 0 {
		return "", errors.New("Username or Password incorrectly")
	} else if len(result) > 1 {
		return "", errors.New("Multi-role in database")
	}

	// Stuff user claims here
	params := make(map[string]interface{})
	params["role"] = result[0]["role"].(string)
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

var secret = `D!a@g#o$g%o^S(e)c_r+e!t@`

func generateReponseJSON(token string) string {
	// Formating token to JSON
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

func generateJWT(params map[string]interface{}) (string, error) {
	alg := jwt.HmacSha256(secret)

	// User defined payload
	claims := jwt.NewClaim()
	for key, value := range params {
		claims.Set(key, value)
	}

	token, err := alg.Encode(claims)
	if err != nil {
		return "", err
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
