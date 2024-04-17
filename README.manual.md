## Lemonway

swagger.json file was retrieved at https://sandbox-api.lemonway.fr/mb/drouot/dev/directkitrest/swagger/ui/index, more exactly: https://sandbox-api.lemonway.fr/mb/drouot/dev/directkitrest/v2/swagger

# Generate client
To generate go client and check convera api models, use:

```shell
java -jar ~/bin/openapitools/openapi-generator-cli.jar generate --global-property skipFormModel=false \
-i https://sandbox-api.lemonway.fr/mb/drouot/dev/directkitrest/v2/swagger -g go -o ./ --skip-validate-spec \
-t ./go_auth_provider_client -c ./config.json
```

## Generate template files
To generate go moustache template files, use:

````shell
java -jar ~/bin/openapitools/openapi-generator-cli.jar author template -g go
````