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
[GET] `/pokemons` : Pokemon list </br>
[GET] `/pokemon/scores` : Pokemon score board </br>
[GET] `/battles` : Get Battle List </br>
[POST] `/battle/auto` </br>
[POST] `/battle/:battleID/eliminate-pokemon` </br>

*All API Docs attached in /docs/Farmacare - Pokemons.postman_collection.json

## ERD
![ERD](https://raw.githubusercontent.com/Tasrifin/pokemonfight-go/develop/docs/POKEMON%20-%20ERD.png)
