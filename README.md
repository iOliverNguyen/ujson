# µjson

[![Build Status](https://travis-ci.org/ng-vu/ujson.svg?branch=master)](https://travis-ci.org/ng-vu/ujson)
[![GoDoc](https://img.shields.io/badge/godoc-reference-blue.svg)](http://godoc.org/github.com/ng-vu/ujson)

A fast and minimal JSON parser and transformer that works on unstructured json.
Example use cases:

1. Walk through unstructured json:
   - [Print all keys and values](https://godoc.org/github.com/ng-vu/ujson#example-Walk)
   - Extract some values
2. Transform unstructured json:
   - [Remove all spaces](https://godoc.org/github.com/ng-vu/ujson#example-Walk--Reconstruct)
   - [Reformat](https://godoc.org/github.com/ng-vu/ujson#example-Walk--Reformat)
   - [Remove blacklist fields](https://godoc.org/github.com/ng-vu/ujson#example-Walk--RemoveBlacklistFields)
   - [Wrap int64 in string for processing by JavaScript](https://godoc.org/github.com/ng-vu/ujson#example-Walk--WrapInt64InString)

without fully unmarshalling it into a `map[string]interface{}`.

See usage and examples on [godoc.org](https://godoc.org/github.com/ng-vu/ujson).

**CAUTION: Behaviour is undefined on invalid json. Use on trusted input only.**

## Usage

The single most important function is [`Walk(input, callback)`](https://godoc.org/github.com/ng-vu/ujson#Walk),
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

| indent | key        | value   |
|:------:|:----------:|:-------:|
|`0`     |            |`{`      |
|`1`     |`"id"`      |`12345`  |
|`1`     |`"name"`    |`"foo"`  |
|`1`     |`"numbers"` |`[`      |
|`2`     |            |`"one"`  |
|`2`     |            |`"two"`  |
|`1`     |            |`]`      |
|`1`     |`"tags"`    |`{`      |
|`2`     |`"color"`   |`"red"`  |
|`2`     |`"priority"`|`"high"` |
|`1`     |            |`}`      |
|`1`     |`"active"`  |`true`   |
|`0`     |            |`}`      |

`indent` indicates the indentation of the key/value pair as if the json is
formatted properly. `key`s and `value`s are provided as raw literal. Strings are
always double-quoted. To get the original string, use
[`Unquote()`](https://godoc.org/github.com/ng-vu/ujson#Unquote).

`value` will never be empty (for valid json). You can test the first byte
(`value[0]`) to get its type:

- `n`: Null (`null`)
- `f`, `t`: Boolean (`false`, `true`)
- `0`-`9`: Number
- `"`: String, see [`Unquote()`](https://godoc.org/github.com/ng-vu/ujson#Unquote)
- `[`, `]`: Array
- `{`, `}`: Object

When processing arrays and objects, first the open bracket (`[`, `{`) will be
provided as `value`, followed by its children, and finally the close bracket
(`]`, `}`). When encounting open brackets, You can make the callback function
return `false` to skip the array/object entirely.

## LICENSE

[MIT License](https://opensource.org/licenses/mit-license.php)
