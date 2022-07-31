## API Endpoints

### Create Account

**Sample Request**

- `POST /createAccount`
  
- `Content-Type:` application/json
  
- `JSON body:` 


```json
{
"name":    "john",
"balance":     4,
}
```

|     **Key**    | **Type** | **Remarks** |
|:--------------:|:--------:|:-----------:|
| `name` | `string` | required  |
| `balance` | `int64` | required  |


**Sample Response**

```json
{
"account": {
"id": "95d694d0-941d-40bf-84d9-3c31b5554be9",
"name": "john",
"balance": 4,
"created_at": "2022-07-30T22:55:19.913509Z",
"last_modified": "2022-07-30T22:55:19.913509Z"
},
"message": "account successfully created"
}
```

### Withdraw from account

**Sample Request**

- `POST /withdraw"`

- `Content-type:` application/json

- `Sample request JSON body:`

```json
{
"id": "95d694d0-941d-40bf-84d9-3c31b5554be9",
"amount":   2000,
}
```

|     **Key**    | **Type** | **Remarks** |
|:--------------:|:--------:|:-----------:|
| `id` | `string` | required  |
| `amount` | `int64` | required  |


**Sample Response**

```json
{
"amount": 10,
"message": "amount successfully withdrawn",
"new_balance": 1990,
"old_balance": 2000
}
```



## Error HTTP Status Codes

 **Error Codes**                              | **Meaning**
----------------------------------------------|------------------------------------------------------------------------------
 **[400 or StatusBadRequest](#400)**          | Malformed parameters or other bad requests to the server.
 **[500 or StatusInternalServerError](#500)** | Processing errors, probably relating to backing services.
  **[404 or StatusNotFoundError](#404)** | Resource not found



**Sample 500 Response**

```json
{
"message": "could not withdraw amount"
},
```
Above error represents processing errors

**Sample 400 Response**
```json
{
"message":"could not withdraw amount due to insufficient balance"
}
```
Above error occurs when user does ot have enough balance in account for a withdrawal

```json
{
"message":"could not parse request. check API documentation"
}
```
Above error occurs when there is a missing request body field, we suggest you look at the docs carefully when this happens

**Sample 404 request**

```json
{
"message":"account origin specified was not found"
}
```
Above error occurs mostly when an account resource is not found

