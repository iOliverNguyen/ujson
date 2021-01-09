# µjson

[![Build Status](http://img.shields.io/travis/olvrng/ujson.svg?style=flat-square)](https://travis-ci.org/olvrng/ujson)
[![GoDoc](http://img.shields.io/badge/go-documentation-blue.svg?style=flat-square)](https://pkg.go.dev/github.com/olvrng/ujson)
[![License](http://img.shields.io/badge/license-mit-blue.svg?style=flat-square)](https://raw.githubusercontent.com/olvrng/ujson/master/LICENSE)

A fast and minimal JSON parser and transformer that works on unstructured json.

## Motivation

Sometimes we just want to make some minimal changes to a json document, or do
some generic transformations without fully unmarshalling it. For example,
removing [blacklist fields](https://godoc.org/github.com/olvrng/ujson#example-Walk--RemoveBlacklistFields2)
from response json. Why spend all the cost on unmarshalling into a `map[string]interface{}`
just to immediate marshal it again.

Read more on [dev.to/olvrng](https://dev.to/olvrng/json-a-minimal-json-parser-and-transformer-in-go-3dhb).

### Example use cases:

1. Walk through unstructured json:
   - [Print all keys and values](https://godoc.org/github.com/olvrng/ujson#example-Walk)
   - Extract some values
2. Transform unstructured json:
   - [Remove all spaces](https://godoc.org/github.com/olvrng/ujson#example-Walk--Reconstruct)
   - [Reformat](https://godoc.org/github.com/olvrng/ujson#example-Walk--Reformat)
   - [Remove blacklist fields](https://godoc.org/github.com/olvrng/ujson#example-Walk--RemoveBlacklistFields2)
   - [Wrap int64 in string for processing by JavaScript](https://godoc.org/github.com/olvrng/ujson#example-Walk--WrapInt64InString)

without fully unmarshalling it into a `map[string]interface{}`.

See usage and examples on [godoc.org](https://godoc.org/github.com/olvrng/ujson) and [dev.to/olvrng](https://dev.to/olvrng/json-a-minimal-json-parser-and-transformer-in-go-3dhb).

**Important**: *Behaviour is undefined on invalid json. Use on trusted input
only. For untrusted input, you may want to run it through
[`json.Valid()`](https://golang.org/pkg/encoding/json/#Valid) first.*

## Usage

The single most important function is [`Walk(input, callback)`](https://godoc.org/github.com/olvrng/ujson#Walk),
which parses the `input` json and call `callback` function for each key/value
pair processed.

Let's see an example:

```json
{
   "id": 12345,
   "name": "foo",
   "numbers": ["one", "two"],
   "tags": {"color": "red", "priority": "high"},
   "active": true
}
```

Calling `Walk()` with the above input will produce:

| level | key        | value   |
|:-----:|:----------:|:-------:|
|`0`    |            |`{`      |
|`1`    |`"id"`      |`12345`  |
|`1`    |`"name"`    |`"foo"`  |
|`1`    |`"numbers"` |`[`      |
|`2`    |            |`"one"`  |
|`2`    |            |`"two"`  |
|`1`    |            |`]`      |
|`1`    |`"tags"`    |`{`      |
|`2`    |`"color"`   |`"red"`  |
|`2`    |`"priority"`|`"high"` |
|`1`    |            |`}`      |
|`1`    |`"active"`  |`true`   |
|`0`    |            |`}`      |

`level` indicates the indentation of the key/value pair as if the json is
formatted properly. `key`s and `value`s are provided as raw literal. Strings are
always double-quoted. To get the original string, use
[`Unquote()`](https://godoc.org/github.com/olvrng/ujson#Unquote).

`value` will never be empty (for valid json). You can test the first byte
(`value[0]`) to get its type:

- `n`: Null (`null`)
- `f`, `t`: Boolean (`false`, `true`)
- `0`-`9`: Number
- `"`: String, see [`Unquote()`](https://godoc.org/github.com/olvrng/ujson#Unquote)
- `[`, `]`: Array
- `{`, `}`: Object

When processing arrays and objects, first the open bracket (`[`, `{`) will be
provided as `value`, followed by its children, and finally the close bracket
(`]`, `}`). When encountering open brackets, You can make the callback function
return `false` to skip the array/object entirely.
