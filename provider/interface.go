package provider

import "github.com/l3uddz/mediarr/config"

type Interface interface {
	Init(MediaType, map[string]string) error
	SetIgnoreExistingMediaItemFn(func(*config.MediaItem) bool)
	SetAcceptMediaItemFn(func(*config.MediaItem) bool)

	GetShowsSearchTypes() []string
	GetMoviesSearchTypes() []string
	SupportsShowsSearchType(string) bool
	SupportsMoviesSearchType(string) bool

	GetShows(string, map[string]interface{}, map[string]string) (map[string]config.MediaItem, error)
	GetMovies(string, map[string]interface{}, map[string]string) (map[string]config.MediaItem, error)
}
