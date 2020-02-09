package provider

import (
	"fmt"
	"github.com/jpillora/backoff"
	"github.com/l3uddz/mediarr/utils/web"
	"strings"
	"time"
)

var (
	providerDefaultTimeout = 15
	providerDefaultRetry   = web.Retry{
		MaxAttempts:          5,
		RetryableStatusCodes: []int{},
		Backoff: backoff.Backoff{
			Jitter: true,
			Min:    1 * time.Second,
			Max:    5 * time.Second,
		},
	}
)

/* Public */

func Get(providerType string) (Interface, error) {
	switch strings.ToLower(providerType) {
	case "tvmaze":
		return NewTvMaze(), nil
	case "tmdb":
		return NewTmdb(), nil
	default:
		break
	}

	return nil, fmt.Errorf("unsupported media provider provided: %q", providerType)
}
