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

type TeamService interface {
	GetTeam(ctx context.Context) (*Team, error)
	CreateTeam(ctx context.Context, req CreateTeamRequest) (*Team, error)
}

type teamService struct {
	*Client
}

type Team struct {
	Id          uint32 `json:"id"`
	Name        string `json:"dame"`
	Description string `json:"description"`
	Owner       string `json:"owner"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

type CreateTeamRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Owner       string `json:"owner"`
}

func (t *teamService) GetTeam(ctx context.Context) (*Team, error) {
	req, _ := t.Client.newRequest("get", "team", nil)
	db := &Team{}
	err := t.Client.do(ctx, req, db)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func (t *teamService) CreateTeam(ctx context.Context, CreateReq CreateTeamRequest) (*Team, error) {
	req, _ := t.Client.newRequest("post", "team", CreateReq)
	db := &Team{}
	err := t.Client.do(ctx, req, db)
	if err != nil {
		return nil, err
	}

	return db, nil
}
