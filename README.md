###What is it about?
This repo is a golang service for uvatel clients which previously resides inside the juvo repo, since the courier concept will remove the uvatel engine entirely, it is a good practice to seperate the demo service into another repository.
###How to start it?
This service is built on beego framework, by running the following command, the server can be started
`bee run`
###Where can I check APIs?
This repo supports autogenerating API docs, by running `bee run -downdoc=true -gendoc=true .` and go to `:8080/swagger/`, you can see the docs
###Demo of API endpoints
####Create a User
Post `http://localhost:8080/user`
```cassandraql
{
    "msisdn": 1234,
    "active":true,
    "country_code":"US"
}
```
You are expecting to get
```cassandraql
{
  "data": {
    "id": 1,
    "msisdn": 1234,
    "active": true,
    "country_code": "US",
    "created_at": "2019-09-15T19:18:52.664177+08:00",
    "updated_at": "2019-09-15T19:18:52.664179+08:00",
    "imsi": 0,
    "expired_at": "2019-09-15T19:18:52.664182+08:00"
  },
  "code": 200,
  "error": {
    "message": ""
  }
}
```
####Insert a Package
Post `http://localhost:8080/package`
```cassandraql
{
    "name":"trial4",
    "total":43.2,
    "unit":"SGD",
    "type":1,
    "amount":3.1,
    "expiration":4
}
```
You are expecting to get
```cassandraql
{
  "data": {
    "id": 1,
    "Name": "trial4",
    "total": 43.2,
    "unit": "SGD",
    "type": 1,
    "amount": 3.1,
    "expiration": 4
  },
  "code": 200,
  "error": {
    "message": ""
  }
}
```
####Add a Package to a User
Post `http://localhost:8080/user/1234/package/trial4`
You are expecting to get
```cassandraql
{
  "data": {
    "id": 1,
    "total": 43.2,
    "used": false,
    "remaining": 0,
    "unit": "",
    "expired_at": "2019-09-15T19:21:40.786754+08:00",
    "user": {
      "id": 1,
      "msisdn": 1234,
      "active": true,
      "country_code": "US",
      "created_at": "2019-09-15T19:18:52.664177+08:00",
      "updated_at": "2019-09-15T19:18:52.664179+08:00",
      "imsi": 0,
      "expired_at": "2019-09-15T19:18:52.664182+08:00"
    },
    "package": {
      "id": 1,
      "Name": "trial4",
      "total": 43.2,
      "unit": "SGD",
      "type": 1,
      "amount": 3.1,
      "expiration": 4
    }
  },
  "code": 200,
  "error": {
    "message": ""
  }
}
```
####Create an account
Post `http://localhost:8080/account`
```cassandraql
{
    "msisdn": 23242,
    "amount":54
}
```
You are expecting to get
```cassandraql
{
  "data": {
    "id": 1,
    "msisdn": 23242,
    "active": true,
    "core_balance": 54,
    "currency": ""
  },
  "code": 200,
  "error": {
    "message": ""
  }
}
```
####Deduct balance from an account
Post `http://localhost:8080/balance/deduct`
```cassandraql
{
    "msisdn": 23242,
    "amount":54
}
```
You are expecting to get
```cassandraql
{
  "data": {
    "id": 1,
    "msisdn": 23242,
    "active": true,
    "core_balance": 0,
    "currency": ""
  },
  "code": 200,
  "error": {
    "message": ""
  }
}
```
####Add balance from an account
Post `http://localhost:8080/balance/add`
```cassandraql
{
    "msisdn": 23242,
    "amount":54
}
```
You are expecting to get
```cassandraql
{
  "data": {
    "id": 1,
    "msisdn": 23242,
    "active": true,
    "core_balance": 54,
    "currency": ""
  },
  "code": 200,
  "error": {
    "message": ""
  }
}
```
####Set balance from an account
Post `http://localhost:8080/balance/add`
```cassandraql
{
    "msisdn": 23242,
    "amount":56
}
```
You are expecting to get
```cassandraql
{
  "data": {
    "id": 1,
    "msisdn": 23242,
    "active": true,
    "core_balance": 56,
    "currency": ""
  },
  "code": 200,
  "error": {
    "message": ""
  }
}
```