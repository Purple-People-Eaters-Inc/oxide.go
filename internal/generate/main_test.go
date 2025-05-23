// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_loadAPI(t *testing.T) {
	type args struct {
		file string
	}
	tests := []struct {
		name    string
		args    args
		wantErr string
	}{
		{
			name:    "file does not exist",
			args:    args{"bob.txt"},
			wantErr: "no such file or directory",
		},
		{
			name:    "invalid version",
			args:    args{"generate/test_utils/INVALID_VERSION"},
			wantErr: "error loading openAPI spec from \"https://raw.githubusercontent.com/oxidecomputer/omicron//openapi/nexus.json\"",
		},
		{
			name: "success",
			args: args{"../VERSION_OMICRON"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := loadAPIFromFile(tt.args.file)
			if err != nil {
				assert.ErrorContains(t, err, tt.wantErr)
				return
			}
			assert.NotNil(t, got)
		})
	}
}
