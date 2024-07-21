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

type AgentService interface {
	Register(ctx context.Context, req CreateAgentRequest) error
	Heartbeat(ctx context.Context, req CreateHeartbeat) error
}

type agentService struct {
	*Client
}

type CreateAgentRequest struct {
	AgentID string `json:"agent_id"`
	Name    string `json:"name"`
	Status  string `json:"status"`
}

type CreateHeartbeat struct {
	AgentID string `json:"agent_id"`
	Status  string `json:"status"`
}

func (s *agentService) Register(ctx context.Context, req CreateAgentRequest) error {
	request, err := s.Client.newRequest("post", "/agent/register", req)
	if err != nil {
		return err
	}
	return s.do(ctx, request, nil)
}

func (s *agentService) Heartbeat(ctx context.Context, req CreateHeartbeat) error {
	request, err := s.Client.newRequest("post", "/agent/register", req)
	if err != nil {
		return err
	}
	return s.do(ctx, request, nil)
}
