# GoProfile

GoProfile is an implementation to migrate data between Users and Profiles identifying them by email.

## Requirements

* [Go](https://golang.org/dl/)
* Docker
* Dump (coatlicue)
* Dump (cuco-restapi-feeds)

## Dependencies

Install `cuco-restapi-feeds` project

```sh
git clone git@github.com:Cultura-Colectiva-Tech/cuco-restapi-feeds.git

cd cuco-restapi-feeds

copy dump in /data folder

cuco-restapi-feeds: docker-compose up
````
## Config coatlicue
Start project
```sh
docker coatlicue start
```
Import DB 
```sh
docker exec -it mictlanDB bash

mongorestore -u ${user} -p ${pass} --authenticationDatabase admin  --db ${path}
```

Login in CMS `https://dev.cms.culturacolectiva.com`

Update your ROLE to `ROLE_ADMIN`

```sh
db.profiles.updateMany({email: '<email>'}, {$set: {status: 'STATUS_ACTIVE', role: 'ROLE_ADMIN'}})
```

## Install GoProfile & Test
```sh
  git clone git@github.com:Cultura-Colectiva-Tech/go-profile.git

  cd go-profile/profile

  go test ./... -v -cover --token ${token}
```

## Basic usage

```sh
  go run main.go --user-email ${email} --profile-email ${email} --token ${token}
```

## Flags

| Flag            | Type   | Description                                            |
|-----------------|--------|--------------------------------------------------------|
|`--token`        | string |Token to make petitions get from CMS                    |
|`--user-email`   | string |User email from last database (CMS)                     |
|`--profile-email` | string |Profile email from new database (CMS)                    |
|`--environment`  | string |Environment to make petitions {dev,staging, prod}       |
|`--data-users`   | string |URL for get users                                       |
