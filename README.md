# Recipe Service

-----

The recipe service contains and holds all endpoints related to managing and creating recipes. Below is a list of all recorded endpoints and their  HTTP verbs

## Service Features

1. **CATEGORIES**
    * Crud operation for recipe categories
1. **RECIPES**
    * Crud operation for recipes
1. **INGREDIENTS**
    * Crud operation for ingredients

### Technologies

-----

 1. Golang
 1. MysQL
 1. Gorillar Mux
 1. Glide
 1. Docker

### Below are the API endpoints and their functions

-----

verbs  | endpoints | description
--     | --        | --
GET    | /recipes                    | Get all recipes
POST   | /recipes                    | Creates new recipe
PATCH  | /recipes/<id>              | Updates a recipe
DELETE | /recipes/<id>              | Deletes a recipe
GET    | /recipes/<id>              | Gets a single recipe
GET    | /categories                | Gets all categories
POST   | /categories                | Creates a new category
PATCH  | /categories/<id>           | Updates a category
DELETE | /categories/<id>           | Deletes a category
GET    | /categories/<id>           | Gets a single category
GET    | /categories/<id>/recipes   | Gets all recipe that belongs to a category
GET    | /categories?q={query}      | Searches for categories
GET    | /categories?limit={limit}  | Limits categories
GET    | /categories?page={offset}  | paginates categories  
GET    | /ingredients               | Gets all ingredients related to recipe
GET    | /ingredients               | Gets all ingredients
POST   | /ingredients               | Creates new ingredients
PATCH  | /ingredients/<id>          | Updates an ingredient
DELETE | /ingredients/<id>          | Deletes an ingredient
GET    | /ingredients/<id>          | Gets a single ingredient
GET    | /ingredients?q={query}     | Searches for ingredients
GET    |  /ingredients?limit={limit} | limits search result
GET    | /ingredients?offset={offset} | paginates result set


### POSTMAN COLLECTION

-----
You can get all endpoints directly on you postman collection by clicking on the link below

[![Run in Postman](https://run.pstmn.io/button.svg)](https://app.getpostman.com/run-collection/9ba34049f2e2545fa7bf)


### How to use

* Clone repository 
* Cd into the recipe folder
* run docker-compose build
* run docker-compose up
* make your API calls to localhost:8080/api/v1/recipes

### How to Contribute
* Fork repo to your branch
* Clone the repo
* Add a feature and raise a PR to  the base branch
* PR is reviewed and merged 
