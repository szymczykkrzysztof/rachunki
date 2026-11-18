package httpapi

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// Pinger sprawdza, czy baza odpowiada. Spełnia go *pgxpool.Pool.
type Pinger interface {
	Ping(ctx context.Context) error
}

const readyTimeout = 2 * time.Second

func NewRouter(db Pinger) *gin.Engine {
	r := gin.Default()
	r.HandleMethodNotAllowed = true
	r.GET("/health", health)
	r.GET("/ready", ready(db))
	return r
}

func health(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}

func ready(db Pinger) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(c.Request.Context(), readyTimeout)
		defer cancel()

		if err := db.Ping(ctx); err != nil {
			_ = c.Error(err)
			c.JSON(http.StatusServiceUnavailable, gin.H{"status": "unavailable"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"status": "ready"})
	}
}
