# MiddleWhere
MiddleWhere is an open source endpoint Auth Header -> Query String Token proxy in Golang.

## Use cases
1. You have a secured endpoint you would like to publicly expose without exposing your secure auth token
2. You would like to have an application make a request to a secure endpoint but have no place to add an auth bearer token

## Default config
```
ENDPOINT_URL="https://google.com"
ENDPOINT_AUTH_TOKEN=
ENDPOINT_REQUST_METHOD=POST
SECURE_TOKEN=
```

## Our use case:
W use directus for the CMS for a number of our sites, but they also use a static site generator. So we have a webhook that sends a POST request to DroneCI to re-build the site with the new data. Directus webhooks do not let you set an auth header, so we needed a way to include a token in a querystring within the URL. This is our solution.