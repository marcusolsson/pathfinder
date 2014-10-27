pathfinder [![wercker status](https://app.wercker.com/status/b8863587e0a6ae46fe3c508732392e09/s "wercker status")](https://app.wercker.com/project/bykey/b8863587e0a6ae46fe3c508732392e09)
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