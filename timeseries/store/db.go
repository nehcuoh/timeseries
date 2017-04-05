package store

import "zeus/utils/redis"

type Client struct {
	*redis.Client
}

