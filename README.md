# baseline-golang-app


## 1. Build the Docker images

```bash
docker build --target development -t baseline-golang-app:dev -f Dockerfile . --progress=plain
docker build --target production  -t baseline-golang-app:prod -f Dockerfile . --progress=plain
```

## 2. Run the Docker images

```bash
docker run -d --name baseline-golang-app-dev -p 1323:1323 baseline-golang-app:dev
docker run -d --name baseline-golang-app-prod -p 1323:1323 baseline-golang-app:prod
```

## 3. Test the application

```bash
curl http://localhost:1323
```

