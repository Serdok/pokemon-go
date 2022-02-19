# Pokémon Go

A Pokémon team management written in Go

## Endpoints

See [the public Endpoint website](https://endpointsportal.serdok-pokemon-go.cloud.goog/) to visualize all deployed endpoints

## Use

No installation required: the API endpoints are deployed at https://serdok-pokemon-go.ew.r.appspot.com/

### Local use

In case you want to run the API locally, a Docker Compose file has been provided. Run `docker-compose up` to build and run both the api and the firebase emulators.

You can manipulate the following values using environment variables:

- The port to use for the server: `PORT=8080` (defaults to first `8080` if not taken)
- The mode of the server: `GIN_MODE=debug|release` (defaults to `debug`)
- The firebase JSON config: `FIREBASE_CONFIG` (either a JSON value or a file name)
- The firebase auth emulator port: `FIREBASE_AUTH_EMULATOR_HOST` (refer to the firebase.json file for the used auth port)


## Author

- [Anass 'Serdok' Lahnin](mailto:l.anass.pro@gmail.com)

