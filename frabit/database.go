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

type DatabaseService interface {
	GetDatabase(ctx context.Context) (*Database, error)
	CreateDatabase(ctx context.Context, req CreateDatabaseRequest) (*Database, error)
}

type databaseService struct {
	*Client
}

type Database struct {
	Workspace string `json:"workspace"`
	Name      string `json:"name"`
	Admin     string `json:"admin"`
}

type CreateDatabaseRequest struct {
	Workspace string `json:"workspace"`
	Name      string `json:"name"`
	Owner     string `json:"owner"`
}

func (d *databaseService) GetDatabase(ctx context.Context) (*Database, error) {
	req, _ := d.Client.newRequest("get", "database", nil)
	db := &Database{}
	err := d.Client.do(ctx, req, db)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func (d *databaseService) CreateDatabase(ctx context.Context, CreateReq CreateDatabaseRequest) (*Database, error) {
	req, _ := d.Client.newRequest("post", "database", CreateReq)
	db := &Database{}
	err := d.Client.do(ctx, req, db)
	if err != nil {
		return nil, err
	}

	return db, nil
}
