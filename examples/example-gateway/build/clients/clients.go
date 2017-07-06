// Code generated by zanzibar
// @generated

// Copyright (c) 2017 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package clients

import (
	barClientGenerated "github.com/uber/zanzibar/examples/example-gateway/build/clients/bar"
	bazClientGenerated "github.com/uber/zanzibar/examples/example-gateway/build/clients/baz"
	contactsClientGenerated "github.com/uber/zanzibar/examples/example-gateway/build/clients/contacts"
	googlenowClientGenerated "github.com/uber/zanzibar/examples/example-gateway/build/clients/google-now"
	quuxClientGenerated "github.com/uber/zanzibar/examples/example-gateway/build/clients/quux"
	quuxClientStatic "github.com/uber/zanzibar/examples/example-gateway/clients/quux"

	"github.com/uber/zanzibar/runtime"
)

// Clients datastructure that holds all the generated clients
// This should only hold clients generate from specs
type Clients struct {
	Bar       barClientGenerated.Client
	Baz       bazClientGenerated.Client
	Contacts  contactsClientGenerated.Client
	GoogleNow googlenowClientGenerated.Client
	Quux      quuxClientStatic.Client
}

// CreateClients will make all clients
func CreateClients(
	gateway *zanzibar.Gateway,
) interface{} {
	_bar := barClientGenerated.NewClient(gateway)
	_baz := bazClientGenerated.NewClient(gateway)
	_contacts := contactsClientGenerated.NewClient(gateway)
	_googlenow := googlenowClientGenerated.NewClient(gateway)
	_quux := quuxClientStatic.NewClient(gateway, quuxClientGenerated.ClientDependencies{
		Bar: _bar,
	})

	return &Clients{
		Bar:       _bar,
		Baz:       _baz,
		Contacts:  _contacts,
		GoogleNow: _googlenow,
		Quux:      _quux,
	}
}
