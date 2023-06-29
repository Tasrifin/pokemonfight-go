# pokemonfight-go : Tasrifin

- Golang
- MySQL

## Application Manual
1. Set-up database (tables will be migrated automatically)
2. Run :
   ```
   go run main.go
   ```

## Endpoints List
[GET] `/pokemons` : Pokemon list
[GET] `/pokemon/scores` : Pokemon score board
[GET] `/battles` : Get Battle List
[POST] `/battle/auto`
[POST] `/battle/:battleID/eliminate-pokemon`

*All API Docs attached in /docs/Farmacare - Pokemons.postman_collection.json

## ERD
