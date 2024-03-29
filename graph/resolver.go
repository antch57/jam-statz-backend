package graph

import (
	"github.com/antch57/jam-statz/pkg/albums"
	"github.com/antch57/jam-statz/pkg/bands"
	"github.com/antch57/jam-statz/pkg/performances"
	"github.com/antch57/jam-statz/pkg/songs"
	"github.com/antch57/jam-statz/pkg/venues"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	BandRepo        bands.BandRepo
	AlbumRepo       albums.AlbumRepo
	SongRepo        songs.SongRepo
	VenueRepo       venues.VenueRepo
	PerformanceRepo performances.PerformanceRepo
}
