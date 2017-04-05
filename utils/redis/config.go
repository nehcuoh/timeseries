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

import "time"

// Config ...
type Config struct {
	MaxIdle      int
	MaxActive    int
	Wait         bool
	IdleTimeout  time.Duration
	ConnTimeout  time.Duration
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	Master       []string
	Slave        []string
}

// DefaultConfig ...
func DefaultConfig() *Config {
	return &Config{
		MaxIdle:      64,
		MaxActive:    64,
		Wait:         true,
		IdleTimeout:  30 * time.Second,
		ConnTimeout:  100 * time.Millisecond,
		ReadTimeout:  100 * time.Millisecond,
		WriteTimeout: 100 * time.Millisecond,
		Master:       []string{},
		Slave:        []string{},
	}
}
