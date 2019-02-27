# µjson

[![Build Status](https://travis-ci.org/ng-vu/ujson.svg?branch=master)](https://travis-ci.org/ng-vu/ujson)
[![GoDoc](https://img.shields.io/badge/godoc-reference-blue.svg)](http://godoc.org/github.com/ng-vu/ujson)

A minimal implementation of a json parser and transformer. Example use cases:

1. Walk through unstructured json:
   - [Print all keys](https://godoc.org/github.com/ng-vu/ujson#example-Walk)
   - Extract some values
2. Transform unstructured json:
   - [Remove all spaces](https://godoc.org/github.com/ng-vu/ujson#example-Walk--Reconstruct)
   - [Reformat](https://godoc.org/github.com/ng-vu/ujson#example-Walk--Reformat)
   - [Remove blacklist fields](https://godoc.org/github.com/ng-vu/ujson#example-Walk--RemoveBlacklistFields)
   - [Wrap int64 in string for processing by JavaScript](https://godoc.org/github.com/ng-vu/ujson#example-Walk--WrapInt64InString)

without fully unmarshalling it into a `map[string]interface{}`.

See usage and examples on [godoc.org](https://godoc.org/github.com/ng-vu/ujson).

**CAUTION: Behaviour is undefined on invalid json. Use on trusted input only.**

## LICENSE

[MIT License](https://opensource.org/licenses/mit-license.php)
