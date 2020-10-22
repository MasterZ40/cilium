// Copyright 2017-2020 Authors of Cilium
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by lister-gen. DO NOT EDIT.

package v2

import (
	v2 "github.com/cilium/cilium/pkg/k8s/apis/cilium.io/v2"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// CiliumEgressRouteLister helps list CiliumEgressRoutes.
// All objects returned here must be treated as read-only.
type CiliumEgressRouteLister interface {
	// List lists all CiliumEgressRoutes in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v2.CiliumEgressRoute, err error)
	// Get retrieves the CiliumEgressRoute from the index for a given name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v2.CiliumEgressRoute, error)
	CiliumEgressRouteListerExpansion
}

// ciliumEgressRouteLister implements the CiliumEgressRouteLister interface.
type ciliumEgressRouteLister struct {
	indexer cache.Indexer
}

// NewCiliumEgressRouteLister returns a new CiliumEgressRouteLister.
func NewCiliumEgressRouteLister(indexer cache.Indexer) CiliumEgressRouteLister {
	return &ciliumEgressRouteLister{indexer: indexer}
}

// List lists all CiliumEgressRoutes in the indexer.
func (s *ciliumEgressRouteLister) List(selector labels.Selector) (ret []*v2.CiliumEgressRoute, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v2.CiliumEgressRoute))
	})
	return ret, err
}

// Get retrieves the CiliumEgressRoute from the index for a given name.
func (s *ciliumEgressRouteLister) Get(name string) (*v2.CiliumEgressRoute, error) {
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v2.Resource("ciliumegressroute"), name)
	}
	return obj.(*v2.CiliumEgressRoute), nil
}
