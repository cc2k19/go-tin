# go-tin
[![Build Status](https://travis-ci.org/cc2k19/go-tin.svg?branch=master)](https://travis-ci.org/cc2k19/go-tin)
[![Coverage Status](https://coveralls.io/repos/github/cc2k19/go-tin/badge.svg?branch=master)](https://coveralls.io/github/cc2k19/go-tin?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/cc2k19/go-tin)](https://goreportcard.com/report/github.com/cc2k19/go-tin)
[![GoDoc](https://godoc.org/github.com/cc2k19/go-tin?status.svg)](https://godoc.org/github.com/cc2k19/go-tin)
[![Licence](https://img.shields.io/github/license/cc2k19/go-tin.svg?style=popout)](https://github.com/cc2k19/go-tin/blob/master/LICENSE)


# Table of Contents

  - [Overview](#overview)
  - [API](#api)
      - [Users](#users)
        - [Register User](#register-user)
        - [Retrieve User](#get-user)
        - [Follow User](#follow-user)
        - [Unfollow User](#unfollow-user)
        - [Retrieve Followers](#retrieve-followers)
        - [Retrieve Following](#retrieve-following)
      - [Posts](#posts)
        - [Create Post](#create-post)
        - [Retrieve Posts](#retrieve-posts)
  - [Database Model](#database-model)


## Overview

Go-tin is new, one of a kind social network developed as a project for FMI course Clean Code in 2019.
The project is currently in active development, for up-to-date status you can visit the `dev` branch.

## API

## Users

### Register User

#### Request

#### Route

`POST /v1/users`

#### Body

```json
{
	"username": "test_user",
	"password": "a"
}
```

#### Response

| Status Code | Description |
| --- | --- |
| 201 Created |  Will be returned if the user is successfully registered. |
| 400 Bad Request | Will be returned if the request is malformed or missing mandatory data. |


### Get User

#### Request

#### Route

`GET /users/{username}`

#### *Requires Basic Authentication*

#### Response

| Status Code | Description |
| --- | --- |
| 200 OK | Will be returned if the user is found. Expected body is below. |
| 401 Unauthorized | Will be returned if authentication fails. |
| 404 Not Found | Will be returned if the user does not exist. |

#### Body

```json
{
  "username": "test_user",
  "birth_date": "1997-12-12",
  "bio": "test bio",
  "hometown": "test town"
}
```

### Follow User

#### Request

#### Route

`PUT /follow/{username}`

#### *Requires Basic Authentication*

#### Response

| Status Code | Description |
| --- | --- |
| 201 Created | Will be returned if the relation is created. |
| 401 Unauthorized | Will be returned if authentication fails. |
| 404 Not Found | Will be returned if the user does not exist. |


### Unfollow User

#### Request

#### Route

`DELETE /follow/{username}`

#### *Requires Basic Authentication*

#### Response

| Status Code | Description |
| --- | --- |
| 204 No Content | Will be returned if the relation is deleted. |
| 401 Unauthorized | Will be returned if authentication fails. |
| 404 Not Found | Will be returned if the user does not exist. |



### Retrieve Followers

#### Request

#### Route

`GET /followers`

#### *Requires Basic Authentication*

#### Response

| Status Code | Description |
| --- | --- |
| 200 OK | MUST be returned upon successful processing of this request. The expected response body is below. |
| 401 Unauthorized | Will be returned if authentication fails. |
| 404 Not Found | MUST be returned when the current user does not have followers. |

#### Body

The response body MUST be a valid JSON Object (`{}`).

```json
[
  {
    "username": "test_user2",
    "birth_date": "2000-07-02",
    "bio": "some bio",
    "hometown": "Burgas"
  },
  {
    "username": "test_user3",
    "birth_date": "1997-09-04",
    "bio": "some bio2",
    "hometown": "Sofia"
  }
]
```

### Retrieve Following

#### Request

#### Route

`GET /following`

#### *Requires Basic Authentication*

#### Response

| Status Code | Description |
| --- | --- |
| 200 OK | The expected response body is below. |
| 401 Unauthorized | Will be returned if authentication fails. |
| 404 Not Found | Will be returned when the current user does not follow anyone. |


#### Body

The response body MUST be a valid JSON Object (`{}`).

```json
[
  {
    "username": "test_user3",
    "birth_date": "2000-07-02",
    "bio": "some bio",
    "hometown": "Burgas"
  },
  {
    "username": "test_user4",
    "birth_date": "1997-09-04",
    "bio": "some bio2",
    "hometown": "Sofia"
  }
]
```
## Posts


### Create Post

#### Request

#### Route

`POST /v1/posts`

#### *Requires Basic Authentication*

#### Body

The request body MUST be a valid JSON Object (`{}`).

```json
{
	"title": "title",
	"content": "content"
}
```

#### Response

| Status Code | Description |
| --- | --- |
| 201 Created | The resource is created |
| 401 Unauthorized | Will be returned if authentication fails. |
| 400 Bad Request | Will be returned if the request body is invalid. |



### Retrieve Posts

#### Request

#### *Route*

`GET /v1/posts`

#### *Requires Basic Authentication*

#### Response

| Status Code | Description |
| --- | --- |
| 200 OK | MUST be returned upon successful processing of this request. The expected response body is below. |
| 401 Unauthorized | Will be returned if authentication fails. |
| 404 Not Found | MUST be returned when no post are presented. |

#### Body

The response body MUST be a valid JSON Object (`{}`).

```json
[
  {
    "title": "title",
    "content": "content",
    "date": "2019-07-02T00:00:00Z"
  },
  {
    "title": "title2",
    "content": "content2",
    "date": "2019-07-02T00:00:00Z"
  }
]
```


# Database Model

![db_model](docs/images/db_schema.png)
