# Go random logger
A tool for sending random logs to the `/var/log/random.log` file.

## I. Prerequisites
- **Docker** must be installed on your system.
- **Docker Compose** must be installed on your system.

## II. Testing
- Run the container using Docker Compose (`-d` for detached mode):
```sh
docker-compose up -d
```
- Verify that the logs are being created in the container at `/var/log/random.log`
```sh
docker exec -it `docker container ls | grep go_random_log | awk '{print $1}'` bash -c "tail -f /var/log/random.log"
```
- You should see output similar to the following:
```plaintext
2024/06/05 21:50:32 [INFO] The spider ventures into a burrow boldly.
2024/06/05 21:50:35 [INFO] A whale swirls in a boulder hurriedly.
2024/06/05 21:50:36 [ERROR] A dolphin discovers a cliff loudly.
2024/06/05 21:50:37 [ERROR] A seagull runs to a tree recklessly.
2024/06/05 21:50:38 [ERROR] The seal dives down a cave earnestly.
2024/06/05 21:50:41 [INFO] The seal catches the tide patiently.
2024/06/05 21:50:44 [WARNING] The bear barks at the stars nonchalantly.
2024/06/05 21:50:47 [WARNING] The rabbit catches the rock suddenly.
2024/06/05 21:50:49 [ERROR] A squirrel sneaks up on the sand briskly.
2024/06/05 21:50:51 [WARNING] The panda swirls in the lagoon relentlessly.
```

## III. Creating a new version of the logger
Follow these steps to create a new version of the logger:

### 1. Generating a Personal Access Token
In order to authenticate to GitHub Packages the first thing we'll need is an access token.

- Open your GitHub account, go to Settings -> Developer Settings -> [Personal access tokens](https://github.com/settings/tokens)
- Click **Generate new Token**
- Give it a name, select **write:packages** and save the token
- Define a `.env` file with your GitHub username and access token:
```text
GITHUB_USERNAME=<your_username>
GITHUB_TOKEN=<your_token>
```

### 2. Evolving the Logger
- Modify the source code as needed
- Update the tag version in the `.env` file, replacing `<version>` with your new version:
```text
TAG_VERSION=<version>
```
- Update the `CHANGELOG.md` to describe the changes in the new version

### 3. Creating a new tag
- Create and push a new tag:
```sh
git tag <version>
git push origin <version>
```

### 4. Building the docker image
- Build using docker compose:
```sh
docker-compose build
```

### 5. Pushing the docker image
- Load the environment variables from the `.env` file
```sh
source .env
```
- Log in to the GitHub Container Registry
```sh
echo $GITHUB_TOKEN | docker login ghcr.io -u $GITHUB_USERNAME --password-stdin
```
- Push the Docker image using Docker Compose:
```sh
docker-compose push
```