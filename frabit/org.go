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

import "context"

type OrgService interface {
	UpdateOrg(ctx context.Context, req OrgUpdateRequest) error
	CreateOrg(ctx context.Context, req OrgCreateRequest) error
}

type orgService struct {
	*Client
}

type OrgCreateRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Country     string `json:"country"`
}

type OrgUpdateRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Country     string `json:"country"`
}

func (s *orgService) UpdateOrg(ctx context.Context, req OrgUpdateRequest) error {
	return nil
}

func (s *orgService) CreateOrg(ctx context.Context, req OrgCreateRequest) error {
	return nil
}
