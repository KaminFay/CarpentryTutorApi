# Vault Setup:

- Pull the container:

```bash
➜  ~ docker pull vault
```

- Run the container, revealing the port as well:

```bash
➜  ~ docker run --cap-add=IPC_LOCK -d --name=dev-vault -p 8200:8200 -e 'VAULT_DEV_ROOT_TOKEN_ID=myroot' -e 'VAULT_DEV_LISTEN_ADDRESS=127.0.0.1:8200' vault
3972146ee21a2605a30db777667ee3e9d29aa41d0c136ac7f729720f7181973d
```

- Set the aliases and export the address for vault:

```bash
➜  ~ alias vault='docker exec -it -e VAULT_ADDR dev-vault vault "$@"'
➜  ~ export VAULT_ADDR=http://127.0.0.1:8200
```

