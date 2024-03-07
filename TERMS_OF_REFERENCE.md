## Library WebAPI. Terms of reference

Need to create a REST API for 3 entities:

- Author
- Book
- Reader

The author is characterized by the following parameters:

- Id
- Full name
- Nickname
- Speciality

Book:

- Id
- Title
- Genre
- ISBN

Reader:

- Id
- Full name
- List of books

Paths:

- /authors
- /books
- /members

Special paths:

- /authors/{id}/books - list of author's available books
- /members/{id}/books - list of books borrowed by the reader

Actions:

- GET - produces a list of entities
- POST - creates an entity
- PATCH - updates the entity (needs id)
- DELETE - deletes the entity (needs id)

Procedure:

1. Create a REST API, all entities can be stored in memory
2. Integrate with Postgres database
3. Run the service in docker

Responses:

- In case of successful GET, PATCH, DELETE should return 200
- In case of missing resource should return error code 404
- In case of successful POST should return 201
- If a crooked request is received, it should return 400
- If other HTTTP methods will be received, it should return 405