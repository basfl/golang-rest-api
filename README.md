# Routes
* GET "/events"
* GET"/events/:id"
* POST "/signup"
* POST "/login"
### Adding middleware (Auhentication required )
* POST "/events"
* PUT "/events/:id"
* DELETE "/events/:id"
* POST "/events/:id/register"
* DELETE "/events/:id/register"

# Structure
  * db package -> Initiate database and created related table
  * models-> contains data structure represented inside database with CRUD methods
  * routes-> contains routes
  * uils->contains hashing and JWT 
  * api-test->for testing

* Geting config values from .env file 
  