package markets

// The MIT License (MIT)
//
// Copyright © 2017 Dmitry Moskowski
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

import (
	"fmt"
)

// ErrUnsupportedMarket is an error indicating that market `m`
// is not supported in code.
type ErrUnsupportedMarket struct {
	m string
}

func (e *ErrUnsupportedMarket) Error() string {
	return fmt.Sprintf(
		"Unsupported market '%s'",
		e.m,
	)
}

// NewErrUnsupportedMarket creates new ErrUnsupportedMarket.
func NewErrUnsupportedMarket(m string) error {
	return &ErrUnsupportedMarket{m}
}

//

// ErrUnsupportedClient is an error indicating that client `c`
// is not supported in code.
type ErrUnsupportedClient struct {
	c interface{}
}

func (e *ErrUnsupportedClient) Error() string {
	return fmt.Sprintf(
		"Unsupported client of type '%T'",
		e.c,
	)
}

// NewErrUnsupportedClient creates new ErrUnsupportedClient.
func NewErrUnsupportedClient(c interface{}) error {
	return &ErrUnsupportedClient{c}
}

//