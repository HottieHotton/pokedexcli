package main

import (
	"testing"
	"time"
	"github.com/hottiehotton/pokedexcli/pokecache"

)

// Test Add and Get from the cache
func TestCleanInput(t *testing.T) {
	cache := pokecache.NewCache(time.Minute)
	key := "testKey"
	value := "testValue"

	cache.Add(key, []byte(value))
	retrievedValue, exists := cache.Get(key)
	if !exists {
		t.Errorf("Expected key %s to exist", key)
	}
	if string(retrievedValue) != value {
		t.Errorf("Expected value %v, got %v", value, string(retrievedValue))
	}
}

func TestCacheGetMiss(t *testing.T) {
	cache := pokecache.NewCache(5 * time.Minute)
	_, ok := cache.Get("missing")
	if ok {
		t.Error("expected cache miss for nonexistent key")
	}
}

func TestCacheReaping(t *testing.T) {
	interval := 10 * time.Millisecond
	cache := pokecache.NewCache(interval)

	cache.Add("old", []byte("old-value"))

	if _, ok := cache.Get("old"); !ok {
		t.Fatal("expected entry to exist right after adding")
	}

	// Wait for the entry to expire and the reaper to run
	time.Sleep(3 * interval)

	if _, ok := cache.Get("old"); ok {
		t.Error("expected entry to be reaped after expiry")
	}
}

func TestCacheReapingKeepsFresh(t *testing.T) {
	interval := 50 * time.Millisecond
	cache := pokecache.NewCache(interval)

	cache.Add("fresh", []byte("fresh-value"))

	// Wait less than the interval — entry should survive
	time.Sleep(interval / 2)

	got, ok := cache.Get("fresh")
	if !ok {
		t.Fatal("expected fresh entry to still exist")
	}
	if string(got) != "fresh-value" {
		t.Errorf("expected %q, got %q", "fresh-value", got)
	}
}