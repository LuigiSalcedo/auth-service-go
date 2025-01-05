
## Authors

- César Luigi Salcedo: [@LuigiSalcedo](https://www.github.com/LuigiSalcedo)




## Deployment

To build server executable:

```bash
go build server.go
```

To run server:
```bash
./server beauty
```

If You want to chage the server port change the server port conts in server.go file and re-build the project.




## Environment Variables

You need **PostgreSQL** installed and running in your machine.

There are 3 environment variable that You need to set in the **.env** file at project root.

`DB_CONNECTION_STRING`: any type of PostgreSQL connection string. Example: postgresql://user:password@localhost:5432/example_db

`JWT_SECRET`: the secret to sing the tokens

`CORS_ORIGINS`: origins that are allowed to request (optional, default = *)



