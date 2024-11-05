package internals

import (
	"sync"

	"github.com/quankori/go-manhattan-distance/server/internals/cinema/services"
)

var (
	once      sync.Once
	container *Container
)

type Container struct {
	CinemaService services.CinemaService
}

// GetContainer returns a singleton instance of the Container
func GetContainer() *Container {
	once.Do(func() {
		// Initialize dependencies here
		cinemaService := services.NewUserService(5, 5, 2)
		container = &Container{
			CinemaService: cinemaService,
		}
	})
	return container
}
