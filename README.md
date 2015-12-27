#Go api shell with examples

This api is based on a couple of packages.

That is:
  * goth - for authentication towards social media. Facebook is used as example
  * jwt-go - for token based api requests
  * gorilla mux - for routing inside the api
  * gorethink - RethinkDB library for go, saving list and userdata

This code is ready to be run in docker throug the docker file.
It also runs pure locally with the rundev.sh script, that is because its dependent on some env variables.
