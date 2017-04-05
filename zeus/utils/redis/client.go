/*
Copyright 2009-2016 Weibo, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package redis

import (
	"errors"
	"fmt"
	"hash/crc32"
	"math"
	"sync"
	"time"

	"github.com/garyburd/redigo/redis"
)

var (
	// ErrNotFound ...
	ErrNotFound = errors.New("redis: client: server not found")
)

// Client ...
type Client struct {
	mutex  sync.RWMutex
	config *Config
	master []*redis.Pool
	slave  []*redis.Pool
}

// NewClient ...
func NewClient(cfg *Config) *Client {
	c := &Client{config: cfg}
	c.master = c.newPool(cfg.Master)
	c.slave = c.newPool(cfg.Slave)
	return c
}

// Set ...
func (c *Client) Set(key string, value []byte) (string, error) {
	return redis.String(c.withKeyWrite(key, "SET", key, value))
}

// Get ...
func (c *Client) Get(key string) ([]byte, error) {
	reply, err := c.withKeyRead(key, "GET", key)
	if err != nil || reply == nil {
		return nil, err
	}
	return redis.Bytes(reply, err)
}

// Del ...
func (c *Client) Del(key string) (uint64, error) {
	return redis.Uint64(c.withKeyWrite(key, "DEL", key))
}

// HSet ...
func (c *Client) HSet(key string, hkey string, value []byte) (bool, error) {
	return redis.Bool(c.withKeyWrite(key, "HSET", key, hkey, value))
}



// HDel ...
func (c *Client) HDel(key string, hkeys ...string) (uint64, error) {
	args := make([]interface{}, len(hkeys)+1)
	args[0] = key
	for i, hk := range hkeys {
		args[i+1] = hk
	}
	return redis.Uint64(c.withKeyWrite(key, "HDEL", args...))
}

// HLen ...
func (c *Client) HLen(key string) (uint64, error) {
	return redis.Uint64(c.withKeyWrite(key, "HLEN", key))
}

// HGetAll ...
func (c *Client) HGetAll(key string) (map[string]string, error) {
	return redis.StringMap(c.withKeyRead(key, "HGETALL", key))
}

// HKeys ...
func (c *Client) HKeys(key string) ([]string, error) {
	return redis.Strings(c.withKeyRead(key, "HKEYS", key))
}

func hscan(rc redis.Conn, key string, pos int64, count int64) (int64, map[string]string, error) {
	reply, err := rc.Do("HSCAN", key, pos, "COUNT", count)
	if err != nil {
		return 0, nil, err
	}
	rps, err := redis.MultiBulk(reply, nil)
	if err != nil || len(rps) != 2 {
		return 0, nil, err
	}
	idx, err := redis.Int64(rps[0], nil)
	if err != nil {
		return 0, nil, err
	}
	kvs, err := redis.StringMap(rps[1], nil)
	if err != nil {
		return 0, nil, err
	}
	return idx, kvs, nil
}

// HScanAll ...
func (c *Client) HScanAll(key string, count int64, cb func(map[string]string) error) error {
	var (
		rc  redis.Conn
		idx int64
		err error
		kv  map[string]string
	)

	c.mutex.RLock()
	defer func() {
		if rc != nil {
			rc.Close()
		}
		c.mutex.RUnlock()
	}()
	pool := c.slave
	if pool == nil {
		pool = c.master
	}
	rp, err := c.peek(key, pool)
	if err != nil {
		return err
	}
	rc = rp.Get()

	idx = int64(0)
	for {
		idx, kv, err = hscan(rc, key, idx, count)
		if err != nil {
			return err
		}
		if kv != nil && len(kv) > 0 && cb != nil {
			if err = cb(kv); err != nil {
				return err
			}
		}
		if idx == 0 {
			break
		}
	}
	return nil
}

// RandomKey ...
func (c *Client) RandomKey() ([]byte, error) {
	key := fmt.Sprintf("%d", time.Now().Nanosecond())
	return redis.Bytes(c.withKeyRead(key, "RANDOMKEY"))
}

// GetM ...
func (c *Client) GetM(key string) ([]byte, error) {
	reply, err := c.withKeyWrite(key, "GET", key)
	if err != nil || reply == nil {
		return nil, err
	}
	return redis.Bytes(reply, err)
}

// HGet ...
func (c *Client) HGet(key string, hkey string) ([]byte, error) {
	reply, err := c.withKeyRead(key, "HGET", key, hkey)
	if err != nil || reply == nil {
		return nil, err
	}
	return redis.Bytes(reply, err)
}

// HGetM ...
func (c *Client) HGetM(key string, hkey string) ([]byte, error) {
	reply, err := c.withKeyWrite(key, "HGET", key, hkey)
	if err != nil || reply == nil {
		return nil, err
	}
	return redis.Bytes(reply, err)
}

// Close ...
func (c *Client) Close() {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.closePool(c.master)
	c.closePool(c.slave)
	c.master = nil
	c.slave = nil
}

func (c *Client) newPool(addrs []string) []*redis.Pool {
	var p []*redis.Pool
	if len(addrs) > 0 {
		p = make([]*redis.Pool, len(addrs))
		for k, addr := range addrs {
			p[k] = c.newConn(addr)
		}
	}
	return p
}

func (c *Client) closePool(pool []*redis.Pool) {
	for k, p := range pool {
		p.Close()
		pool[k] = nil
	}
}

func (c *Client) newConn(addr string) *redis.Pool {
	cfg := c.config
	return &redis.Pool{
		Dial: func() (redis.Conn, error) {
			c, err := redis.DialTimeout(
				"tcp", addr,
				cfg.ConnTimeout,
				cfg.ReadTimeout,
				cfg.WriteTimeout,
			)
			if err != nil {
				return nil, err
			}
			return c, nil
		},
		MaxIdle:     cfg.MaxIdle,
		MaxActive:   cfg.MaxActive,
		Wait:        cfg.Wait,
		IdleTimeout: cfg.IdleTimeout,
		TestOnBorrow: func(rc redis.Conn, t time.Time) error {
			if time.Since(t) < time.Minute {
				return nil
			}
			_, err := rc.Do("PING")
			return err
		},
	}
}

func (c *Client) withKeyWrite(key string, cmd string, args ...interface{}) (interface{}, error) {
	var rc redis.Conn
	c.mutex.RLock()
	defer func() {
		if rc != nil {
			rc.Close()
		}
		c.mutex.RUnlock()
	}()
	rp, err := c.peek(key, c.master)
	if err != nil {
		return nil, err
	}
	rc = rp.Get()
	return rc.Do(cmd, args...)
}

func (c *Client) withKeyRead(key string, cmd string, args ...interface{}) (interface{}, error) {
	var rc redis.Conn
	c.mutex.RLock()
	defer func() {
		if rc != nil {
			rc.Close()
		}
		c.mutex.RUnlock()
	}()
	pool := c.slave
	if pool == nil {
		pool = c.master
	}
	rp, err := c.peek(key, pool)
	if err != nil {
		return nil, err
	}
	rc = rp.Get()
	return rc.Do(cmd, args...)
}

func crc(data []byte) int {
	v := int32(crc32.ChecksumIEEE(data))
	return int(math.Abs(float64(v)))
}

func (c *Client) peek(key string, pool []*redis.Pool) (*redis.Pool, error) {
	if pool == nil || len(pool) == 0 {
		return nil, ErrNotFound
	}
	return pool[(crc([]byte(key)) % len(pool))], nil
}
