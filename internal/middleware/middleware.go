package middleware

import (
	"errors"

	"github.com/Markuysa/courceWorkBackendDev/internal/auth/cache"
	"github.com/Markuysa/courceWorkBackendDev/internal/models"
	"github.com/Markuysa/courceWorkBackendDev/pkg/constants"
	"github.com/Markuysa/courceWorkBackendDev/pkg/lists"
	"github.com/gofiber/fiber/v2"
	"github.com/redis/go-redis/v9"
)

type Mw struct {
	sessionCache cache.Cache
}

func New(sessionCache cache.Cache) *Mw {
	return &Mw{
		sessionCache: sessionCache,
	}
}

func (mw *Mw) AdminAuth(c *fiber.Ctx) error {

	session := c.Cookies(constants.SessionKey)

	cachedSession, err := mw.sessionCache.GetSession(c.Context(), models.GetSessionRequest{
		SessionKey: session,
	})
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return c.SendStatus(fiber.StatusUnauthorized)
		}

		return c.SendStatus(fiber.StatusInternalServerError)
	}

	if cachedSession.Role != lists.RoleAdmin {
		return c.SendStatus(fiber.StatusForbidden)
	}

	c.Locals(constants.UserIDKey, cachedSession.UserID)

	return c.Next()
}

func (mw *Mw) ClientAuth(c *fiber.Ctx) error {
	session := c.Cookies(constants.SessionKey)

	cachedSession, err := mw.sessionCache.GetSession(c.Context(), models.GetSessionRequest{
		SessionKey: session,
	})
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return c.SendStatus(fiber.StatusUnauthorized)
		}

		return c.SendStatus(fiber.StatusInternalServerError)
	}

	if cachedSession.Role != lists.RoleUser {
		return c.SendStatus(fiber.StatusForbidden)
	}

	c.Locals(constants.UserIDKey, cachedSession.UserID)

	return c.Next()
}

func (mw *Mw) DefaultAuth(c *fiber.Ctx) error {

	return c.Next()
}
