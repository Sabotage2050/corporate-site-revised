// tests/feature/forklift_api_test.go
package feature

import (
	api "corporation-site/infra/api/forklift"
	"corporation-site/tests/common/setup"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestForkliftAPI(t *testing.T) {
	tests := []struct {
		name string
		fn   func(t *testing.T)
	}{
		{"GetForkliftsByType", func(t *testing.T) {
			setupTest(t, "Forklift")
			resp, err := http.Get(setup.TestServer.URL + "/forklifts/type/battery")
			require.NoError(t, err)
			assert.Equal(t, http.StatusOK, resp.StatusCode)

			var forklifts []api.Forklift
			err = json.NewDecoder(resp.Body).Decode(&forklifts)
			require.NoError(t, err)
			assert.NotEmpty(t, forklifts)
			for _, forklift := range forklifts {
				assert.Equal(t, "battery", forklift.Enginetype)
			}
		}},

		{"GetForkliftByEngineTypeModelSerial", func(t *testing.T) {
			setupTest(t, "Forklift")
			resp, err := http.Get(setup.TestServer.URL + "/forklifts/type/battery/7FB15/29693")
			require.NoError(t, err)
			assert.Equal(t, http.StatusOK, resp.StatusCode)

			var forklift api.Forklift
			err = json.NewDecoder(resp.Body).Decode(&forklift)
			require.NoError(t, err)
			assert.Equal(t, "battery", forklift.Enginetype)
			assert.Equal(t, "7FB15", forklift.Model)
			assert.Equal(t, "29693", forklift.SerialNo)
		}},
	}

	for _, tt := range tests {
		t.Run(tt.name, tt.fn)
	}
}
