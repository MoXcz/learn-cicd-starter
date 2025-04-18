package auth_test

import (
	"github.com/bootdotdev/learn-cicd-starter/internal/auth"
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		headers http.Header
		want    string
		wantErr bool
	}{
		{
			name: "missing header",
			headers: http.Header{
				"Authorization": []string{""}},
			wantErr: true,
			want:    "",
		},
		{
			name: "malformed header",
			headers: http.Header{
				"Authorization": []string{"Key value"}},
			wantErr: true,
			want:    "",
		},
		{
			name: "valid header",
			headers: http.Header{
				"Authorization": []string{"ApiKey valid_ke"}},
			wantErr: false,
			want:    "valid_key",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, gotErr := auth.GetAPIKey(tt.headers)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("GetAPIKey() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("GetAPIKey() succeeded unexpectedly")
			}
			// TODO: update the condition below to compare got with tt.want.
			if got != tt.want {
				t.Errorf("GetAPIKey() = %v, want %v", got, tt.want)
			}
		})
	}
}
