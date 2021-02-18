stringify
=========

```go
stringify.StringifyMap(map[string]interface{}{
  "a": 123,
  "b": big.NewInt(3232332),
  "c": map[string]interface{}{
    "b": 456,
  },
  "d": []byte{1, 2},
})
```

will return:

```
{
  "a": "123",
  "b": "3232332",
  "c": {
    "b": "456"
  },
  "d": "\x1\x2"
}
```

## LICENSE

MIT
