package azure

import (
	"context"
	"sync"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
)

// cachedToken struct holds information about the token and
// the expires time. This is useful so the token can be reused
// instead of fetching it again from azure (avoiding the rate
// limiting implemented by azure)
type cachedToken struct {
	token     string
	expiresOn time.Time
}

var (
	// tokenCache maps the clientId into the cachedToken struct.
	// clientId is unique
	tokenCache   = map[string]cachedToken{}
	tokenCacheMu sync.RWMutex
)

func GetBearerToken(clientId, clientSecret, tenantId string) (string, error) {
	tokenCacheMu.RLock()

	if t, ok := tokenCache[clientId]; ok && time.Now().Before(t.expiresOn.Add(-5*time.Minute)) {
		tokenCacheMu.RUnlock()

		return t.token, nil
	}
	tokenCacheMu.RUnlock()

	cred, err := azidentity.NewClientSecretCredential(tenantId, clientId, clientSecret, nil)
	if err != nil {
		return "", err
	}

	opts := policy.TokenRequestOptions{
		Scopes: []string{"https://management.azure.com/.default"},
	}

	tokenResponse, err := cred.GetToken(context.Background(), opts)
	if err != nil {
		return "", err
	}

	tokenCacheMu.Lock()
	tokenCache[clientId] = cachedToken{
		token:     tokenResponse.Token,
		expiresOn: tokenResponse.ExpiresOn,
	}
	tokenCacheMu.Unlock()

	return tokenResponse.Token, nil
}
