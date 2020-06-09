/*
 * Copyright The Spinnaker Authors
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package convert_test

import (
	"testing"

	"github.com/spinnaker/kleat/api/client/config"
	"github.com/spinnaker/kleat/api/client/security"
	"github.com/spinnaker/kleat/api/client/security/authz"
	"github.com/spinnaker/kleat/internal/convert"
	"google.golang.org/protobuf/proto"
)

var fiatTests = []testCase{
	{
		"Empty hal config",
		&config.Hal{},
		&config.Fiat{},
	},
	{
		"Authorization",
		&config.Hal{
			Security: &security.Security{
				Authz: &authz.Authorization{
					Enabled: true,
					GroupMembership: &authz.GroupMembership{
						Service: authz.GroupMembership_GITHUB,
						Github: &authz.GithubRoleProvider{
							AccessToken:  "my-token",
							Organization: "my-org",
						},
						File: &authz.FileRoleProvider{
							Path: "/var/secrets/my-file.csv",
						},
					},
				},
			},
		},
		&config.Fiat{
			Auth: &authz.Authorization{
				Enabled: true,
				GroupMembership: &authz.GroupMembership{
					Service: authz.GroupMembership_GITHUB,
					Github: &authz.GithubRoleProvider{
						AccessToken:  "my-token",
						Organization: "my-org",
					},
					File: &authz.FileRoleProvider{
						Path: "/var/secrets/my-file.csv",
					},
				},
			},
		},
	},
}

func TestHalToFiat(t *testing.T) {
	for _, tt := range fiatTests {
		t.Run(tt.name, func(t *testing.T) {
			got := convert.HalToFiat(tt.hal)
			if !proto.Equal(got, tt.want) {
				t.Errorf("Expected hal config to generate %v, got %v", tt.want, got)
			}
		})
	}
}
