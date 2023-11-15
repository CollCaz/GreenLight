# GreenLight

## What is this?
GreenLight is an api for retrieving movie data
Taken from the wonderful [Let's Go Further](https://lets-go-further.alexedwards.net) book by Alex Edwards.

## How to use?
This API consists of the following end points
- GET /v1/healthcheck : Retrieve system information
- POST /v1/movies : Add a movie to the database
- GET /v1/movies/:id : Retrieve a movie by id
- PATCH /v1/movies/:id : Update movie information by id
- DELETE /v1/movie/:id : Delete a movie by id
And more to come.

## How does this code differ from the book?
- [x] Using [log](https://github.com/charmbracelet/log) for logging
- [ ] Unit Tests
- [ ] Integration Tests
- [ ] CI/CD using github actions

## Note on liscene
The author did not put any licence on his code, this does not mean it is opensource; so keep that in mind.

