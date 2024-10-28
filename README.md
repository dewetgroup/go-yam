# go-yam
![build and test](https://github.com/ndewet/go-yam/actions/workflows/go.yml/badge.svg)

Yet Another Multiplexer - lightweight, ergonomic routing module for Go.

## Why do we need Yet Another Multiplexer?

While this module draws inspiration from:
- [gorilla/mux](https://github.com/gorilla/mux)
- [go-chi/chi](https://github.com/go-chi/chi)

Continuing the philisophies of compositional routing, Yet Another Multiplexer seeks to solve issues with developer ergonomics. 

This inital version seeks to solve:
- Automatic conversion of errors to http responses.
- Easier serialization of objects for response.

## What does the future look like?

In the future Yet Another Multiplexer will also implement functionality as seen in similar packages for other languages, such as:
- [fastapi/fastapi](https://github.com/fastapi/fastapi)

The following might be implemented:
- Typing of request bodies (to the best extend that the language will allow)
- Model Validation
