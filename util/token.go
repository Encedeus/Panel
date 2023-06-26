package util

import (
	"errors"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"golang.org/x/exp/slices"
	"panel/config"
	"panel/dto"
	"strings"
	"time"
)

// GenerateAccessToken generates an access token containing the uuid of a user that expires in 15 minutes
func GenerateAccessToken(userData dto.AccessTokenDTO) (string, error) {

	userData.ExpiresAt = time.Now().Add(15 * time.Minute).Unix()

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, userData)
	accessTokenString, err := accessToken.SignedString([]byte(config.Config.Auth.JWTSecretAccess))
	if err != nil {
		return "", err
	}

	return accessTokenString, nil
}

// GenerateRefreshToken generates a refresh token containing the uuid of a user that expires in a week
func GenerateRefreshToken(userData dto.RefreshTokenDTO) (string, error) {
	// generate a token containing the user's uuid
	userData.ExpiresAt = time.Now().Add(168 * time.Hour).Unix()

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, userData)
	accessTokenString, err := accessToken.SignedString([]byte(config.Config.Auth.JWTSecretRefresh))
	if err != nil {
		return "", err
	}

	return accessTokenString, nil
}

// GetTokenPair returns an access and a refresh token
func GetTokenPair(userData dto.AccessTokenDTO) (string, string, error) {
	accessToken, err1 := GenerateAccessToken(userData)
	refreshToken, err2 := GenerateRefreshToken(dto.RefreshTokenDTO{UserId: userData.UserId})

	if err1 != nil {
		log.Errorf("error generating access token %v", err1)
		return "", "", err1
	}
	if err2 != nil {
		log.Errorf("error generating refresh token %v", err2)
		return "", "", err2
	}

	return accessToken, refreshToken, nil
}

// GetTokenFromHeader extracts the token from the auth header
func GetTokenFromHeader(ctx echo.Context) string {
	// removes "Bearer" in front of the token and returns the token
	return strings.Split(ctx.Request().Header.Get("Authorization"), " ")[1]
}

func ValidateAccessJWT(tokenString string) (bool, dto.AccessTokenDTO, error) {
	// parse the JWT and check the signing method

	claims := dto.AccessTokenDTO{}

	_, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
		if token.Method != jwt.SigningMethodHS256 {
			return nil, errors.New("unexpected jwt signing method")
		}
		return []byte(config.Config.Auth.JWTSecretAccess), nil
	})

	if err != nil {
		return false, claims, err
	}

	return true, claims, err
}
func ValidateRefreshJWT(tokenString string) (bool, dto.RefreshTokenDTO, error) {
	// parse the JWT and check the signing method

	claims := dto.RefreshTokenDTO{}

	_, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
		if token.Method != jwt.SigningMethodHS256 {
			return nil, errors.New("unexpected jwt signing method")
		}
		return []byte(config.Config.Auth.JWTSecretRefresh), nil
	})

	if err != nil {
		return false, claims, err
	}

	return true, claims, err
}

// DoesTokenContainPermission checks whether a JWT contains a permission
func DoesTokenContainPermission(permission string, jwt string) bool {
	// verifies JWT is valid
	isValid, token, err := ValidateAccessJWT(jwt)
	if err != nil || !isValid {
		return false
	}

	// checks if JWT contains the permission or was granted all permissions via '*'
	if !slices.Contains(token.Permissions, permission) && !slices.Contains(token.Permissions, "*") {
		return false
	}
	return true
}