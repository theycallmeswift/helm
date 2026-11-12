/*
Copyright The Helm Authors.

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

package urlutil

import (
	"net/url"
	"path"
	"path/filepath"
	"strings"
)

// URLJoin joins a base URL to one or more path components.
//
// It's like filepath.Join for URLs. If the baseURL is pathish, this will still
// perform a join.
//
// If the URL is unparsable, this returns an error.
func URLJoin(baseURL string, paths ...string) (string, error) {
	u, err := url.Parse(baseURL)
	if err != nil {
		return "", err
	}
	// We want path instead of filepath because path always uses /.
	all := []string{u.Path}
	all = append(all, paths...)
	u.Path = path.Join(all...)
	return u.String(), nil
}

// Equal normalizes two URLs and then compares for equality.
//
// The scheme and host are compared case-insensitively; the path is not.
func Equal(a, b string) bool {
	au, err := url.Parse(a)
	if err != nil {
		a = filepath.Clean(a)
		b = filepath.Clean(b)
		// If urls are paths, return true only if they are an exact match
		return a == b
	}
	bu, err := url.Parse(b)
	if err != nil {
		return false
	}

	for _, u := range []*url.URL{au, bu} {
		if u.Path == "" {
			u.Path = "/"
		}
		u.Path = filepath.Clean(u.Path)
		// The scheme and host are case-insensitive (RFC 3986 section 3.1 and 3.2.2).
		u.Scheme = strings.ToLower(u.Scheme)
		u.Host = strings.ToLower(u.Host)
	}
	return au.String() == bu.String()
}

// ExtractHostname returns hostname from URL
func ExtractHostname(addr string) (string, error) {
	u, err := url.Parse(addr)
	if err != nil {
		return "", err
	}
	return u.Hostname(), nil
}
