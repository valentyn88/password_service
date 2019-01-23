## How to run password_server?

Run in the project root folder (Mac)

```
 make build-mac && make run
```

Run in the project root folder (Linux)

```
 make build-linux && make run
```

## How to run test ?

Run in the project root folder

```
 make test
```

## How to test password_server API ?

Lets assume that our password should not be less than 5 characters and not more that 15, otherwise we will get an error </br>
User can't generate more than 50 passwords once, otherwise we will get an error </br>
Our json request body has format {"letters_len":2,"spec_ch_len":2,"numbers_len":5, "count":10} </br>
'letters_len' - count letters in out password </br>
'spec_ch_len' - count special characters in out password </br>
'numbers_len' - count numbers in out password </br>
'count' - count passwords that we can generate once </br>

You can test API in your browser:

200 OK response

```
curl --header "Content-Type: application/json" \
  --request POST \
  --data '{"letters_len":2,"spec_ch_len":2,"numbers_len":5, "count":10}' \
  http://localhost:8080/generate
```

400 Bad Request invalid request params

```
curl --header "Content-Type: application/json" \
  --request POST \
  --data '{"letters_len":5,"spec_ch_len":5,"numbers_len":6, "count":10}' \
  http://localhost:8080/generate
```

400 Bad Request invalid count passwords

```
curl --header "Content-Type: application/json" \
  --request POST \
  --data '{"letters_len":5,"spec_ch_len":5,"numbers_len":5, "count":51}' \
  http://localhost:8080/generate
```
