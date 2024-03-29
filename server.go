package main

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-contrib/cors"

	"github.com/gin-gonic/gin"

	"github.com/antch57/jam-statz/graph"

	"github.com/antch57/jam-statz/pkg/albums"
	"github.com/antch57/jam-statz/pkg/bands"
	"github.com/antch57/jam-statz/pkg/db"
	"github.com/antch57/jam-statz/pkg/performances"
	"github.com/antch57/jam-statz/pkg/songs"
	"github.com/antch57/jam-statz/pkg/venues"
)

// Defining the Graphql handler
func graphqlHandler(resolver *graph.Resolver) gin.HandlerFunc {
	// NewExecutableSchema and Config are in the generated.go file
	// Resolver is in the resolver.go file
	h := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: resolver}))

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

// Defining the Playground handler
func playgroundHandler() gin.HandlerFunc {
	h := playground.Handler("GraphQL", "/query")

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func main() {

	// Open up DB connection to use throughout the app
	DB, err := db.InitDB()
	if err != nil {
		panic(err)
	}

	resolvers := &graph.Resolver{
		BandRepo:        bands.BandRepo{DB: DB},
		AlbumRepo:       albums.AlbumRepo{DB: DB},
		SongRepo:        songs.SongRepo{DB: DB},
		VenueRepo:       venues.VenueRepo{DB: DB},
		PerformanceRepo: performances.PerformanceRepo{DB: DB},
	}

	// Setting up Gin
	r := gin.Default()

	// Apply the middleware to the router (Allow all origins)
	r.Use(cors.Default())

	r.POST("/query", graphqlHandler(resolvers))
	r.GET("/", playgroundHandler())
	r.Run()
}
