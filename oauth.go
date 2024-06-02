package provesdkservergo

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/prove-identity/prove-sdk-server-go/models/components"
)

// OAuthServerList contains the list of OAuth servers available to the SDK.
var OAuthServerList = map[string]string{
	ServerUatUs:  "https://link.uat.proveapis.com",
	ServerProdUs: "https://link.proveapis.com",
}

// OAuthClient contains the OAuth connection details.
type OAuthClient struct {
	// Client is the HTTP client.
	Client *http.Client
	// RefreshThreshold is time before the expiration to attempt a refresh.
	RefreshThreshold time.Duration

	username    string
	password    string
	accessToken string
	expiresAt   time.Time

	// Needed for international.
	clientID    string
	subClientID string

	mutex sync.RWMutex
}

// NewOAuthClientFromEnv returns an OAuth client using environment variables.
func NewOAuthClientFromEnv() (*OAuthClient, error) {
	// Load the environment variables for the OAuth credentials.
	oauthUsername, oauthPassword, clientID, subClientID, err := LoadEnv()
	if err != nil {
		return nil, err
	}

	return NewOAuthClient(oauthUsername, oauthPassword, clientID, subClientID, nil)
}

// LoadEnv returns a OAuth values from environment variables.
func LoadEnv() (oauthUsername string, oauthPassword string, clientID string, subClientID string, err error) {
	oauthUsername = os.Getenv("PROVE_USERNAME")
	if len(oauthUsername) == 0 {
		err = fmt.Errorf("missing env variable: PROVE_AUTH_USERNAME")

		return
	}

	oauthPassword = os.Getenv("PROVE_PASSWORD")
	if len(oauthPassword) == 0 {
		err = fmt.Errorf("missing env variable: PROVE_AUTH_PASSWORD")

		return
	}

	clientID = os.Getenv("PROVE_CLIENT_ID")

	subClientID = os.Getenv("PROVE_SUBCLIENT_ID")

	return
}

// NewOAuthClient returns an OAuth client.
func NewOAuthClient(
	username string,
	password string,
	clientID string,
	subClientID string,
	httpClient *http.Client,
) (*OAuthClient, error) {
	if len(username) == 0 {
		return nil, fmt.Errorf("oauth username missing")
	}
	if len(password) == 0 {
		return nil, fmt.Errorf("oauth password missing")
	}

	client := httpClient

	// Set the client.
	if httpClient == nil {
		// Set a default timeout of 10 seconds.
		client = &http.Client{
			Timeout: time.Second * 10,
		}
	}

	return &OAuthClient{
		username:         username,
		password:         password,
		clientID:         clientID,
		subClientID:      subClientID,
		Client:           client,
		RefreshThreshold: 30 * time.Second,
	}, nil
}

// OAuthTokenResponse is the response back from the OAuth service.
type OAuthTokenResponse struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int32  `json:"expires_in"`
}

// validToken returns a valid token and a boolean to determine if there is a valid token.
func (o *OAuthClient) validToken() (string, bool) {
	o.mutex.RLock()
	defer o.mutex.RUnlock()

	if o.accessToken == "" {
		return "", false
	}

	// Set to 30 seconds in the future to prevent race condition right at the expiration.
	tnow := time.Now().Add(o.RefreshThreshold)

	if tnow.After(o.expiresAt) {
		return "", false
	}

	return o.accessToken, true
}

// AccessToken returns a new or cached access token based on expiration.
func (o *OAuthClient) AccessToken(URL string) (string, error) {
	// Determine if there is an existing valid (not expired) token, if so use it.
	token, valid := o.validToken()
	if valid {
		return token, nil
	}

	o.mutex.Lock()
	defer o.mutex.Unlock()

	tokenServiceURL := fmt.Sprintf("%s/token", URL)

	data := url.Values{}
	data.Set("username", o.username)
	data.Set("password", o.password)
	data.Set("grant_type", "password")
	if len(o.clientID) > 0 {
		data.Set("client_id", o.clientID)
	}
	if len(o.subClientID) > 0 {
		data.Set("pf-subclientid", o.subClientID)
	}

	r, err := http.NewRequestWithContext(context.Background(), http.MethodPost, tokenServiceURL, strings.NewReader(data.Encode()))
	if err != nil {
		return "", err
	}

	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	r.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))

	res, err := o.Client.Do(r)
	if err != nil {
		return "", err
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		return "", fmt.Errorf("invalid response code %s", res.Status)
	}

	responseBody := new(bytes.Buffer)
	_, err = io.Copy(responseBody, res.Body)
	if err != nil {
		return "", err
	}

	decoded := &OAuthTokenResponse{}

	err = json.Unmarshal(responseBody.Bytes(), decoded)
	if err != nil {
		return "", err
	}

	o.accessToken = decoded.AccessToken
	o.expiresAt = time.Now().Add(time.Duration(decoded.ExpiresIn) * time.Second)

	return o.accessToken, nil
}

// WithAuthorization returns the access token for use with Prove OAuth. Parameter serverEnv can either be the name of
// the server or the full OAuth URL.
func WithAuthorization(oauthClient *OAuthClient, serverEnv string) func(context.Context) (components.Security, error) {
	serverURL, ok := OAuthServerList[serverEnv]
	if !ok {
		serverURL = serverEnv
	}

	return func(ctx context.Context) (components.Security, error) {
		token, err := oauthClient.AccessToken(serverURL)
		if err != nil {
			return components.Security{}, err
		}

		return components.Security{Auth: &token}, nil
	}
}
