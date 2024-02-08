package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.43

import (
	"context"
	"fmt"

	"github.com/antch57/jam-statz/graph/model"
)

// CreateBand is the resolver for the createBand field.
func (r *mutationResolver) CreateBand(ctx context.Context, input *model.BandInput) (*model.Band, error) {
	res, err := r.BandRepo.CreateBand(input)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// UpdateBand is the resolver for the updateBand field.
func (r *mutationResolver) UpdateBand(ctx context.Context, id int, input *model.BandInput) (*model.Band, error) {
	panic(fmt.Errorf("not implemented: UpdateBand - updateBand"))
}

// DeleteBand is the resolver for the deleteBand field.
func (r *mutationResolver) DeleteBand(ctx context.Context, id int) (*bool, error) {
	panic(fmt.Errorf("not implemented: DeleteBand - deleteBand"))
}

// CreateAlbum is the resolver for the createAlbum field.
func (r *mutationResolver) CreateAlbum(ctx context.Context, input *model.AlbumInput) (*model.Album, error) {
	res, err := r.AlbumRepo.CreateAlbum(input)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// UpdateAlbum is the resolver for the updateAlbum field.
func (r *mutationResolver) UpdateAlbum(ctx context.Context, id int, input *model.AlbumInput) (*model.Album, error) {
	panic(fmt.Errorf("not implemented: UpdateAlbum - updateAlbum"))
}

// DeleteAlbum is the resolver for the deleteAlbum field.
func (r *mutationResolver) DeleteAlbum(ctx context.Context, id int) (*bool, error) {
	panic(fmt.Errorf("not implemented: DeleteAlbum - deleteAlbum"))
}

// CreateVenue is the resolver for the createVenue field.
func (r *mutationResolver) CreateVenue(ctx context.Context, input *model.VenueInput) (*model.Venue, error) {
	panic(fmt.Errorf("not implemented: CreateVenue - createVenue"))
}

// UpdateVenue is the resolver for the updateVenue field.
func (r *mutationResolver) UpdateVenue(ctx context.Context, id int, input *model.VenueInput) (*model.Venue, error) {
	panic(fmt.Errorf("not implemented: UpdateVenue - updateVenue"))
}

// DeleteVenue is the resolver for the deleteVenue field.
func (r *mutationResolver) DeleteVenue(ctx context.Context, id int) (*bool, error) {
	panic(fmt.Errorf("not implemented: DeleteVenue - deleteVenue"))
}

// CreatePerformance is the resolver for the createPerformance field.
func (r *mutationResolver) CreatePerformance(ctx context.Context, input *model.PerformanceInput) (*model.Performance, error) {
	panic(fmt.Errorf("not implemented: CreatePerformance - createPerformance"))
}

// UpdatePerformance is the resolver for the updatePerformance field.
func (r *mutationResolver) UpdatePerformance(ctx context.Context, id int, input *model.PerformanceInput) (*model.Performance, error) {
	panic(fmt.Errorf("not implemented: UpdatePerformance - updatePerformance"))
}

// DeletePerformance is the resolver for the deletePerformance field.
func (r *mutationResolver) DeletePerformance(ctx context.Context, id int) (*bool, error) {
	panic(fmt.Errorf("not implemented: DeletePerformance - deletePerformance"))
}

// CreateSong is the resolver for the createSong field.
func (r *mutationResolver) CreateSong(ctx context.Context, input *model.SongInput) (*model.Song, error) {
	res, err := r.SongRepo.CreateSong(input)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// UpdateSong is the resolver for the updateSong field.
func (r *mutationResolver) UpdateSong(ctx context.Context, id int, input *model.SongInput) (*model.Song, error) {
	res, err := r.SongRepo.UpdateSong(id, input)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// DeleteSong is the resolver for the deleteSong field.
func (r *mutationResolver) DeleteSong(ctx context.Context, id int) (*bool, error) {
	res, err := r.SongRepo.DeleteSong(id)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

// CreatePerformanceSong is the resolver for the createPerformanceSong field.
func (r *mutationResolver) CreatePerformanceSong(ctx context.Context, input *model.PerformanceSongInput) (*model.PerformanceSong, error) {
	panic(fmt.Errorf("not implemented: CreatePerformanceSong - createPerformanceSong"))
}

// UpdatePerformanceSong is the resolver for the updatePerformanceSong field.
func (r *mutationResolver) UpdatePerformanceSong(ctx context.Context, id int, input *model.PerformanceSongInput) (*model.PerformanceSong, error) {
	panic(fmt.Errorf("not implemented: UpdatePerformanceSong - updatePerformanceSong"))
}

// DeletePerformanceSong is the resolver for the deletePerformanceSong field.
func (r *mutationResolver) DeletePerformanceSong(ctx context.Context, id int) (*bool, error) {
	panic(fmt.Errorf("not implemented: DeletePerformanceSong - deletePerformanceSong"))
}

// CreateAlbumSong is the resolver for the createAlbumSong field.
func (r *mutationResolver) CreateAlbumSong(ctx context.Context, input *model.AlbumSongInput) (*model.AlbumSong, error) {
	res, err := r.AlbumRepo.CreateAlbumSong(input)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// UpdateAlbumSong is the resolver for the updateAlbumSong field.
func (r *mutationResolver) UpdateAlbumSong(ctx context.Context, id int, input *model.AlbumSongInput) (*model.AlbumSong, error) {
	panic(fmt.Errorf("not implemented: UpdateAlbumSong - updateAlbumSong"))
}

// DeleteAlbumSong is the resolver for the deleteAlbumSong field.
func (r *mutationResolver) DeleteAlbumSong(ctx context.Context, id int) (*bool, error) {
	panic(fmt.Errorf("not implemented: DeleteAlbumSong - deleteAlbumSong"))
}

// Band is the resolver for the band field.
func (r *queryResolver) Band(ctx context.Context, id int) (*model.Band, error) {
	panic(fmt.Errorf("not implemented: Band - band query resolver"))
}

// Bands is the resolver for the bands field.
func (r *queryResolver) Bands(ctx context.Context) ([]*model.Band, error) {
	res, err := r.BandRepo.GetBands()
	if err != nil {
		return nil, err
	}

	return res, nil
}

// Album is the resolver for the album field.
func (r *queryResolver) Album(ctx context.Context, id int) (*model.Album, error) {
	panic(fmt.Errorf("not implemented: Album - album"))
}

// Albums is the resolver for the albums field.
func (r *queryResolver) Albums(ctx context.Context) ([]*model.Album, error) {
	panic(fmt.Errorf("not implemented: Albums - albums, query resolver"))
}

// Venue is the resolver for the venue field.
func (r *queryResolver) Venue(ctx context.Context, id int) (*model.Venue, error) {
	panic(fmt.Errorf("not implemented: Venue - venue"))
}

// Venues is the resolver for the venues field.
func (r *queryResolver) Venues(ctx context.Context) ([]*model.Venue, error) {
	panic(fmt.Errorf("not implemented: Venues - venues"))
}

// Performance is the resolver for the performance field.
func (r *queryResolver) Performance(ctx context.Context, id int) (*model.Performance, error) {
	panic(fmt.Errorf("not implemented: Performance - performance"))
}

// Performances is the resolver for the performances field.
func (r *queryResolver) Performances(ctx context.Context) ([]*model.Performance, error) {
	panic(fmt.Errorf("not implemented: Performances - performances"))
}

// Song is the resolver for the song field.
func (r *queryResolver) Song(ctx context.Context, id int) (*model.Song, error) {
	res, err := r.SongRepo.GetSongById(id)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// Songs is the resolver for the songs field.
func (r *queryResolver) Songs(ctx context.Context) ([]*model.Song, error) {
	res, err := r.SongRepo.GetSongs()
	if err != nil {
		return nil, err
	}

	return res, nil
}

// PerformanceSong is the resolver for the performanceSong field.
func (r *queryResolver) PerformanceSong(ctx context.Context, id int) (*model.PerformanceSong, error) {
	panic(fmt.Errorf("not implemented: PerformanceSong - performanceSong"))
}

// AlbumSong is the resolver for the albumSong field.
func (r *queryResolver) AlbumSong(ctx context.Context, id int) (*model.AlbumSong, error) {
	panic(fmt.Errorf("not implemented: AlbumSong - albumSong"))
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
