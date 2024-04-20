## Pixelated Pipeline

The backend infrastructure of a **Pixelated Pen** written using GoLang.
Frontend is written using Svelte Kit (Pixelated Pen).

## Architecture

![Pixelated Pipeline Architecture][system_design]

[system_design]: ./lld.excalidraw.png

## Notes

1. To view the database using `Adminer` UI, use the name of the service in `docker-compose.yml` file as server.
2. After setting up the project locally, we need to create a file named `application.yml` in the root directory which contains the configurations for the application to be deployed. To use intellisense while writing this file, install [YAML](https://marketplace.visualstudio.com/items?itemName=redhat.vscode-yaml) and add the following line on top of newly created `application.yml`.

```yaml
# yaml-language-server: $schema=application.schema.json
```
3. The project uses `make` to build the project inside a `/build` directory. There is no need to manually invoke the go compiler. Similarly, this project contains a `docker-compose.yml` file which can be used to spin up the `postgresql` database required.

    **TODO:** Work has been ongoing for creating a Dockerfile which can be used to run the project using a single command and will not require the use of `make`. That will lead to easier deployments.

4. For the developers of the software, if provided, read the `doc.go` under whichever package you are using to understand potential implementation details and pitfalls to look for.