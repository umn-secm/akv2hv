# akv2hv
Go app for migrating secrets from Azure KeyVault to Hashicorp Vault.

## Prerequisites

1. Download the latest akv2hv binary for your OS from https://github.com/umn-secm/akv2hv/releases or [build locally](./README.md#building-locally) and open a command line window in the directory that you downloaded the binary to.

2. Install azure cli <https://learn.microsoft.com/en-us/cli/azure/install-azure-cli>

3. Login to azure cli `az login`

4. Get a vault token with permissions to write secrets

    - If you have the vault enterprise cli installed, run `vault login --method=saml --namespace=admin`

    - If you do not have the vault enterprise cli installed login to the vault GUI and go to the `Person Icon>Copy token` and then set your environmental variable `VAULT_TOKEN` with that token.

5. Export your vault token `export VAULT_TOKEN=TOKEN_FROM_STEP_3`

## Copy Secrets

1. The first step is to generate a json file with the list of all secrets you have in KeyVault. This will only retrieve their names, not their values.

    ```bash
    # Linux
    ./akv2hv --kv=INSERT_AZ_KV_NAME --gen

    # Windows
    akv2hv.exe --kv=INSERT_AZ_KV_NAME --gen
    ```

2. The second step is to edit the secrets.json file that was generated in step 1. The fields that you will want to edit include:

    - `vault_secret_mount`      - the kvv2 mount location (defaults to secret)
    - `vault_secret_name` 	  - the name of the secret including path that you would like to store the value to i.e. github/super_secret
    - `vault_secret_key`      - the key of the field within the secret (each secret can contain multiple key/value pairs)
    - `copy`                  - true if you would like it copied to vault (defaults to false so nothing will be copied)

3. The final step is to run the copy function to retrieve the secret data from Azure KeyVault and write the secrets to Hashicorp Vault.

    ```bash
    # Linux
    ./akv2hv --kv=INSERT_AZ_KV_NAME --vault_addr=https://EXAMPLE.z1.hashicorp.cloud:8200/ --vault_namespace=admin/namespace --copy

    # Windows
    akv2hv.exe --kv=INSERT_AZ_KV_NAME --vault_addr=https://EXAMPLE.z1.hashicorp.cloud:8200/ --vault_namespace=admin/namespace --copy
    ```

## Flags

``` bash
  -copy
        Run the function to copy the secrets from KeyVault to HashiCorp Vault based on the secrets.json locations.
  -file string
        json file to write or read list of secrets from/to. Defaults to secrets.json in the current directory
  -gen
        Generate json file secrets.json with a list of secrets from KeyVault as keys.
  -kv string
        The name of the Azure Key Vault.
  -mount string
        The path of the kvv2 mount (will default to secret).
  -vault_addr string
        The url for vault (i.e. https://examplevault.com).
  -vault_namespace string
        The namespace for vault (i.e. https://examplevault.com).
```


## Building Locally

```bash
git clone git@github.com:umn-secm/akv2hv.git
cd akv2hv
go build -o .
```
