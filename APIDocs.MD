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


### Deposit to account

**Sample Request**

- `POST /deposit"`

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
"message": "amount successfully deposited",
"new_balance": 4000,
"old_balance": 2000
}
```


### Transfer to another account

**Sample Request**

- `POST /transfer"`

- `Content-type:` application/json

- `Sample request JSON body:`

```json
{
"origin": "95d694d0-941d-40bf-84d9-3c31b5554be9",
"destination": "95ff94d0-941d-40bf-8rd9-3c31b3214be9",
"amount": 2000
}
```

|     **Key**    | **Type** | **Remarks** |
|:--------------:|:--------:|:-----------:|
| `origin` | `string` | required  |
| `destination` | `string` | required  |
| `amount` | `int64` | required  |


**Sample Response**

```json
{
"account_destination_id":"95ff94d0-941d-40bf-8rd9-3c31b3214be9" ,
"account_origin_id": "95d694d0-941d-40bf-84d9-3c31b5554be9",
"amount": 2000,
}
```


### Find all transactions for a given ID

**Sample Request**

- `GET /transactions?id=<account_id>"`

- `Content-type:` application/json


|     **query string**    | **Type** | **Remarks** |
|:--------------:|:--------:|:-----------:|
| `id` | `string` | required  |

**Sample Response**

```json
{
"account_destination_id":"95ff94d0-941d-40bf-8rd9-3c31b3214be9" ,
"account_origin_id": "95d694d0-941d-40bf-84d9-3c31b5554be9",
"amount": 2000,
}
```
### Find all withdrawals(debits) and deposits(credit) for a user

**Sample Request**

- `GET /transactions/filter?credit=<account_id>"`

- `Content-type:` application/json


|     **query string**    | **Type** | **Remarks** |
|:--------------:|:--------:|:-----------:|
| `credit` | `string` | optional  | Depending on if you want to find deposits
| `debit` | `string` | optional  | Depending on if you want to find withdrawals

**Sample Response**

```json
{
"message": "transactions successfully retreived",
"transactions": {
"id": "85cce29b-cfa4-4185-a617-f0ccf2edf60b",
"credit": "ebb18d44-8d19-44bf-8a87-b8015411f3ab",
"debit": "",
"amount": 1000000,
"created_at": "2022-08-01T01:14:06.055102Z"
    }
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

