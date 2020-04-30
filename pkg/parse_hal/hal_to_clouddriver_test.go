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
package parse_hal_test

import (
	"bytes"
	"io/ioutil"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/spinnaker/kleat/api/client/artifact"
	"github.com/spinnaker/kleat/api/client/cloudprovider"
	"github.com/spinnaker/kleat/api/client/config"
	"github.com/spinnaker/kleat/internal/protoyaml"
	"github.com/spinnaker/kleat/pkg/parse_hal"
)

var halToClouddriverTests = []struct {
	name string
	hal  *config.Hal
	want *config.Clouddriver
}{
	{
		"Empty hal config",
		&config.Hal{},
		&config.Clouddriver{},
	},
	{
		"Empty providers",
		&config.Hal{
			Providers: &cloudprovider.Providers{},
		},
		&config.Clouddriver{},
	},
	{
		"Empty Kubernetes provider",
		&config.Hal{
			Providers: &cloudprovider.Providers{
				Kubernetes: &cloudprovider.Kubernetes{},
			},
		},
		&config.Clouddriver{
			Kubernetes: &cloudprovider.Kubernetes{},
		},
	},
	{
		"Kubernetes account",
		&config.Hal{
			Providers: &cloudprovider.Providers{
				Kubernetes: &cloudprovider.Kubernetes{
					Enabled: true,
					Accounts: []*cloudprovider.KubernetesAccount{
						{
							Name:           "my-account",
							Kinds:          []string{"deployment"},
							OmitNamespaces: []string{"kube-system"},
						},
					},
					PrimaryAccount: "my-account",
				},
			},
		},
		&config.Clouddriver{
			Kubernetes: &cloudprovider.Kubernetes{
				Enabled: true,
				Accounts: []*cloudprovider.KubernetesAccount{
					{
						Name:           "my-account",
						Kinds:          []string{"deployment"},
						OmitNamespaces: []string{"kube-system"},
					},
				},
				PrimaryAccount: "my-account",
			},
		},
	},
	{
		"Empty artifacts",
		&config.Hal{
			Artifacts: &artifact.Artifacts{},
		},
		&config.Clouddriver{
			Artifacts: &artifact.Artifacts{},
		},
	},
	{
		"Empty GCS artifacts",
		&config.Hal{
			Artifacts: &artifact.Artifacts{
				Gcs: &artifact.Gcs{},
			},
		},
		&config.Clouddriver{
			Artifacts: &artifact.Artifacts{
				Gcs: &artifact.Gcs{},
			},
		},
	},
	{
		"GCS artifact account",
		&config.Hal{
			Artifacts: &artifact.Artifacts{
				Gcs: &artifact.Gcs{
					Enabled: true,
					Accounts: []*artifact.GcsAccount{
						{
							Name:     "my-account",
							JsonPath: "/var/secrets/my-key.json",
						},
					},
				},
			},
		},
		&config.Clouddriver{
			Artifacts: &artifact.Artifacts{
				Gcs: &artifact.Gcs{
					Enabled: true,
					Accounts: []*artifact.GcsAccount{
						{
							Name:     "my-account",
							JsonPath: "/var/secrets/my-key.json",
						},
					},
				},
			},
		},
	},
}

func TestHalToClouddriver(t *testing.T) {
	for _, tt := range halToClouddriverTests {
		t.Run(tt.name, func(t *testing.T) {
			got := parse_hal.HalToClouddriver(tt.hal)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Expected hal config to generate %v, got %v", tt.want, got)
			}
		})
	}
}

func TestHalToClouddriverYaml(t *testing.T) {
	data, err := ioutil.ReadFile(filepath.Join("../../testdata", "halconfig.yml"))
	if err != nil {
		t.Fatal(err)
	}

	h := &config.Hal{}
	if err := protoyaml.Unmarshal(data, h); err != nil {
		t.Fatal(err)
	}

	gotC := parse_hal.HalToClouddriver(h)

	wantC, err := parseClouddriverConfig(filepath.Join("../../testdata", "clouddriver.yml"))
	if err != nil {
		t.Fatal(err)
	}

	want, err := protoyaml.Marshal(wantC)
	if err != nil {
		t.Fatal(err)
	}

	got, err := protoyaml.Marshal(gotC)
	if err != nil {
		t.Fatal(err)
	}

	res := bytes.Compare(want, got)
	if res != 0 {
		t.Errorf("Expected generated Clouddriver config to match contents of clouddriver.yml, but got:\n" + string(got))
	}
}

func parseClouddriverConfig(fn string) (*config.Clouddriver, error) {
	dat, err := ioutil.ReadFile(fn)

	h := config.Clouddriver{}
	err = protoyaml.UnmarshalStrict([]byte(dat), &h)
	if err != nil {
		return nil, err
	}

	return &h, nil
}
