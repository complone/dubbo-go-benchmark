/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func (u *User) JavaClassName() string {
	return "com.dubbogo.pixiu.UserService1"
}

// User user obj.
type User struct {
	ID   string    `json:"id,omitempty"`
	Code int64     `json:"code,omitempty"`
	Name string    `json:"name,omitempty"`
	Age  int32     `json:"age,omitempty"`
	Time time.Time `json:"time,omitempty"`
}

type UserProvider struct {
}

var cache *userDB

// userDB cache user.
type userDB struct {
	// key is name, value is user obj
	nameIndex map[string]*User
	// key is code, value is user obj
	codeIndex map[int64]*User
	lock      sync.Mutex
}

// GetUserByName query by name, single param, PX config GET.
func (u *UserProvider) GetUserByName(ctx context.Context, name string) (*User, error) {
	outLn("Req GetUserByName name:%#v", name)
	r, ok := cache.GetByName(name)
	if ok {
		outLn("Req GetUserByName result:%#v", r)
		return r, nil
	}
	return nil, nil
}

func (db *userDB) GetByName(n string) (*User, bool) {
	db.lock.Lock()
	defer db.lock.Unlock()

	r, ok := db.nameIndex[n]
	return r, ok
}

func outLn(format string, args ...interface{}) {
	fmt.Printf("\033[32;40m"+format+"\033[0m\n", args...)
}
