// Copyright (c) 2020 SAP SE or an SAP affiliate company. All rights reserved. This file is licensed under the Apache Software License, v. 2 except as noted otherwise in the LICENSE file
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package predicate

import (
	"sigs.k8s.io/controller-runtime/pkg/event"
	"sigs.k8s.io/controller-runtime/pkg/predicate"
	"sigs.k8s.io/controller-runtime/pkg/runtime/inject"
)

type or struct {
	predicates []predicate.Predicate
}

func (o *or) orRange(f func(predicate.Predicate) bool) bool {
	for _, p := range o.predicates {
		if f(p) {
			return true
		}
	}
	return false
}

// Create implements Predicate.
func (o *or) Create(event event.CreateEvent) bool {
	return o.orRange(func(p predicate.Predicate) bool { return p.Create(event) })
}

// Delete implements Predicate.
func (o *or) Delete(event event.DeleteEvent) bool {
	return o.orRange(func(p predicate.Predicate) bool { return p.Delete(event) })
}

// Update implements Predicate.
func (o *or) Update(event event.UpdateEvent) bool {
	return o.orRange(func(p predicate.Predicate) bool { return p.Update(event) })
}

// Generic implements Predicate.
func (o *or) Generic(event event.GenericEvent) bool {
	return o.orRange(func(p predicate.Predicate) bool { return p.Generic(event) })
}

// InjectFunc implements Injector.
func (o *or) InjectFunc(f inject.Func) error {
	for _, p := range o.predicates {
		if err := f(p); err != nil {
			return err
		}
	}
	return nil
}

// Or builds a logical OR gate of passed predicates.
func Or(predicates ...predicate.Predicate) predicate.Predicate {
	return &or{predicates}
}