# Go random logger
Allow to send random logs to `/var/log/random.log` file

## I. Prerequisites
**Docker** installed on your system

## II. Generating personal access token
In order to authenticate to GitHub Packages the first thing we'll need is an access token. Open your GitHub account, go to Settings -> Developer Settings -> [Personal access tokens](https://github.com/settings/tokens), click **Generate new Token**, give it a name, select **write:packages** and save

## III. Setting environment variables
- Define a `.env` file wih Github username and access token like below:
```text
GITHUB_USERNAME=veben
GITHUB_TOKEN=******
```

## IV. Building the docker image
- Open `docker-compose.yaml` file
- Modify the tag (`0.0.1` by default)
- Build using docker compose:
```sh
docker-compose build
```

## V. Pushing the docker image
- Source `.env` file
```sh
source .env
```
- Log to docker package
```sh
echo $GITHUB_TOKEN | docker login ghcr.io -u $GITHUB_USERNAME --password-stdin
```
- Push using docker compose:
```sh
docker-compose push
```

## VI. Testing
- Run using docker compose (`-d` for detached mode):
```sh
docker-compose up -d
```
- Check that the logs are created in the container, in the file `/var/log/random.log`
```sh
docker exec -it `docker container ls | grep go_random_log | awk '{print $1}'` bash -c "tail -f /var/log/random.log"
```
- Something like below should be displayed:
```plaintext
2024/04/27 21:53:14 [ERROR] Random log message number three
2024/04/27 21:53:17 [INFO] Random log message number three
2024/04/27 21:53:20 [ERROR] This is a random log message
2024/04/27 21:53:21 [ERROR] Another random log message
2024/04/27 21:53:22 [INFO] This is a random log message
2024/04/27 21:53:24 [ERROR] This is a random log message
2024/04/27 21:53:25 [INFO] This is a random log message
```
- Check that the logs are created in the host:
```sh
docker exec -it `docker container ls | grep go_random_log | awk '{print $1}'` bash -c "tail -f /var/log/random.log"
```