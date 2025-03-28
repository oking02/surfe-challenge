# surfe-challenge


## Run Instructions

There is a makefile to provide a number of handy commands


### Configuration

The service uses a few environment variables to configure how it runs:

There are defaults that will work out the gate.

- `HTTP_PORT` - port to use with the http server
- `STORAGE_DRIVER` - type of storage will be used.
- `USER_DATA_LOCATION` - location of users.json file to be loaded on bootstrap
- `ACTION_DATA_LOCATION` - location of actions.json file to be loaded on bootstrap

### Run

Go 1.24 is required

`make run` - starts the application

### Development

To install tooling run: e.g. linter

`make setup` - installs tools into separate go mod file

## Client ID
This appears in all the interfaces and controller. Whilst currently it will always be "". 
The idea is for this to model the service supporting multiple client with a single datasource.

Assumption would be that its added by an API gateway type of process
Possibly using JWTs or other Token based auth systems.

## Routes

### Get User By ID
``[GET] api/v1/users/{id}``

### Get User Actions
``[GET] api/v1/users/actions``

Return all actions from this user.

A count is found in the metadata object

### Get Next Action Probability
``[GET] /api/v1/actions/{type}/probability/next``

`type` refers to the action type. e.g. `VIEW_CONTACTS`

### Get Referral Index
``[GET] /api/v1/actions/REFER_USER/referral-index``

