/*
Copyright 2014 The Camlistore Authors

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

package internal

import (
	"sync"
	"sync/atomic"
)

// A Once will perform a successful action exactly once.
// Once 将会执行一次成功的操作。
//
// Unlike a sync.Once, this Once's func returns an error
// and is re-armed on failure.
// 与 sync.Once 不同，此 Once 函数会返回错误，并在失败时重新启动。
type Once struct {
	m    sync.Mutex
	done uint32
}

// Do calls the function f if and only if Do has not been invoked
// without error for this instance of Once.  In other words, given
//
//	var once Once
//
// if once.Do(f) is called multiple times, only the first call will
// invoke f, even if f has a different value in each invocation unless
// f returns an error.  A new instance of Once is required for each
// function to execute.
//
// Do is intended for initialization that must be run exactly once.  Since f
// is niladic, it may be necessary to use a function literal to capture the
// arguments to a function to be invoked by Do:
//
//	err := config.once.Do(func() error { return config.init(filename) })
//
// --------------------------------------------------------------------------
// ch:
// 当且仅当 Do 从未被无错误地调用过此 Once 实例时，Do 才会调用函数 f。换句话说，给定
// var once Once
//
// 如果 once.Do(f) 被多次调用，则只有第一次调用会调用 f，即使 f 在每次调用中都有不同的值，除非 f 返回错误。
// 每次执行函数都需要一个新的 Once 实例。
func (o *Once) Do(f func() error) error {
	if atomic.LoadUint32(&o.done) == 1 {
		return nil
	}
	// Slow-path.
	o.m.Lock()
	defer o.m.Unlock()
	var err error
	if o.done == 0 {
		err = f()
		if err == nil {
			atomic.StoreUint32(&o.done, 1)
		}
	}
	return err
}
