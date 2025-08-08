// Copyright 2018 The Fire Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

//go:build fireinject
// +build fireinject

package main

import (
	"github.com/google/fire"
)

func injectFoo() Foo {
	// provideFoo returns an error, but injectFoo does not.
	fire.Build(provideFoo)
	return Foo(0)
}
