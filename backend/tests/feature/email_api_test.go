package feature

import (
	"bytes"
	"corporation-site/domain"
	"corporation-site/tests/common/setup"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestEmailAPI(t *testing.T) {
	tests := []struct {
		name           string
		request        domain.EmailMessage
		setupRequired  bool
		expectedError  bool
		expectedStatus int
		expectedErrMsg string
	}{
		{
			name: "SendTemplatedEmail_Success",
			request: domain.EmailMessage{
				Subject:  "フォークリフトメンテナンス通知",
				TextBody: "maintenance-notification",
			},
			setupRequired:  true,
			expectedError:  false,
			expectedStatus: http.StatusOK,
		},
		{
			name: "Validation_EmptySubject",
			request: domain.EmailMessage{
				Subject:  "", // 空の件名
				TextBody: "maintenance-notification",
			},
			setupRequired:  false,
			expectedError:  true,
			expectedStatus: http.StatusBadRequest,
			expectedErrMsg: "Validation failed",
		},
		{
			name: "Validation_MissingBothBodies",
			request: domain.EmailMessage{
				Subject: "Test Subject",
				// TextBodyとHTMLBodyの両方が空
			},
			setupRequired:  false,
			expectedError:  true,
			expectedStatus: http.StatusBadRequest,
			expectedErrMsg: "Validation failed",
		},
		{
			name: "Validation_InvalidAttachment",
			request: domain.EmailMessage{
				Subject:  "Test Subject",
				TextBody: "maintenance-notification",
				Attachments: []domain.EmailAttachment{
					{
						Filename: "", // 空のファイル名
						Content:  []byte("test"),
						MIMEType: "text/plain",
					},
				},
			},
			setupRequired:  false,
			expectedError:  true,
			expectedStatus: http.StatusBadRequest,
			expectedErrMsg: "Validation failed",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.setupRequired {
				setupTest(t, "Forklift")
			}

			// テンプレートデータを追加
			reqBody := struct {
				domain.EmailMessage
				Data map[string]string `json:"data"`
			}{
				EmailMessage: tt.request,
				Data: map[string]string{
					"forkliftId":   "fork-001",
					"forkliftName": "Test Electric Forklift 1",
					"status":       "maintenance required",
				},
			}

			jsonData, err := json.Marshal(reqBody)
			require.NoError(t, err)

			resp, err := http.Post(setup.TestServer.URL+"/email/send", "application/json", bytes.NewBuffer(jsonData))
			require.NoError(t, err)
			assert.Equal(t, tt.expectedStatus, resp.StatusCode)

			if tt.expectedError {
				var errorResp struct {
					Error    string   `json:"error"`
					Messages []string `json:"messages"`
				}
				err = json.NewDecoder(resp.Body).Decode(&errorResp)
				require.NoError(t, err)
				assert.Contains(t, errorResp.Error, tt.expectedErrMsg)
			} else {
				var response domain.EmailResponse
				err = json.NewDecoder(resp.Body).Decode(&response)
				require.NoError(t, err)
				assert.NotEmpty(t, response.MessageID)
				assert.Equal(t, "sent", response.Status)
			}
		})
	}
}
