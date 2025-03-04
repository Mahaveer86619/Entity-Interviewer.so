package middleware

import (
	"context"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	types "github.com/Mahaveer86619/Entity-Interviewer.so/go-server/internal/types"
	"github.com/golang-jwt/jwt/v5"
)

type contextKey string

const userContextKey = contextKey("user")

type Claims struct {
	Email string `json:"email"`
	jwt.MapClaims
}

var JwtKey = []byte(os.Getenv("JWT_SECRET"))

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			failureResponse := types.Failure{}
			failureResponse.SetStatusCode(http.StatusUnauthorized)
			failureResponse.SetMessage("Authorization header is required")
			failureResponse.JSON(w)
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		claims := &Claims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return JwtKey, nil
		})

		if err != nil || !token.Valid {
			failureResponse := types.Failure{}
			failureResponse.SetStatusCode(http.StatusUnauthorized)
			failureResponse.SetMessage("Invalid token")
			failureResponse.JSON(w)
			return
		}

		// Token is valid, proceed to set the context
		ctx := context.WithValue(r.Context(), userContextKey, claims.Email)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func UserFromContext(ctx context.Context) (string, bool) {
	email, ok := ctx.Value(userContextKey).(string)
	return email, ok
}

func LoggingMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Get the remote address
		remoteAddr := r.RemoteAddr

		// Extract the IP address from remoteAddr (host:port format)
		ipAddress := remoteAddr[:strings.LastIndex(remoteAddr, ":")]

		// Log the request with IP address
		logString := `
		Request received from IP address: %s
		Method: %s & Path: %s
		User-Agent: %s
		`
		log.Printf(logString, ipAddress, r.Method, r.URL.Path, r.UserAgent())

		// Call the next handler in the chain
		next(w, r)
	}
}

func GenerateToken(email string) (string, error) {
	expirationTime := time.Now().Add(25 * time.Hour) // 1 day + 1 hour
	claims := &Claims{
		Email: email,
		MapClaims: jwt.MapClaims{
			"email": email,
			"exp":   expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(JwtKey)
}

func GenerateRefreshToken(email string) (string, error) {
	expirationTime := time.Now().Add(721 * time.Hour) // 30 days + 1 hour
	claims := &Claims{
		Email: email,
		MapClaims: jwt.MapClaims{
			"email": email,
			"exp":   expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(JwtKey)
}
