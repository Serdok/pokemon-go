# Pokémon Go

A Pokémon team management written in Go

## Endpoints

See [the public Endpoint website](https://endpointsportal.serdok-pokemon-go.cloud.goog/) to visualize all deployed endpoints

## Use

No installation required: the API endpoints are deployed at https://serdok-pokemon-go.ew.r.appspot.com/

### Local use

In case you want to run the API locally, a Dockerfile has been provided. Run `docker build -t serdok-pokemon-go .` to build your docker image (tag may change) and `docker run --rm -it -p 8080:8080 serdok-pokemon-go` to run the container.

The application uses Firebase for its authentication and database functionality. To run a mock database locally, you can configure a simple firebase project and get the emulators. More info [here](https://firebase.google.com/docs/emulator-suite)


## Author

- [Anass 'Serdok' Lahnin](mailto:l.anass.pro@gmail.com)

