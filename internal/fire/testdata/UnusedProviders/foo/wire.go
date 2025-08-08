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
	"github.com/myyrakle/fire"
)

func injectFooBar() FooBar {
	fire.Build(
		provideFoo,                       // needed as input for provideBar
		provideBar,                       // needed for FooBar
		partiallyUsedSet,                 // 1/2 providers in the set are needed
		provideUnused,                    // not needed -> error
		fire.Value("unused"),             // not needed -> error
		unusedSet,                        // nothing in set is needed -> error
		fire.Bind(new(Fooer), new(*Foo)), // binding to Fooer is not needed -> error
		fire.FieldsOf(new(S), "Cfg"),     // S.Cfg not needed -> error
		fire.Struct(new(FooBar), "MyFoo", "MyBar"), // needed for FooBar
	)
	return FooBar{}
}
