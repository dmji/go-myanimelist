package mal_client

import (
	"context"
	"crypto/rand"
	"fmt"
	"net/http"

	"golang.org/x/oauth2"
)

type (
	fnLoadCache  func() (*oauth2.Token, error)
	fnStoreCache func(oauth2.Token) error
)

type OAuth2CacheClientProvider struct {
	conf    *oauth2.Config
	fnLoad  fnLoadCache
	fnStore fnStoreCache
}

func NewOAuth2Cacher(conf *oauth2.Config, fnLoad fnLoadCache, fnStore fnStoreCache) *OAuth2CacheClientProvider {
	//	"auth-example-token-cache.txt"
	return &OAuth2CacheClientProvider{
		conf:    conf,
		fnLoad:  fnLoad,
		fnStore: fnStore,
	}
}

/*
func loadCachedToken() (*oauth2.Token, error) {
	b, err := os.ReadFile(cacheName)
	if err != nil {
		return nil, fmt.Errorf("reading oauth2 token from cache file %q: %v", cacheName, err)
	}
	token := new(oauth2.Token)
	if err := json.Unmarshal(b, token); err != nil {
		return nil, fmt.Errorf("unmarshaling oauth2 token: %v", err)
	}
	return token, nil
}

func cacheToken(token oauth2.Token) error {
	b, err := json.MarshalIndent(token, "", "   ")
	if err != nil {
		return fmt.Errorf("marshaling token %s: %v", token.AccessToken, err)
	}
	err = os.WriteFile(cacheName, b, 0o644)
	if err != nil {
		return fmt.Errorf("writing token %s to file %q: %v", token.AccessToken, cacheName, err)
	}
	return nil
} */

func (o *OAuth2CacheClientProvider) RestoreClientFromCached(ctx context.Context) (*http.Client, error) {
	oauth2Token, err := o.fnLoad()
	if err != nil {
		return nil, nil
	}

	refreshedToken, err := o.conf.TokenSource(ctx, oauth2Token).Token()

	if err == nil && (oauth2Token != refreshedToken) {

		// Caching refreshed oauth2 token...
		if err := o.fnStore(*refreshedToken); err != nil {
			return nil, fmt.Errorf("caching refreshed oauth2 token: %s", err)
		}
		return o.conf.Client(ctx, refreshedToken), nil

	}

	return o.conf.Client(ctx, oauth2Token), nil
}

func (o *OAuth2CacheClientProvider) RequestRedirectToMalOAuth2(ctx context.Context, conf *oauth2.Config, state string) (*http.Client, string, error) {
	codeVerifier, err := newCodeVerifier()
	if err != nil {
		return nil, "", err
	}

	// Produce the authentication URL where the user needs to be redirected and
	// allow your application to access their MyAnimeList data.
	authURL := conf.AuthCodeURL(state,
		oauth2.SetAuthURLParam("code_challenge", codeVerifier),
	)
	return nil, authURL, nil
}

func (o *OAuth2CacheClientProvider) CatchRedirect(ctx context.Context, conf *oauth2.Config, code string) (*http.Client, error) {
	codeVerifier, err := newCodeVerifier()
	if err != nil {
		return nil, err
	}

	// Exchange the authentication code for a token. MyAnimeList currently only
	// supports the plain code_challenge_method so to verify the string, just
	// make sure it is the same as the one you entered in the code_challenge.
	token, err := conf.Exchange(ctx, code,
		oauth2.SetAuthURLParam("code_verifier", codeVerifier),
	)
	if err != nil {
		return nil, fmt.Errorf("exchanging code for token: %v", err)
	}

	// Authentication was successful. Caching oauth2 token...
	if err := o.fnStore(*token); err != nil {
		return nil, fmt.Errorf("caching oauth2 token: %s", err)
	}

	return conf.Client(ctx, token), nil
}

// Generate a code verifier, a high-entropy cryptographic random string. It
// will be set as the code_challenge in the authentication URL. It should
// have a minimum length of 43 characters and a maximum length of 128
// characters.
func newCodeVerifier() (string, error) {
	const codeVerifierLength = 128
	codeVerifier, err := generateCodeVerifier(codeVerifierLength)
	if err != nil {
		return "", fmt.Errorf("generating code verifier: %v", err)
	}
	return codeVerifier, nil
}

func generateCodeVerifier(length int) (string, error) {
	const charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZ" +
		"abcdefghijklmnopqrstvuwxyz" +
		"0123456789-._~"
	bytes := make([]byte, length)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	for i, b := range bytes {
		bytes[i] = charset[b%byte(len(charset))]
	}
	return string(bytes), nil
}

// Prepare the oauth2 configuration with your application ID, secret, the
// MyAnimeList authentication and token URLs as specified in:
//
// https://myanimelist.net/apiconfig/references/authorization
func NewConfigOAuth2(clientID string, clientSecret string) *oauth2.Config {
	return &oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		Endpoint: oauth2.Endpoint{
			AuthURL:   "https://myanimelist.net/v1/oauth2/authorize",
			TokenURL:  "https://myanimelist.net/v1/oauth2/token",
			AuthStyle: oauth2.AuthStyleInParams,
		},
	}
}
