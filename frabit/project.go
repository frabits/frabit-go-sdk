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

type ProjectService interface {
	GetProject(ctx context.Context) (*Project, error)
	CreateProject(ctx context.Context, req CreateProjectRequest) (*Project, error)
}

type projectService struct {
	*Client
}

type Project struct {
	Workspace string `json:"workspace"`
	Name      string `json:"name"`
	Owner     string `json:"owner"`
}

type CreateProjectRequest struct {
	Workspace string `json:"workspace"`
	Name      string `json:"name"`
	Owner     string `json:"owner"`
}

func (u *projectService) GetProject(ctx context.Context) (*Project, error) {
	req, _ := u.Client.newRequest("get", "database", nil)
	user := &Project{}
	err := u.Client.do(ctx, req, user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *projectService) CreateProject(ctx context.Context, CreateReq CreateProjectRequest) (*Project, error) {
	req, _ := u.Client.newRequest("post", "database", CreateReq)
	user := &Project{}
	err := u.Client.do(ctx, req, user)
	if err != nil {
		return nil, err
	}

	return user, nil
}
