//go:build !enterprise

package storage

import (
	"context"
	"fmt"
)

// ColdTierConfig specifies the remote cold-storage backend and policy.
type ColdTierConfig struct {
	Endpoint        string `json:"endpoint"`
	RemoteBucket    string `json:"remote_bucket"`
	Region          string `json:"region"`
	AccessKey       string `json:"access_key,omitempty"`
	SecretKey       string `json:"secret_key,omitempty"`
	MinAgeDays      int    `json:"min_age_days"`
	ScanIntervalMin int    `json:"scan_interval_min"`
}

type ColdTierManager struct {
	cfg ColdTierConfig
}

func newColdTierManager(store *LocalStore, cfg ColdTierConfig) *ColdTierManager {
	return nil
}

// stubPath returns the .cold stub path for a given data-file path.
// In the OSS build cold tiering is not active, but local_store.go references
// this helper for the re-hydration guard; return the path so the guard is
// structurally correct but will never trigger (coldTier is always nil).
func stubPath(dataPath string) string {
	return dataPath + ".cold"
}

// FetchBack is a no-op in the OSS build (coldTier is always nil, so this
// method is never called). It exists purely to satisfy the compiler.
func (m *ColdTierManager) FetchBack(_ context.Context, _ string) error {
	return fmt.Errorf("cold storage tiering requires ServStore Enterprise Edition")
}

func (s *LocalStore) SetColdTier(cfg ColdTierConfig) error {
	return fmt.Errorf("cold storage tiering requires ServStore Enterprise Edition")
}

func (s *LocalStore) GetColdTierConfig() (ColdTierConfig, bool) {
	return ColdTierConfig{}, false
}

func (s *LocalStore) RunColdSweep(ctx context.Context) (int, []error) {
	return 0, []error{fmt.Errorf("cold storage tiering requires ServStore Enterprise Edition")}
}
