package main

import (
	"context"
	"fmt"
	"strings"
	"time"

	// "go.etcd.io/etcd/clientv3"
	// "go.etcd.io/etcd/clientv3/concurrency"
	"github.com/coreos/etcd/clientv3"
	"github.com/coreos/etcd/clientv3/concurrency"
)

type EtcdClient struct {
	client *clientv3.Client
	config *EtcdConfig
	Log    *Log
}

func NewConnect(config *EtcdConfig) *EtcdClient {
	client, err := clientv3.New(clientv3.Config{
		Endpoints:   config.Endpoints,
		DialTimeout: 10 * time.Second,
		Username:    config.Username,
		Password:    config.Password,
	})

	if err != nil {
		fmt.Printf("Create new connection failed, err: %s", err.Error())
		return nil
	}

	return &EtcdClient{client: client, config: config, Log: Logger}
}

func (etcd *EtcdClient) Close() {
	etcd.client.Close()
}

func (etcd *EtcdClient) GetPrefix() string {
	return etcd.config.GetPrefix()
}

func (etcd *EtcdClient) keyJoin(key string) string {
	// return key
	return strings.Join([]string{etcd.config.GetPrefix(), key}, "/")
}

func (etcd *EtcdClient) SetString(ctx context.Context, key string, value string) {
	key = etcd.keyJoin(key)
	ctx, cancel := context.WithTimeout(ctx, time.Second)
	_, err := etcd.client.Put(ctx, key, value)
	cancel()

	if err != nil {
		etcd.Log.ErrorF("SET key: %s, value: %s failed: %s", key, value, err.Error())
		return
	}

	etcd.Log.DebugF("SET key: %s, value: %s Success", key, value)
}

func (etcd *EtcdClient) GetString(ctx context.Context, key string) (string, error) {
	key = etcd.keyJoin(key)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	resp, err := etcd.client.Get(ctx, key)
	cancel()

	if err != nil {
		etcd.Log.ErrorF("GET key: %s, failed: %s", key, err.Error())
		return "", err
	}
	if len(resp.Kvs) == 0 {
		return "", fmt.Errorf("Key: %s not exist", key)
	}
	value := string(resp.Kvs[0].Value)
	etcd.Log.DebugF("GET key: %s Success, value: %s", key, value)
	return value, nil
}

func (etcd *EtcdClient) Delete(key string) error {
	key = etcd.keyJoin(key)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	_, err := etcd.client.Delete(ctx, key)
	cancel()

	if err != nil {
		etcd.Log.ErrorF("GET key: %s, failed: %s", key, err.Error())
		return err
	}
	return nil
}

func (etcd *EtcdClient) watch(key string) {
	rch := etcd.client.Watch(context.Background(), key)
	etcd.Log.DebugF("Watch Key: %s", key)
	for wresp := range rch {
		for _, ev := range wresp.Events {
			etcd.Log.DebugF("Operation Type: %s Key: %s Value: %s\n", ev.Type, ev.Kv.Key, ev.Kv.Value)
		}
	}
}

func (etcd *EtcdClient) watchPrefix(key string) {
	rch := etcd.client.Watch(context.Background(), key, clientv3.WithPrefix())
	etcd.Log.DebugF("Watch Prefix: %s", key)
	for wresp := range rch {
		for _, ev := range wresp.Events {
			etcd.Log.DebugF("Operation Type: %s Key: %s Value: %s\n", ev.Type, ev.Kv.Key, ev.Kv.Value)
		}
	}
}

func (etcd *EtcdClient) Watch(key string) {
	if len(key) == 0 {
		key = etcd.GetPrefix() + "/"
		go etcd.watchPrefix(key)
	} else {
		key = etcd.keyJoin(key)
		go etcd.watch(key)
	}
}

func (etcd *EtcdClient) LeaseGrant(ctx context.Context, ttl int64) (clientv3.LeaseID, error) {
	resp, err := etcd.client.Grant(ctx, ttl)
	if err != nil {
		etcd.Log.ErrorF("LeaseGrant failed: %s", err.Error())
		return -1, err
	}
	etcd.Log.DebugF("LeaseGrant: %d\n", resp.ID)
	return resp.ID, nil
}

func (etcd *EtcdClient) LeaseKeepAlive(ctx context.Context, id clientv3.LeaseID) error {
	_, err := etcd.client.KeepAliveOnce(ctx, id)
	if err != nil {
		etcd.Log.ErrorF("LeaseKeepAlive failed: %s", err.Error())
		return err
	}
	return nil
}

func (etcd *EtcdClient) LeaseRevoke(ctx context.Context, id clientv3.LeaseID) error {
	_, err := etcd.client.Revoke(ctx, id)
	if err != nil {
		etcd.Log.ErrorF("LeaseRevoke failed: %s", err.Error())
		return err
	}
	return nil
}

func (etcd *EtcdClient) LeaseTTL(ctx context.Context, id clientv3.LeaseID) (int64, error) {
	resp, err := etcd.client.TimeToLive(ctx, id)
	if err != nil {
		etcd.Log.ErrorF("LeaseTTL failed: %s", err.Error())
		return -1, err
	}
	return resp.TTL, nil
}

func (etcd *EtcdClient) Lock(ctx context.Context, key string, timeout int64, x func()) {
	key = etcd.keyJoin(key)

	leaseID, err := etcd.LeaseGrant(ctx, timeout)
	if err != nil {
		etcd.Log.ErrorF("lease err : %s", err)
	}
	session, err := concurrency.NewSession(etcd.client, concurrency.WithLease(leaseID))
	if err != nil {
		etcd.Log.ErrorF("session err : %s", err)
	}
	defer session.Close()

	mutex := concurrency.NewMutex(session, key)
	err = mutex.Lock(ctx)
	if err != nil {
		etcd.Log.ErrorF("Get Mutex failed: %s", err)
	}

	x()

	err = mutex.Unlock(ctx)
	if err != nil {
		etcd.Log.ErrorF("Unlock mutex error: %s", err)
	}
}
