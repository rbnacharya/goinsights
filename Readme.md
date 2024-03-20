# GoInsights

GoInsights is a demo project designed to showcase the integration of Go with PostgreSQL databases.

## Prerequisites

Before starting, ensure you have the following prerequisites installed:

- Go 1.22.1

## Installing and Running GoInsights

To get GoInsights up and running, follow these steps:

1. **Clone the repository**

    ```sh
    git clone https://github.com/rbnacharya/goinsights.git
    ```

2. **Navigate into the project directory**

    ```sh
    cd goinsights
    ```

3. **Build the project**

   Using Go:
    ```sh
    go build .
    ```

4. **Run the application**

    ```sh
    go run .
    ```

## Application Configuration

To configure your application, you'll need to set up environment variables as per `.env.example`. Copy this file and adjust it to your needs:

```sh
cp .env.example .env
```

Then, update `.env` with the appropriate database configuration.

### Example `.env` file:

```
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=yourpassword
DB_NAME=goinsightsdb
```

## API Usage

### Sending a Request

To send a request to the GoInsights API, use the following `curl` command:

```sh
curl --location 'http://localhost:8080/request' \
--header 'Content-Type: application/json' \
--data '{
    "customerID": 1,
    "tagID": 1,
    "userID": "abc",
    "timestamp": 1710055228,
    "remoteIP": "192.168.1.2",
    "userAgent": "Firefox/1201"
}'
```

#### Request Constraints:

- All fields are required.
- `customerID` should already exist in the database. (By default, 1L is available. Check src/main/resources/data.sql to prepopulate data.)
- `timestamp` should not be after the current time and should be within 1 year from the current time.
- If `remoteIP` or `userAgent` is blacklisted, the request will be marked as invalid.

### Retrieving Statistics

To retrieve statistics for a single customer for a single day, use the following `curl` command:

```sh
curl --location 'http://localhost:8080/statistics?customerID=1&date=2024-03-10'
```

#### Sample Response Format:

```json
{
    "message": "Success",
    "data": {
        "request_count": 3,
        "invalid_count": 1,
        "customer_id": 3
    }
}
```

This response indicates the number of requests and the number of invalid requests for the specified customer on the given date.

## Contact

For further information or queries, you can reach out at rbn@gerer.io