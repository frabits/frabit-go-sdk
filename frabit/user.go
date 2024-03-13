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

type UserService interface {
	GetTeam(ctx context.Context) (*User, error)
	CreateTeam(ctx context.Context, req CreateTeamRequest) (*User, error)
}

type userService struct {
	*Client
}

type User struct {
	Workspace string `json:"workspace"`
	Name      string `json:"name"`
	Owner     string `json:"owner"`
}

type CreateUserRequest struct {
	Workspace string `json:"workspace"`
	Name      string `json:"name"`
	Owner     string `json:"owner"`
}

func (u *userService) GetUser(ctx context.Context) (*User, error) {
	req, _ := u.Client.newRequest("get", "database", nil)
	user := &User{}
	err := u.Client.do(ctx, req, user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *userService) CreateUser(ctx context.Context, CreateReq CreateUserRequest) (*User, error) {
	req, _ := u.Client.newRequest("post", "database", CreateReq)
	user := &User{}
	err := u.Client.do(ctx, req, user)
	if err != nil {
		return nil, err
	}

	return user, nil
}
