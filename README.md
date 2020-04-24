## GO API

## URLS

### GET request
- Used to get all users `http://localhost:3000/users`
- Used to get a specific user `http://localhost:3000/users/<id>`  example `http://localhost:3000/users/1`

### POST request
- Used to post(add) a user url `http://localhost:3000/users`
- POST sample {"FirstName":"<name>","SecondName":"<name>"}
- POST example {"FirstName":"John123","SecondName":"test33"}


### PUT request/Update user
- used to update a user with a specific id `http://localhost:3000/users/<id>`  example `http://localhost:3000/users/1`
- PUT sample {"ID":"<id>","FirstName":"<name>","SecondName":"<name>"}
- PUT example {"ID":"1","FirstName":"John","SecondName":"peter"}

### DELETE USER
- used to delete a user with a specific id `http://localhost:3000/users/<id>`  example `http://localhost:3000/users/1`

