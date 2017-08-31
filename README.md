# Drawbridge API Gateway

<img align="right" src="https://media.giphy.com/media/xT5LMTMnW8KpO34mQ0/giphy.gif" width="260 "/>

Drawbridge is a lightweight API gateway written in Go.

- No datastore needed - APIs are defined by a yaml config file.

## Configuration

Drawbridge acts as a reverse proxy for one or more APIs.
The APIs can be configured in a yaml file called `config.yaml`.

Each API definition must have the following properties:
 - `name`- a friendly name for the API
 - `prefix` - the path prefix
 - `upstream_url` - the target URL to proxy to

```yaml
apis:
  reqres:
      name: "JSONPlaceholder"
      prefix: "typicode"
      upstream_url: "http://jsonplaceholder.typicode.com"
```

If the drawbridge server was running at localhost:5000, http://localhost:5000/typicode/posts would proxy to http://jsonplaceholder.typicode.com/posts.

## Development

The project includes a Docker Compose file and some useful targets in the Makefile to allow you to easily set up a
development environment.

To prepare the image for development, use `make install`.

```bash
$ make install
```

This will build the docker image and then install dependencies with the mounted volumes.

Even though Glide dependencies are installed when the Docker image is built, they must be installed again when developing because the volume mount will overwrite the /vendor directory in the container.
Notice that after running `make install`, you have a local `vendor/` directory. Please do not commit this directory to git.

Similarly, the application must be compiled again when using the development volume mounts. This is easy using `make build`.
Once compiled, the application can be run using `make start`.
These can be combined into `make build start`.
A normal development workflow might look like

```Shell
$ make install
$ make build start
$ Ctrl+c
[make some code changes]
$ make build start
$ Ctrl+c
```
Where `Ctrl+c` stops the running Docker container.

### Installing new packages

If, during development, you need to use a new package that has not been previously installed, use `make glide-get`.

```bash
$ make glide-get pkg=[package name]
```

For example, to add the `github.com/fsnotify/fsnotify` package, run `make glide-get pkg=github.com/fsnotify/fsnotify`.
