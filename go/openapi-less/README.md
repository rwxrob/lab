# OpenAPI without the OpenAPI Generators

This repo is full of experimentation developing code "by hand" that matches an OpenAPI specification in a way that can be validated, but not necessarily generated.

## Motivation

Generating code with OpenAPI can be really a pain to maintain long term. The generators are woefully limited and ancient --- especially for Go stuff. They are also all written in a *single* Java compiled artifact. To be honest, the entire OpenAPI implementation is absolutely butt-fucking-ugly because of this. This is probably why the project gets absolutely no love from the Go community.

However, we do have to use OpenAPI these days to cover the documentation and submission to Apigee (which has become something of an enterprise gate keeping thing for all microservices).

* [Test root link](/)
* [Test root relative link](/)
