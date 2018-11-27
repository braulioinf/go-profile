# GoProfile

GoProfile is an implementation to migrate data between Users and Profiles identifying them by email.

## Basic usage

```sh
  go run main.go --email-user ${email} --email-profile ${email} --token ${token}
```

## Flags

| Flag            | Type   | Description                        |
|-----------------|--------|------------------------------------|
|`--token`        | string |Token to make petitions|
|`--email-user`   | string |Email user from last database (CMS) |
|`--email-profile` | string |Email profile from new dabatase (CMS)|
|`--env`          | string |Environment to make petitions       |
|`--data-users`   | string |URL for get users                   |


## Testing

* Go to `/profile` module

```sh
  go test ./... -v -cover --token ${token}
```
