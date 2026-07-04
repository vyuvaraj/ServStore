//go:build enterprise

package storage

import (
	"context"
)

// SetColdTier configures and starts the cold-storage tiering background sweep.
// Calling SetColdTier again replaces the previous configuration and restarts
// the sweep goroutine.
func (s *LocalStore) SetColdTier(cfg ColdTierConfig) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	// Cancel any existing sweep goroutine
	if s.coldCancel != nil {
		s.coldCancel()
	}

	mgr := newColdTierManager(s, cfg)
	s.coldTier = mgr

	sweepCtx, cancel := context.WithCancel(context.Background())
	s.coldCancel = cancel
	mgr.Start(sweepCtx)
	return nil
}

// GetColdTierConfig returns the active cold-tier config and whether one is set.
func (s *LocalStore) GetColdTierConfig() (ColdTierConfig, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	if s.coldTier == nil {
		return ColdTierConfig{}, false
	}
	return s.coldTier.cfg, true
}

// RunColdSweep triggers an immediate cold-tier archival sweep.
// Satisfies the optional sweeper interface used by the S3 API handler.
// Returns the number of blocks archived and any accumulated errors.
func (s *LocalStore) RunColdSweep(ctx context.Context) (int, []error) {
	s.mu.RLock()
	mgr := s.coldTier
	s.mu.RUnlock()
	if mgr == nil {
		return 0, nil
	}
	return mgr.RunSweep(ctx)
}
