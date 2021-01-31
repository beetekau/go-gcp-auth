# go-cloudrun-auth
<!-- ALL-CONTRIBUTORS-BADGE:START - Do not remove or modify this section -->
[![All Contributors](https://img.shields.io/badge/all_contributors-1-orange.svg?style=flat-square)](#contributors-)
<!-- ALL-CONTRIBUTORS-BADGE:END -->

### code initially taken from https://github.com/CodeLinkIO/go-cloudfunction-auth

## Introduction

`go-gcp-auth` contains a client implementation for Google GCP following Google's Oauth2 spec at https://developers.google.com/identity/protocols/oauth2/service-account

Google provides their own official [oauth2](https://godoc.org/golang.org/x/oauth2) and higher level [Cloud API](https://github.com/googleapis/google-cloud-go), which include authentication part for most of all services. Unfortunately, these libraries does not support well authenticating with cloud run. There are also not so many useful documents or articles relate to this topic. Investigating on this take me much longer than my expectation. Therefore, I decide to publish a library here in case other people are also stuck at this step.

## Reference

The implementation inside this library follows Google's [official developers guideline](https://developers.google.com/identity/protocols/oauth2/service-account):

![Google OAuth2 Flow](https://developers.google.com/accounts/images/serviceaccount.png)

- Create and sign JWT using a custom wrapper of [google oauth2](https://github.com/golang/oauth2/tree/master/google)
- Send the signed JWT to Google's server to get an authenticated token
- Use the received token to create a custom http Client. This client attachs the token in all requests by default.


## License

MIT
