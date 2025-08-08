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
	"strings"

	"github.com/google/fire"
)

func inject() Foo {
	// fail: provideFoo and provideFooAgain both provide Foo.
	panic(fire.Build(provideFoo, provideFooAgain))
}

func injectFromSet() Foo {
	// fail: provideFoo is also provided by Set.
	panic(fire.Build(provideFoo, Set))
}

func injectFromNestedSet() Foo {
	// fail: provideFoo is also provided by SuperSet, via Set.
	panic(fire.Build(provideFoo, SuperSet))
}

func injectFromSetWithDuplicateBindings() Foo {
	// fail: DuplicateBindingsSet has two providers for Foo.
	panic(fire.Build(SetWithDuplicateBindings))
}

func injectDuplicateValues() Foo {
	// fail: provideFoo and fire.Value both provide Foo.
	panic(fire.Build(provideFoo, fire.Value(Foo("foo"))))
}

func injectDuplicateInterface() Bar {
	// fail: provideBar and fire.Bind both provide Bar.
	panic(fire.Build(provideBar, fire.Bind(new(Bar), new(*strings.Reader))))
}
