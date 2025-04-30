package hooks

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"
)

type TokenInfo struct {
	Token     string
	ExpiresAt time.Time
}

type CriblTerraformAuthHook struct {
	client   HTTPClient
	sessions sync.Map
}

type CriblConfig struct {
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
}

var (
	_ sdkInitHook       = (*CriblTerraformAuthHook)(nil)
	_ beforeRequestHook = (*CriblTerraformAuthHook)(nil)
)

func NewCriblTerraformAuthHook() *CriblTerraformAuthHook {
	log.Printf("[DEBUG] Creating new Cribl Terraform authentication hook")
	return &CriblTerraformAuthHook{}
}

func (o *CriblTerraformAuthHook) SDKInit(baseURL string, client HTTPClient) (string, HTTPClient) {
	log.Printf("[DEBUG] Initializing Cribl Terraform authentication hook with baseURL: %s", baseURL)
	o.client = client
	return baseURL, client
}

func (o *CriblTerraformAuthHook) BeforeRequest(ctx BeforeRequestContext, req *http.Request) (*http.Request, error) {
	log.Printf("[DEBUG] Cribl Terraform authentication hook: BeforeRequest for operation: %s", ctx.OperationID)

	// Check for bearer token in environment variable
	if bearerToken := os.Getenv("CRIBL_BEARER_TOKEN"); bearerToken != "" {
		log.Printf("[DEBUG] Found CRIBL_BEARER_TOKEN in environment")
		req.Header.Set("Authorization", "Bearer "+bearerToken)
		return req, nil
	}

	// Get credentials from config or environment
	config, err := o.getCredentialsFromConfig()
	if err != nil {
		log.Printf("[DEBUG] Failed to get credentials: %v", err)
		return req, nil
	}

	if config != nil {
		log.Printf("[DEBUG] Found OAuth credentials: clientID=%s", config.ClientID)
		
		// Get audience from environment or use default
		audience := "https://api.cribl-playground.cloud"
		if os.Getenv("CRIBL_AUDIENCE") != "" {
			audience = os.Getenv("CRIBL_AUDIENCE")
			log.Printf("[DEBUG] Using audience from environment: %s", audience)
		} else {
			log.Printf("[DEBUG] Using default audience: %s", audience)
		}

		// Get or create session
		sessionKey := fmt.Sprintf("%s:%s", config.ClientID, config.ClientSecret)
		var tokenInfo *TokenInfo
		
		if cachedTokenInfo, ok := o.sessions.Load(sessionKey); ok {
			tokenInfo = cachedTokenInfo.(*TokenInfo)
			// Check if token is expired or about to expire (within 60 minutes)
			if time.Until(tokenInfo.ExpiresAt) < 60*time.Minute {
				log.Printf("[DEBUG] Token expired or about to expire, fetching new token")
				tokenInfo = nil
			} else {
				log.Printf("[DEBUG] Using cached token, expires at: %v", tokenInfo.ExpiresAt)
			}
		}

		if tokenInfo == nil {
			// Get new token
			newTokenInfo, err := o.getBearerToken(ctx.Context, config.ClientID, config.ClientSecret, audience)
			if err != nil {
				log.Printf("[ERROR] Failed to get bearer token: %v", err)
				return req, err
			}
			tokenInfo = newTokenInfo
			o.sessions.Store(sessionKey, tokenInfo)
			log.Printf("[DEBUG] Got new token and cached it, expires at: %v", tokenInfo.ExpiresAt)
		}

		req.Header.Set("Authorization", "Bearer "+tokenInfo.Token)
	} else {
		log.Printf("[DEBUG] No OAuth credentials found")
	}

	return req, nil
}

func (o *CriblTerraformAuthHook) getCredentialsFromConfig() (*CriblConfig, error) {
	// First try environment variables
	clientID := os.Getenv("CRIBL_CLIENT_ID")
	clientSecret := os.Getenv("CRIBL_CLIENT_SECRET")

	if clientID != "" && clientSecret != "" {
		return &CriblConfig{
			ClientID:     clientID,
			ClientSecret: clientSecret,
		}, nil
	}

	// Then try ~/.cribl/credentials file
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Printf("[ERROR] Failed to get home directory: %v", err)
		return nil, fmt.Errorf("failed to get home directory: %v", err)
	}

	// Check for credentials in ~/.cribl/credentials
	configDir := filepath.Join(homeDir, ".cribl")
	configPath := filepath.Join(configDir, "credentials")
	
	file, err := os.ReadFile(configPath)
	if err != nil {
		// Fallback to legacy ~/.cribl file if credentials file doesn't exist
		legacyPath := filepath.Join(homeDir, ".cribl")
		file, err = os.ReadFile(legacyPath)
		if err != nil {
			log.Printf("[ERROR] Failed to read both credentials files: %v", err)
			return nil, fmt.Errorf("failed to read credentials file: %v", err)
		}
	}

	var config CriblConfig
	if err := json.Unmarshal(file, &config); err != nil {
		log.Printf("[ERROR] Failed to parse config file: %v", err)
		return nil, fmt.Errorf("failed to parse config file: %v", err)
	}

	return &config, nil
}

func (o *CriblTerraformAuthHook) getBearerToken(ctx context.Context, clientID, clientSecret, audience string) (*TokenInfo, error) {
	authURL := "https://login.cribl-playground.cloud/oauth/token"
	
	// Create form data
	formData := url.Values{}
	formData.Set("grant_type", "client_credentials")
	formData.Set("client_id", clientID)
	formData.Set("client_secret", clientSecret)
	formData.Set("audience", audience)

	req, err := http.NewRequestWithContext(ctx, "POST", authURL, strings.NewReader(formData.Encode()))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %v", err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := o.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to make request: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to get token: %s", string(body))
	}

	var result struct {
		AccessToken string `json:"access_token"`
		ExpiresIn   int    `json:"expires_in"`
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("failed to parse response: %v", err)
	}

	// Calculate expiration time
	expiresAt := time.Now().Add(time.Duration(result.ExpiresIn) * time.Second)

	return &TokenInfo{
		Token:     result.AccessToken,
		ExpiresAt: expiresAt,
	}, nil
} 