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

Generate execute
```sh
go build
```

Run `profiles` task

* Search by email
```sh
./go-ptofile --task profiles --user-email ${email} --profile-email ${email} --token ${token}
```
* Search by slug
```sh
./go-ptofile --task profiles --user-slug ${email} --profile-email ${email} --token ${token}
```

Run `articles` task and filter by authorId

```sh
./go-ptofile --task articles --start-date ${start} --end-date ${end} --author-id ${hash} --token ${token}
```

Run `articles` task and filter by authorSlug

```sh
./go-ptofile --task articles --start-date ${start} --end-date ${end} --author-slug ${slug} --token ${token}
```
## Flags

| Flag            | Type   | Description                                            |
|-----------------|--------|--------------------------------------------------------|
|`--token`        | string | Token to make petitions get from CMS                   |
|`--user-email`   | string | User email from last database (CMS)                    |
|`--user-slug`    | string | User slug from last database (CMS)                     |
|`--profile-email` | string | Profile email from new database (CMS)                   |
|`--environment`  | string | Environment to make petitions {dev,staging, prod}      |
|`--data-users`   | string | URL for get users                                      |
|`--limit`        | string | Limit paginate, Default: 50                            |
|`--page`         | string | Page paginate, Default: Default: 1                     |
|`--type-post`    | string | Article type to search, Default: POST                  |
|`--status-post`  | string | Article status to search, Default: STATUS_PUBLISHED    |
|`--start-date`   | string | StartDate to filter article, Default: 2018-01-01        |
|`--end-date`     | string | EndDate to filter article, Default: 2018-12-31          |
|`--author-slug`  | string | Search articles by author slug                         |
|`--author-id`    | string | Search articles by author id                           |
