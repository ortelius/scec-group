# Ortelius v11 Group Microservice

> Version 11.0.0

RestAPI for the Group Object
![Release](https://img.shields.io/github/v/release/ortelius/scec-group?sort=semver)
![license](https://img.shields.io/github/license/ortelius/scec-group)

![Build](https://img.shields.io/github/actions/workflow/status/ortelius/scec-group/build-push-chart.yml)
[![MegaLinter](https://github.com/ortelius/scec-group/workflows/MegaLinter/badge.svg?branch=main)](https://github.com/ortelius/scec-group/actions?query=workflow%3AMegaLinter+branch%3Amain)
![CodeQL](https://github.com/ortelius/scec-group/workflows/CodeQL/badge.svg)
[![OpenSSF-Scorecard](https://api.securityscorecards.dev/projects/github.com/ortelius/scec-group/badge)](https://api.securityscorecards.dev/projects/github.com/ortelius/scec-group)

![Discord](https://img.shields.io/discord/722468819091849316)

## Path Table

| Method | Path | Description |
| --- | --- | --- |
| GET | [/msapi/group](#getmsapigroup) | Get a List of Groups |
| POST | [/msapi/group](#postmsapigroup) | Create a Group |
| GET | [/msapi/group/:key](#getmsapigroupkey) | Get a Group |

## Reference Table

| Name | Path | Description |
| --- | --- | --- |

## Path Details

***

### [GET]/msapi/group

- Summary  
Get a List of Groups

- Description  
Get a list of groups for the user.

#### Responses

- 200 OK

***

### [POST]/msapi/group

- Summary  
Create a Group

- Description  
Create a new Group and persist it

#### Responses

- 200 OK

***

### [GET]/msapi/group/:key

- Summary  
Get a Group

- Description  
Get a group based on the _key or name.

#### Responses

- 200 OK

## References
