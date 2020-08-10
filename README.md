# JSON to CSV converter

Simple JSON to CSV converter. Only able to convert single level json as per the example. Can't handle nesting. The output will be a pipe separated CSV.

**Run Example**

```
go run main.go ./example.json
```

**Arguments**

Accepts a single argument that is the file path for the json/csv file you wish to convert. CSV's are expected to be pipe separated values.

If a CSV is supplied it will output json and vice versa.