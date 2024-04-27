# Go random logger
Allow to send random logs to `/var/log/random.log` file

## I. Generating personal access token
In order to authenticate to GitHub Packages the first thing we'll need is an access token. Open your GitHub account, go to Settings -> Developer Settings -> [Personal access tokens](https://github.com/settings/tokens), click **Generate new Token**, give it a name, select **write:packages** and save

## II. Building the docker image
- Open `docker-compose.yaml` file
- Modify the tag (`0.0.1` by default)
- Build using docker compose:
```sh
docker-compose build
```

## III. Pushing the docker image
- Save the token in an envirnment variable:
```sh
export GITHUB_TOKEN=***
```
- Log to docker pacakge
```sh
echo $GITHUB_TOKEN | docker login https://docker.pkg.github.com -u veben --password-stdin
```
- Push using docker compose:
```sh
docker-compose build
```

## IV. Testing
TODO
