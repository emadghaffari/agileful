# get postgres connected database queries

## filters
you can filter by: 
`1- query: SELECT,INSERT,UPDATE,DELETE`<br />
`2- order: desc, asc`<br />
`3- limit: limit the responses`<br />
`4- offset: add offset for your response`<br />

## pagination
with limit and offset you can set your pagination model

## sorting by time spent
you can order the response in ascending or descending for time spend every query

### example:
`endpoint: localhost:3000/search`
`method: POST`
`body:`
```json
{
    "query":"SELECT",
    "limit":100,
    "offset":0,
    "order":"asc"
}
 ```
