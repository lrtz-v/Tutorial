package main

import (
	"context"
	"testing"
	"time"
)

func TestNewConnect(t *testing.T) {
	client := NewConnect(Config)
	if client == nil {
		t.Log("Create new connection failed")
	}
	defer client.Close()
}

func TestGetAndPut(t *testing.T) {
	client := NewConnect(Config)
	if client == nil {
		t.Log("Create new connection failed")
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// set
	client.SetString(ctx, "test", "test")
	// get
	value, err := client.GetString(ctx, "test")
	if err != nil {
		t.Fatal(err.Error())
	}
	if value != "test" {
		t.Fatalf("Get wrong value: %s ,expect value: test", value)
	}

	defer client.Close()
}

func TestDelete(t *testing.T) {
	client := NewConnect(Config)
	if client == nil {
		t.Log("Create new connection failed")
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	// set
	client.SetString(ctx, "test", "test")
	// get
	err := client.Delete("test")
	if err != nil {
		t.Fatal(err.Error())
	}

	defer client.Close()
}

func TestWatchKey(t *testing.T) {
	client := NewConnect(Config)
	if client == nil {
		t.Log("Create new connection failed")
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second * time.Duration(3))
	defer cancel()

	key := "test"
	value := "test"
	valueAfterChange := "test_after_change"
	// set
	client.SetString(ctx, key, value)
	// get
	client.Watch(key)

	go func() {
		time.Sleep(1 * time.Second)
		client.SetString(ctx, key, valueAfterChange)
	}()

	time.Sleep(3 * time.Second)
	defer client.Close()
}

func TestWatchPrefix(t *testing.T) {
	client := NewConnect(Config)
	if client == nil {
		t.Log("Create new connection failed")
	}

	// watch prefix
	client.Watch("")
	time.Sleep(1 * time.Second)

	key := "test"
	value := "test"
	valueAfterChange := "test_after_change"

	ctx, cancel := context.WithTimeout(context.Background(), time.Second * time.Duration(3))
	defer cancel()
	// set
	client.SetString(ctx, key, value)
	client.SetString(ctx, key, valueAfterChange)

	time.Sleep(1 * time.Second)
	defer client.Close()
}

func TestLease(t *testing.T) {
	client := NewConnect(Config)
	if client == nil {
		t.Log("Create new connection failed")
	}

	TTL := 5
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(TTL)*time.Second)
	defer cancel()

	idc, err := client.LeaseGrant(ctx, 5)
	if err != nil {
		t.Fatal(err)
	}
	client.Log.DebugF("Lease ID: %d", idc)

	err = client.LeaseKeepAlive(ctx, idc)
	if err != nil {
		t.Fatal(err)
	}

	ttl, err := client.LeaseTTL(ctx, idc)
	if err != nil {
		t.Fatal(err)
	}
	client.Log.DebugF("TTL: %d", ttl)

	err = client.LeaseRevoke(ctx, idc)
	if err != nil {
		t.Fatal(err)
	}

	defer client.Close()
}

func TestLock(t *testing.T) {
	client := NewConnect(Config)
	if client == nil {
		t.Log("Create new connection failed")
	}
	defer client.Close()

	TTL := 10
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(TTL)*time.Second)
	defer cancel()

	client.Lock(ctx, "app_lock", 5, func() {
		client.Log.Info("Run function with lock")
	})
	time.Sleep(1 * time.Second)
}
