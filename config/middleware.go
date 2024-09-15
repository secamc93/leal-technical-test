package config

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

// TokenManager es la estructura que manejará la generación y validación de tokens y el middleware.
type TokenManager struct {
	jwtKey []byte
	env    *Env
}

// Claims define la estructura de los datos que queremos almacenar en el token JWT.
type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

// NewTokenManager es el constructor que carga la clave JWT desde las variables de entorno.
func NewTokenManager() *TokenManager {
	jwtKey := NewGetEnv().JwtKey
	if jwtKey == "" {
		panic("JWT_KEY not set in environment variables")
	}
	return &TokenManager{
		jwtKey: []byte(jwtKey),
	}
}

// GenerateToken genera un nuevo token JWT para el usuario proporcionado.
func (tm *TokenManager) GenerateToken(username string) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour) // Token válido por 24 horas
	claims := &Claims{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	// Crear el token con el método de firma HS256 y los claims definidos.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(tm.jwtKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// ValidateToken valida el token JWT proporcionado.
func (tm *TokenManager) ValidateToken(tokenString string) (*Claims, error) {
	claims := &Claims{}

	// Analizar el token y verificar si es válido.
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return tm.jwtKey, nil
	})
	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	return claims, nil
}

// AuthMiddleware es el middleware que valida el token JWT.
func (tm *TokenManager) AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Obtener el token de la cabecera Authorization.
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header required"})
			c.Abort()
			return
		}

		var tokenString string
		// El token puede tener el formato "Bearer {token}" o solo "{token}"
		if strings.HasPrefix(authHeader, "Bearer ") {
			tokenString = strings.TrimPrefix(authHeader, "Bearer ")
		} else {
			tokenString = authHeader
		}

		// Validar el token
		claims, err := tm.ValidateToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			c.Abort()
			return
		}

		// Establecer el usuario en el contexto para acceder a él en los controladores
		c.Set("username", claims.Username)
		c.Next()
	}
}
