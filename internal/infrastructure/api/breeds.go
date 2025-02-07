package api

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type BreedsService struct {
	client *http.Client
}

func NewBreedsService() *BreedsService {
	return &BreedsService{
		client: &http.Client{
			Timeout: time.Second * 5,
		},
	}
}

func (bs *BreedsService) CheckBreed(ctx context.Context, name string) (bool, error) {
	rq, err := http.NewRequestWithContext(ctx, http.MethodGet, "https://api.thecatapi.com/v1/breeds", nil)
	if err != nil {
		return false, fmt.Errorf("error creating request: %w", err)
	}
	rs, err := bs.client.Do(rq)
	if err != nil {
		return false, fmt.Errorf("error doing request: %w", err)
	}
	defer rs.Body.Close()
	var breeds []struct {
		Name string `json:"name"`
	}
	if err := json.NewDecoder(rs.Body).Decode(&breeds); err != nil {
		return false, fmt.Errorf("error decoding response: %w", err)
	}
	for _, b := range breeds {
		if b.Name == name {
			return true, nil
		}
	}
	return false, nil
}
