package backends

import (
	"strings"
	"encoding/json"
	"github.com/cyberark/conjur-api-go/conjurapi/response"
    "github.com/cyberark/conjur-api-go/conjurapi"
	"strconv"
)

// Vault is a struct for working with a Vault backend
type ConjurVault struct {
	client *conjurapi.Client
}

// NewVaultBackend initializes a new Vault Backend
func NewConjurVaultBackend(client *conjurapi.Client) *ConjurVault {
	return &ConjurVault{
		client: client,
	}
}

// Login authenticates with the auth type provided
func (v *ConjurVault) Login() error {
	return nil
}

// GetSecrets gets secrets from vault and returns the formatted data
func (v *ConjurVault) GetSecrets(path string, version string, annotations map[string]string) (map[string]interface{}, error) {

	// Retrieve a secret into []byte.
	var filter conjurapi.ResourceFilter
	filter.Kind = "variable"
	filter.Search = string(path)

	req, err := v.client.ResourcesRequest(&filter)
	if err != nil {
		return nil, err
	}
	resp, err := v.client.SubmitRequest(req)
	if err != nil {
		return nil, err
	}
	data, err := response.DataResponse(resp)
	if err != nil {
		return nil, err
	}
	result := make(map[string]interface{})
	if data != nil && string(data) != "[]" {
		var secretsFound []map[string]string

		json.Unmarshal(data, &secretsFound)

		for s := range secretsFound {
			splits := strings.Split(secretsFound[s]["id"], ":")
			secret_slices := strings.Split(splits[2], "/")
			secretValue, err := v.client.RetrieveSecret(splits[2])
			if err != nil {
				panic(err)
			}
			result[secret_slices[len(secret_slices)-1]] = string(secretValue)
		}
	}


	return result, nil
}

// GetIndividualSecret will get the specific secret (placeholder) from the SM backend
// For Vault, we only support placeholders replaced from the k/v pairs of a secret which cannot be individually addressed
// So, we use GetSecrets and extract the specific placeholder we want
func (v *ConjurVault) GetIndividualSecret(kvpath, secret, version string, annotations map[string]string) (interface{}, error) {
	if version != ""{
		numberVersion, err := strconv.Atoi(version)
		if err != nil {
			return nil, err
		}
		secretValue, err := v.client.RetrieveSecretWithVersion(kvpath+"/"+secret,numberVersion)
		if err != nil {
			return nil, err
		}
		return string(secretValue), nil
	}else{
		secretValue, err := v.client.RetrieveSecret(kvpath+"/"+secret)
		if err != nil {
			return nil, err
		}
		return string(secretValue), nil
	}
   
}
