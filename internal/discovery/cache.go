// Package discovery provides simple cache management
package discovery

import (
	"errors"
	"log"
)

// Initialize loads or creates the path cache
// This is the main entry point that should be called from main.go
func Initialize() (*PathCache, error) {
	cache, err := NewPathCache()
	if err != nil {
		return nil, errors.New("failed to initialize path cache: " + err.Error())
	}

	log.Printf("Path cache initialized successfully")
	return cache, nil
}

// ForceRescan forces a complete rescan of hardware paths
func (c *PathCache) ForceRescan() error {
	log.Printf("Performing forced hardware rescan...")
	
	if err := c.Scan(); err != nil {
		return errors.New("failed to rescan hardware: " + err.Error())
	}

	if err := c.Save(); err != nil {
		return errors.New("failed to save rescanned cache: " + err.Error())
	}

	log.Printf("Hardware rescan completed and cache updated")
	return nil
}