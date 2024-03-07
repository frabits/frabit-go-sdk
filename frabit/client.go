// Frabit - The next-generation database automatic operation platform
// Copyright © 2022-2024 Frabit Team
//
// Licensed under the GNU General Public License, Version 3.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	https://www.gnu.org/licenses/gpl-3.0.txt
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package frabit

import "net/http"

const Version = "2.0.19"
const UserAgent = "frabit-go-sdk" + Version

type Client struct {
	BaseURL   string
	Client    http.Client
	UserAgent string
	Headers   map[string]interface{}
}

func NewClient(baseURL, token string) *Client {
	client := &Client{UserAgent: UserAgent}

	return client
}
