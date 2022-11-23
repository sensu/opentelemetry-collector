// Copyright The OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//       http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package opampextension

import (
	"errors"

	"go.opentelemetry.io/collector/config"
)

// Config has the configuration for the opamp extension.
type Config struct {
	config.ExtensionSettings `mapstructure:",squash"` // squash ensures fields are correctly decoded in embedded struct

	// Endpoint is the OpAMP server URL. Transport based on the scheme of the URL.
	Endpoint string `mapstructure:"endpoint"`

	// InstanceUID is a ULID formatted as a 36 character string in canonical
	// representation. Auto-generated on start if missing.
	InstanceUID string `mapstructure:"instance_uid"`
}

// Validate checks if the extension configuration is valid
func (cfg *Config) Validate() error {
	// no need to validate less than 0 case for uint64
	if len(cfg.InstanceUID) > 36 {
		return errors.New("instance_uid has more than 36 characters")
	}
	return nil
}