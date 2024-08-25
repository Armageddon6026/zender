# Zender

[![Golang](https://img.shields.io/badge/golang-_1.22-blue)](https://go.dev/)
[![NodeJs](https://img.shields.io/badge/node-_22.1.0-orange)](https://nodejs.org)
[![workflows](https://img.shields.io/badge/license-_MIT-greenyellow)](https://github.com/Armageddon6026/zender/blob/master/LICENSE)

<img src="web/public/logo.png" width="50px">

---
Zender is a monitoring system written in Go + Vue3, providing simple and user-friendly real-time data and status monitoring across multiple servers.


<table>
  <tr>
     <td width="50%" align="center"><b>Login</b></td>
     <td width="50%" align="center"><b>Home</b></td>
  </tr>
  <tr>
     <td><img src="assets/login.jpg"/></td>
     <td><img src="assets/home.jpg"/></td>
  </tr>
  <tr>
      <td width="50%" align="center"><b>Group</b></td>
      <td width="50%" align="center"><b>Service</b></td>
  </tr>
  <tr>
     <td><img src="assets/group.jpg"/></td>
     <td><img src="assets/service.jpg"/></td>
  </tr>
  <tr>
      <td width="50%" align="center"><b>Setting</b></td>
      <td width="50%" align="center"><b>Status Notification</b></td>
  </tr>
  <tr>
     <td><img src="assets/setting.jpg"/></td>
     <td><img src="assets/waring.jpg"/></td>
  </tr>
</table>

## Features
Server support features:
- Restful api, wirtten by [gin](https://gin-gonic.com/)
- MVC structure
- [Mariadb](https://mariadb.org/) storage, via self-made orm libirary
- Graceful shutdown
- Authentication, support jwt
- [GRPC](https://grpc.io/) supported
- SSE supported
- Containerization supported by docker / podman

Frontend support features:
- [Vue3](https://vuejs.org/) supported
- [Pinia](https://pinia.vuejs.org/)🍍 for store library
- [Unocss](https://unocss.dev/) CSS
- SPA supported by [vue-router](https://router.vuejs.org/)
- Build with [vite](https://vitejs.dev/)
- [PWA](https://developer.mozilla.org/zh-TW/docs/Web/Progressive_web_apps) supported
- SSE supported,via [fetch-event-source](https://github.com/Azure/fetch-event-source)

TODOs:
- [ ] Redis cache
- [ ] Jwt black list 
- [ ] Request rate limit
- [ ] Oath2 Authentication
- [ ] Mobile UI 
- [ ] Registry page
- [ ] document page

## Get started
Before starting, you should already install [golang](https://go.dev/), [mariadb](https://mariadb.org/) and [nodejs](https://nodejs.org/en/download/) in your develop env.
### Build & Run server

Env:
- golang (1.22 or later)
- mariadb (11 or later)

Build:
```bash
go mod download
go build
```
Before run:
- You can change the server and the database setting in `conig/server.json`
- You can build the database automatically by script in `script/database.sql`

Run locally:
```bash
./zender
```


### Test api
>See more api in https://localhost:8081/document


### Build & Run UI

Env:
- node (20 or later)

Build ui:
```bash
cd web
npm install --legacy-peer-deps
npm run build
```
Before run ui:

You sould check `web/vite.config.mts` first
>See [Vite Configuration Reference](https://vitejs.dev/config/).

Run ui with vite:
```bash
cd web
npm run dev 
```

View Web
- Explore in http://localhost:5173 for dev (use `vite` server)

- Explore in https://localhost:8081 for release (use `zender` server)

Default user account/password is `Zak/12345`


## Containerization

Zender is supported with [docker](https://docs.docker.com/engine/install/) or [podman](https://podman.io/docs/installation)

### Build & Run server in container

```bash
# build image
docker(podman) build -t zender:latest .
# run server
docker(podman) run --detach --name zenfer-server --env GIN_MODE=release  zender:latestt
```

> build database, see [Mariadb Official Image](https://hub.docker.com/_/mariadb)
