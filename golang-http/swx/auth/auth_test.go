package auth

import (
	"encoding/json"
	"function/swx"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"testing"
)

const apiUrlEnvVarName = "SWX_API_URL"

func TestTokenRevoke(t *testing.T) {

	t.Run("Successful token request and revoking", func(t *testing.T) {
		expectedToken := &Token{
			AccessToken: "valid-access-token",
			ExpiresIn:   604799,
			Scope:       "app function",
			TokenType:   "bearer",
		}

		handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			switch r.URL.Path {
			case "/oauth2/token":
				tokenBytes, err := json.Marshal(expectedToken)
				require.NoError(t, err)
				w.Header().Add("Content-Type", "application/json")
				w.WriteHeader(http.StatusOK)
				w.Write(tokenBytes)

			case "/oauth2/revoke":
				assert.Equal(t, "valid-access-token", r.FormValue("token"))
				assert.Equal(t, "client-id", r.FormValue("client_id"))
				assert.Equal(t, "client-secret", r.FormValue("client_secret"))
				w.WriteHeader(http.StatusOK)
			default:
				t.Error("unexpected url path")
			}
		})

		server := httptest.NewServer(handler)
		defer server.Close()

		os.Setenv(apiUrlEnvVarName, server.URL)

		actualToken, err := GetToken("client-id", "client-secret", []string{"app", "function"})
		require.NoError(t, err)
		assertTokenEquals(t, expectedToken, actualToken)

		err = actualToken.Revoke()
		require.NoError(t, err)
	})
}

func TestGetToken(t *testing.T) {

	t.Run("Token request successful", func(t *testing.T) {
		expectedToken := &Token{
			AccessToken: "valid-access-token",
			ExpiresIn:   604799,
			Scope:       "app function",
			TokenType:   "bearer",
		}

		handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/oauth2/token", r.URL.Path)
			tokenBytes, err := json.Marshal(expectedToken)
			require.NoError(t, err)
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write(tokenBytes)
		})

		server := httptest.NewServer(handler)
		defer server.Close()

		os.Setenv(apiUrlEnvVarName, server.URL)

		actualToken, err := GetToken("client-id", "client-secret", []string{"app", "function"})
		require.NoError(t, err)
		assertTokenEquals(t, expectedToken, actualToken)
	})

	t.Run("Token request failed due to invalid credentials", func(t *testing.T) {
		expectedErr := &swx.OAuth2Error{
			ErrorMessage:     "invalid_client",
			ErrorDescription: "Client authentication failed",
		}

		handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/oauth2/token", r.URL.Path)
			tokenBytes, err := json.Marshal(expectedErr)
			require.NoError(t, err)
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnauthorized)
			w.Write(tokenBytes)
		})

		server := httptest.NewServer(handler)
		defer server.Close()

		os.Setenv(apiUrlEnvVarName, server.URL)

		actualToken, err := GetToken("client-id", "invalid-client-secret", []string{"app", "function"})

		expectedErr.Err.Status = http.StatusUnauthorized
		require.Nil(t, actualToken)
		require.Error(t, err)
		require.Equal(t, expectedErr, err)
	})

	t.Run("Token request failed due to a server error", func(t *testing.T) {
		expectedErr := &swx.ErrorResponse{
			Err: swx.ErrorBody{
				Status:  500,
				Message: "Service currently unavailable",
			},
		}

		handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/oauth2/token", r.URL.Path)
			tokenBytes, err := json.Marshal(expectedErr)
			require.NoError(t, err)
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(tokenBytes)
		})

		server := httptest.NewServer(handler)
		defer server.Close()

		os.Setenv(apiUrlEnvVarName, server.URL)

		actualToken, err := GetToken("client-id", "invalid-client-secret", []string{"app", "function"})
		require.Nil(t, actualToken)
		require.Error(t, err)
		require.Equal(t, expectedErr, err)
	})

	t.Run("Token request failed due to network error", func(t *testing.T) {
		os.Setenv(apiUrlEnvVarName, "")

		actualToken, err := GetToken("client-id", "invalid-client-secret", []string{"app", "function"})
		require.Nil(t, actualToken)
		require.Error(t, err)
		require.IsType(t, &url.Error{}, err)
	})
}

func TestRevokeToken(t *testing.T) {

	t.Run("Successful revoke token request", func(t *testing.T) {
		handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/oauth2/revoke", r.URL.Path)
			assert.Equal(t, "some-access-token", r.FormValue("token"))
			assert.Equal(t, "client-id", r.FormValue("client_id"))
			assert.Equal(t, "client-secret", r.FormValue("client_secret"))
			w.WriteHeader(http.StatusOK)
		})

		server := httptest.NewServer(handler)
		defer server.Close()

		os.Setenv(apiUrlEnvVarName, server.URL)

		err := RevokeToken("some-access-token", "client-id", "client-secret")
		require.NoError(t, err)
	})

	t.Run("Successful revoke token request without secret (authorization code)", func(t *testing.T) {
		handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/oauth2/revoke", r.URL.Path)
			assert.Equal(t, "some-access-token", r.FormValue("token"))
			assert.Equal(t, "client-id", r.FormValue("client_id"))
			assert.Empty(t, r.FormValue("client_secret"))
			w.WriteHeader(http.StatusOK)
		})

		server := httptest.NewServer(handler)
		defer server.Close()

		os.Setenv(apiUrlEnvVarName, server.URL)

		err := RevokeToken("some-access-token", "client-id", "")
		require.NoError(t, err)
	})

	t.Run("Token request failed due to invalid credentials", func(t *testing.T) {
		expectedErr := &swx.OAuth2Error{
			ErrorMessage:     "invalid_client",
			ErrorDescription: "Client authentication failed (e.g., unknown client, no client authentication included, or unsupported authentication method)",
		}

		handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/oauth2/revoke", r.URL.Path)
			tokenBytes, err := json.Marshal(expectedErr)
			require.NoError(t, err)
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnauthorized)
			w.Write(tokenBytes)
		})

		server := httptest.NewServer(handler)
		defer server.Close()

		os.Setenv(apiUrlEnvVarName, server.URL)

		err := RevokeToken("some-access-token", "client-id", "invalid-client-secret")
		require.Error(t, err)
		require.Equal(t, expectedErr, err)
	})
}

func assertTokenEquals(t *testing.T, expected, actual *Token) {
	t.Helper()
	assert.Equal(t, expected.AccessToken, actual.AccessToken)
	assert.Equal(t, expected.ExpiresIn, actual.ExpiresIn)
	assert.Equal(t, expected.IDToken, actual.IDToken)
	assert.Equal(t, expected.RefreshToken, actual.RefreshToken)
	assert.Equal(t, expected.Scope, actual.Scope)
	assert.Equal(t, expected.TokenType, actual.TokenType)
}
