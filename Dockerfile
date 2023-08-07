# Development stage
FROM public.ecr.aws/docker/library/golang:1.18-alpine AS development

RUN apk add --no-cache git

WORKDIR /app

COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download \
    && go mod verify

COPY . .

RUN go build -o ./bin/main  ./src/simple-app

EXPOSE 1323

CMD ["./bin/main"]


# Production stage
FROM public.ecr.aws/docker/library/golang:1.18-alpine AS production

WORKDIR /app

RUN pwd && ls -al

# copy only the built artifact and dependencies from the development stage
COPY --from=development /app/bin/main ./bin/main

EXPOSE 1323

CMD ["./bin/main"]