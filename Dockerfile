# Development stage
FROM public.ecr.aws/docker/library/golang:1.19 AS development

WORKDIR /app

COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download \
    && go install github.com/cosmtrek/air@latest \
    && go mod verify

COPY . .

RUN go build -o ./bin/main 


ENV WAIT_VERSION 2.7.2
ADD https://github.com/ufoscout/docker-compose-wait/releases/download/$WAIT_VERSION/wait /wait
RUN chmod +x /wait

EXPOSE 1323

# CMD ["./bin/main"]
ENTRYPOINT ["air", "run", "."]


# Production stage
FROM public.ecr.aws/docker/library/golang:1.19 AS production

WORKDIR /app

# copy only the built artifact and dependencies from the development stage
COPY --from=development /app/bin/main ./bin/main

EXPOSE 1323

CMD ["./bin/main"]