/*
   Copyright 2014 CoreOS, Inc.

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

package httptypes

import (
	"encoding/json"
	"net/url"
)

type Member struct {
	ID         uint64
	Name       string
	PeerURLs   []url.URL
	ClientURLs []url.URL
}

func (m *Member) UnmarshalJSON(data []byte) (err error) {
	rm := struct {
		ID         uint64
		Name       string
		PeerURLs   []string
		ClientURLs []string
	}{}

	if err := json.Unmarshal(data, &rm); err != nil {
		return err
	}

	parseURLs := func(strs []string) ([]url.URL, error) {
		urls := make([]url.URL, len(strs))
		for i, s := range strs {
			u, err := url.Parse(s)
			if err != nil {
				return nil, err
			}
			urls[i] = *u
		}

		return urls, nil
	}

	if m.PeerURLs, err = parseURLs(rm.PeerURLs); err != nil {
		return err
	}

	if m.ClientURLs, err = parseURLs(rm.ClientURLs); err != nil {
		return err
	}

	m.ID = rm.ID
	m.Name = rm.Name

	return nil
}

type MemberCollection struct {
	Members []Member
}
