
### Banking System
This banking project written in Golang that allows you to create and manage bank accounts, record balance changes, and perform money transfers between accounts. It provides a set of RESTful HTTP APIs built using the Gin framework and uses PostgreSQL to store account information and transaction history. Docker is used for local development and GitHub Actions for running unit tests automatically.
### Features:
- User account creation, balance tracking, and transaction history.
- Advanced features: Concurrent transaction handling with ACID compliance.

- Concurrent Transaction Handling: While not explicitly stated in the README, the use of RabbitMQ for transactions suggests some level of concurrency handling.
- ACID Compliance: Since PostgreSQL is used for storage, the database operations are inherently ACID-compliant when properly implemented.
- You'd need to review the implementation of the transaction logic (e.g., ensuring proper use of database transactions) to confirm ACID compliance.

- Concurrent Transaction Handling: Supported by RabbitMQ's queuing and QoS mechanisms.But not specific code mentioned
- ACID Compliance: Likely implemented in the database layer, not directly in this file.
- Additional code review (e.g., database transaction logic) is needed to confirm full compliance.
- Designed and implemented an ACID-compliant transaction system with support for real-time rollbacks and concurrency.
- Enabled secure multi-currency transactions with encryption (TLS/SSL) and fraud detection mechanisms.
- Optimized transaction processing time using Redis caching and database indexing, improving system throughput by X%.

- Money Transfer:

- Review methods or transactions in the code for ensuring atomic updates to both accounts.
- Look for RabbitMQ integration with money transfer APIs for asynchronous operations.

### Sqlc Generate Go Code:
Run the sqlc generate command in your project directory where the sqlc.yaml or sqlc.josn file is located:
- sqlc generate

C:\Users\Acer\Desktop\Bank-system\Bank-service\
├── internal/
│   ├── db/
│   │   ├── migration/    <- This folder should contain your migration files (like `schema.sql`)
│   │   ├── query/        <- This folder should contain your SQL query files (like `account.sql`, `user.sql`, etc.)
│   │   ├── sqlc/         <- This is where the generated Go code will be output
├── sqlc.yaml

- Manually create migration file: like this: 000001_init_schema.down.sql

### Documenting Important questions
 - Low level design
 - Detail design of project infrastructure
 - Corner cases & What library are used
 - What problem is solved
 - What issues you faces during development & how you solve it
 - What skills or techniques you learn in this project
 - Every thing from line by Line.


### Grpc-server multiline command:
protoc --proto_path=./proto `
    --go_out=./pb --go_opt=paths=source_relative `
    --go-grpc_out=./pb --go-grpc_opt=paths=source_relative `
    ./proto/create_user.proto `
    ./proto/bank_service.proto `
    ./proto/user.proto `
    ./proto/login_user.proto `
    ./proto/update_user.proto

### with Grpc-Gateway Setup:
protoc --proto_path=./proto `
    --go_out=./pb --go_opt=paths=source_relative `
    --go-grpc_out=./pb --go-grpc_opt=paths=source_relative `
    --grpc-gateway_out=./pb --grpc-gateway_opt=paths=source_relative `
    --openapiv2_out=./pb --openapiv2_opt=logtostderr=true `
    ./proto/create_user.proto `
    ./proto/bank_service.proto `
    ./proto/user.proto `
    ./proto/login_user.proto `
    ./proto/update_user.proto

- The swagger.json file generated by the protoc command is a Swagger/OpenAPI definition file that describes our gRPC API in a machine-readable format. It is useful if we're working with gRPC-Gateway and want to expose our gRPC services over HTTP, or if we want to generate API documentation.
- This --openapiv2_out=./pb flag generates OpenAPI v2 (Swagger) files.


### Single line command:
protoc --proto_path=./proto --go_out=./pb --go_opt=paths=source_relative --go-grpc_out=./pb --go-grpc_opt=paths=source_relative ./proto/create_user.proto ./proto/bank_service.proto ./proto/user.proto ./proto/login_user.proto ./proto/update_user.proto



### sqlc query:
- Create sqlc.yaml file in root directory
- run command: sqlc generate


### Generate the Statik Files: Swagger

Now, you can generate the Statik code for your swagger folder. Navigate to your project root where project_docs is located, and run the following command:
First install: go install github.com/rakyll/statik@latest

- statik -src=swagger

What this does:
- The statik command will take all the files inside the swagger directory and generate a Go file (statik.go) inside a statik directory.
- This statik.go file contains the embedded Swagger files, which you can then use in your Go project.

#### Pros of using swagger:
- Swagger provides a standard and clear way to document your API endpoints, request parameters, and response
- Swagger UI provides an interactive interface where users can try out your API endpoints directly from the browser.But, If you're comfortable using tools like Postman for API testing, Swagger UI may not add much value.
- It improves coordination between team members and makes it easy for new developers to understand the API contract.

### Use of db.dbml(for schema design) :
DBML is used in projects to abstract and simplify the database schema design, provide better documentation, and enable easier integration with tools that generate SQL schema, migrations, or even ORM models. It’s a helpful approach if you want to separate database design from implementation or if you're working in a larger team where schema changes need to be versioned and easily communicated.

### Use of Makefile:
But i prefer ci/cd instead of it.The purpose of a Makefile is to automate our development tasks, standardize how they are executed, and simplify complex or repetitive commands, making our development process more efficient and less error-prone.

### Facing problem during development
- sqlc setup due to 64 bit not suppoted for sqlc installation..for solve this i use msys installer
