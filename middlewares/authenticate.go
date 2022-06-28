package middlewares

import (
	"blacklist-service/database"
	"encoding/base64"
	"strings"

	"github.com/gofiber/fiber/v2"
)

const (
	CheckCredentials = "SELECT is_revoked FROM auth_clients WHERE id = $1 and secret = $2"
)

func ClientAuthorization(c *fiber.Ctx) error {
	authHeader := c.GetReqHeaders()["Authorization"]
	if authHeader == "" {
		return c.SendStatus(401)
	}

	b, err := base64.StdEncoding.DecodeString(authHeader)
	if err != nil {
		return c.SendStatus(401)
	}

	token := strings.Split(string(b), ":")

	clientId := token[0]
	clientSecret := token[1]

	var isRevoked bool

	if err := database.QueryRow(CheckCredentials, clientId, clientSecret).Scan(&isRevoked); err != nil {
		return err
	}

	if isRevoked {
		return c.SendStatus(403)
	} else {
		return c.Next()
	}
}
