# rest-api-tutorial

# user-service

# REST API

GET /users -- list of users -- 200, 404, 500
GET /users/:id -- user by id -- 200, 404, 500
POST /users/:id -- create user -- 201, 4xx, Header Location: url
PUT /users/:id -- fully update user -- 204/200, 404, 400, 500
PATCH /users/:id -- partically update user -- 204/200, 404, 400, 500
DELETE /users/:id -- delete user by id -- 204, 404, 400

# About config.yml

If you use VS code, you need to add this file to cmd/main
If you use Goland. you need to add thif file to your workspace directory
I don't know which problem, but it's work
