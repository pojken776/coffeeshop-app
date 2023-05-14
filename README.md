# coffeeshop-app

An app to help choose a coffeeshop to work at.

rest api documentation: <https://golang.cafe/blog/golang-rest-api-example.html>

## Database Design Considerations

- Auto increment the id column since we do not want to have to compute the value on every insert. [Example](https://www.postgresqltutorial.com/postgresql-tutorial/postgresql-serial/)
