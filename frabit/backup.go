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

type BackupService interface {
	GetBackup(ctx context.Context) (*Backup, error)
	CreateBackup(ctx context.Context, req CreateBackupRequest) (*Backup, error)
}

type backupService struct {
	*Client
}

type Backup struct {
	Workspace string `json:"workspace"`
	Name      string `json:"name"`
	Owner     string `json:"owner"`
}

type CreateBackupRequest struct {
	Workspace string `json:"workspace"`
	Name      string `json:"name"`
	Owner     string `json:"owner"`
}

func (u *backupService) GetBackup(ctx context.Context) (*Backup, error) {
	req, _ := u.Client.newRequest("get", "database", nil)
	cls := &Backup{}
	err := u.Client.do(ctx, req, cls)
	if err != nil {
		return nil, err
	}

	return cls, nil
}

func (u *backupService) CreateBackup(ctx context.Context, CreateReq CreateBackupRequest) (*Backup, error) {
	req, _ := u.Client.newRequest("post", "database", CreateReq)
	user := &Backup{}
	err := u.Client.do(ctx, req, user)
	if err != nil {
		return nil, err
	}

	return user, nil
}
