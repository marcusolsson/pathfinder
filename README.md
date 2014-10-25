pathfinder [![wercker status](https://app.wercker.com/status/e2cf7e5e72cb1eb6ea30cb2e07ac0ea0 "wercker status")](https://app.wercker.com/project/bykey/e2cf7e5e72cb1eb6ea30cb2e07ac0ea0)
==========

The routing context from the original DDD Sample Application, written in Go.

Pathfinder is accessible as a REST API, currently deployed to Heroku. Try it out here:

[ddd-pathfinder on Heroku](http://ddd-pathfinder.herokuapp.com/paths?from=SESTO&to=CNHKG)

#### GET /paths
Returns an array of candidate paths for a given route.

| URL Param | Description |
|:----------|:------------|
|from=[string]|UN locode of the origin|
|to=[string]|UN locode of the destination|