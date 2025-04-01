package api

import (
	"cats/internal/domain/service"
	"context"
	"errors"
	"github.com/sony/gobreaker/v2"
)

type BreedsWithFallbackService struct {
	root  service.BreedService
	cache map[string]bool
	cb    *gobreaker.CircuitBreaker[bool]
}

func NewBreedsWithFallbackService(root service.BreedService) service.BreedService {
	var st gobreaker.Settings
	st.Name = "Breeds Service"
	st.ReadyToTrip = func(counts gobreaker.Counts) bool {
		failureRatio := float64(counts.TotalFailures) / float64(counts.Requests)
		return counts.Requests >= 3 && failureRatio >= 0.6
	}
	return &BreedsWithFallbackService{
		root:  root,
		cache: map[string]bool{},
		cb:    gobreaker.NewCircuitBreaker[bool](st),
	}
}

func (b *BreedsWithFallbackService) CheckBreed(ctx context.Context, name string) (bool, error) {
	res, err := b.cb.Execute(func() (bool, error) {
		r, err := b.root.CheckBreed(ctx, name)
		if err == nil && r {
			b.cache[name] = true
		}
		return r, err
	})
	if err != nil {
		if errors.Is(err, gobreaker.ErrOpenState) {
			if _, ok := b.cache[name]; ok {
				return true, nil
			}
		}
		return false, err
	}
	return res, nil
}
