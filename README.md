pathfinder 
==========

[![Build Status](https://travis-ci.org/marcusolsson/pathfinder.svg?branch=master)](https://travis-ci.org/marcusolsson/pathfinder)
[![License MIT](https://img.shields.io/badge/license-MIT-lightgrey.svg?style=flat)](LICENSE)

The routing context from the original DDD Sample Application, written in Go.

Pathfinder is accessible as a REST API, currently deployed to Heroku. Try it out here:

[ddd-pathfinder on Heroku](http://ddd-pathfinder.herokuapp.com/paths?from=SESTO&to=CNHKG)

#### GET /paths
Returns an array of candidate paths for a given route.

| URL Param | Description |
|:----------|:------------|
|from=[string]|UN locode of the origin|
|to=[string]|UN locode of the destination|
