# Cortex API client 
A Cortex API client writen in Go.

### Requisites
1. `Go >= 1.20.0`.
2. A valid Cortex API Bearer Token (see [api-keys](https://app.getcortexapp.com/admin/settings/api-keys)).

### Setup
The client is expecting a file `config.env` in the root folder of this repo. Create that file with the content below, and update the information accordingly:
```bash
BASE_URL = "https://api.getcortexapp.com/api/v1"
API_KEY = "<YOUR-CORTEX-API-KEY>"
```

---
#### More info:
* The layout of the project is based on the official [modules layout](https://go.dev/doc/modules/layout#package-or-command-with-supporting-packages) for packages or commands with supporting packages.
