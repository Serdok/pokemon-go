# Pokémon Go

A Pokémon team management written in Go

## Endpoints

See [the public Endpoint website](https://app.swaggerhub.com/apis/estiam-go/pokemon-go_by_serdok/1.0.0) to visualize all deployed endpoints

## Public use

No installation required: the API endpoints are deployed at https://serdok-pokemon-go.ew.r.appspot.com/

### Local use

In case you want to run the API locally, a Docker Compose file has been provided. Run `docker-compose up` to build and run both the api and the firebase emulators.

You can manipulate the following values using environment variables:

- The port to use for the server: `PORT` (defaults to first `8080` if not taken)
- The mode of the server: `GIN_MODE=debug|release` (defaults to `debug`)
- The firebase JSON config: `FIREBASE_CONFIG` (either a JSON value or a file name)
- The firebase auth emulator port: `FIREBASE_AUTH_EMULATOR_HOST` (use `9099` if using  docker-compose, otherwise refer to the firebase.json file for the used port)
- The firestore emulator port: `FIRESTORE_EMULATOR_HOST` (use `8080` if using  docker-compose, otherwise refer to the firebase.json file for the used port)

## How to use

Check out the [wiki](https://github.com/Serdok/pokemon-go.wiki.git) for all the details

## Author

- [Anass 'Serdok' Lahnin](mailto:l.anass.pro@gmail.com)

