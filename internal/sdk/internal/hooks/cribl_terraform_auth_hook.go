package hooks

import (
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
)

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
		var token string
		if cachedToken, ok := o.sessions.Load(sessionKey); ok {
			token = cachedToken.(string)
			log.Printf("[DEBUG] Using cached token")
		} else {
			// Get new token
			newToken, err := o.getBearerToken(config.ClientID, config.ClientSecret, audience)
			if err != nil {
				log.Printf("[ERROR] Failed to get bearer token: %v", err)
				return req, err
			}
			token = newToken
			o.sessions.Store(sessionKey, token)
			log.Printf("[DEBUG] Got new token and cached it")
		}

		req.Header.Set("Authorization", "Bearer "+token)
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
		log.Printf("[DEBUG] Found credentials in environment variables")
		return &CriblConfig{
			ClientID:     clientID,
			ClientSecret: clientSecret,
		}, nil
	}

	// Then try ~/.cribl file
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return nil, fmt.Errorf("failed to get home directory: %v", err)
	}

	configPath := filepath.Join(homeDir, ".cribl")
	file, err := os.ReadFile(configPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read config file: %v", err)
	}

	var config CriblConfig
	if err := json.Unmarshal(file, &config); err != nil {
		return nil, fmt.Errorf("failed to parse config file: %v", err)
	}

	log.Printf("[DEBUG] Found credentials in config file")
	return &config, nil
}

func (o *CriblTerraformAuthHook) getBearerToken(clientID, clientSecret, audience string) (string, error) {
	authURL := "https://login.cribl-playground.cloud/oauth/token"
	
	// Create form data
	formData := url.Values{}
	formData.Set("grant_type", "client_credentials")
	formData.Set("client_id", clientID)
	formData.Set("client_secret", clientSecret)
	formData.Set("audience", audience)

	req, err := http.NewRequest("POST", authURL, strings.NewReader(formData.Encode()))
	if err != nil {
		return "", fmt.Errorf("failed to create request: %v", err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := o.client.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to make request: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("failed to get token: %s", string(body))
	}

	var result struct {
		AccessToken string `json:"access_token"`
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return "", fmt.Errorf("failed to parse response: %v", err)
	}

	return result.AccessToken, nil
} 