# jsonny-walker

Basic HTTP server to add json objects, and get the set of json paths along with the top k leaf elements.

## Running the server
To run and test this server, first clone this repository and run `go install`, to pull in all dependencies and create an executable `jsonny-walker`, in `$HOME/go/bin/`

```
$ go install ./...
```
Once the installation is completed, start the server by running the executable.
```
$ ~/go/bin/jsonny-walker 

   ____    __
  / __/___/ /  ___
 / _// __/ _ \/ _ \
/___/\__/_//_/\___/ v3.3.10-dev
High performance, minimalist Go web framework
https://echo.labstack.com
____________________________________O/_______
                                    O\
⇨ http server started on [::]:1323
```

The HTTP server runs on port `1323` by default and serves POST calls on the following endpoints:
 - `/addJSONObject`
 - `/getPathsData`

## Making calls to the server
You can use any API client to make POST calls to the endpoints mentioned above. The examples here use cURL commands to demonstrate the usecases.

### Adding JSON objects
`/addJSONObject` adds JSON object under `data` to an in-memory representation of a collection of JSON objects added so far using this endpoint.

Example:
```
curl -X POST \
 http://127.0.0.1:1323/addJSONObject \
 -H 'content-type: application/json' \
 -d '{"data": <json/object>}}'
```
### Get set of JSON paths along with top K leaf elements
`/getPathsData` can be called to get a set of JSON paths with occurance fraction above a `Threshold` and along with the top K leaf elements for leaf paths.

Example (to get JSON paths with occurance fraction > 0.0 along top 3 leaf elements):
```
curl -X POST \
  http://127.0.0.1:1323/getPathsData \
  -H 'content-type: application/json' \
  -d '{"K": 3, "Threshold": 0.0}'
```
