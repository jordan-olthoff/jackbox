FROM node:lts as frontend-build
WORKDIR /
COPY ./web/app/package*.json ./
RUN npm install --force
COPY ./web/app .
RUN npm run build


FROM golang:1.20 as backend-build
WORKDIR /
COPY go.mod ./
RUN go mod download
COPY .. ./
COPY --from=frontend-build /dist /cmd/user/frontend/dist
RUN make prepare_backend
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build ./cmd/user/main.go


FROM alpine:3.18.2
RUN apk add --no-cache bash
COPY --from=backend-build /main /
ENTRYPOINT ["/main"]