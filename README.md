# Simple Bank

A secure and scalable banking system API built with Go, featuring account management, secure money transfers, and user authentication.

## 🌟 Features

- **Account Management**
  - Create and manage bank accounts
  - Support for multiple currencies (USD, EUR)
  - Real-time balance tracking
  - List accounts with pagination

- **Money Transfers**
  - Secure money transfers between accounts
  - Transaction consistency with database transactions
  - Currency validation
  - Deadlock prevention with ordered locking

- **Authentication & Security**
  - User authentication system
  - Token-based authorization
  - Secure API endpoints
  - Password protection

## 🛠️ Technology Stack

- **Backend**: Go (Golang)
- **Framework**: Gin Web Framework
- **Database**: PostgreSQL
- **SQL Toolkit**: SQLC
- **Authentication**: JWT (JSON Web Tokens)
- **Testing**: Go's built-in testing package with testify assertions
- **Configuration**: YAML-based configuration

## 🔧 Project Structure

```
simple_bank/
├── api/              # API handlers and server setup
├── db/               # Database related code
│   ├── sqlc/        # SQL queries and generated Go code
│   └── util/        # Utility functions
├── token/           # Authentication token management
└── main.go          # Application entry point
```

## ⚙️ Setup and Installation

1. Clone the repository
```bash
git clone https://github.com/sathwikshetty33/simple_bank.git
cd simple_bank
```

2. Install dependencies
```bash
go mod tidy
```

3. Set up the database
```bash
# Create the database (PostgreSQL must be installed)
createdb simple_bank

# Run database migrations (if applicable)
# Add migration steps here
```

4. Configure the application
- Create a configuration file based on the provided template
- Set up environment variables as needed

5. Run the application
```bash
go run main.go
```

## 🚀 API Endpoints

### Account Operations
- `POST /accounts` - Create a new account
- `GET /accounts/:id` - Get account details
- `GET /accounts` - List accounts with pagination

### Transfer Operations
- `POST /transfers` - Create a new money transfer

### User Operations
- User-related endpoints (authentication, registration, etc.)

## 💡 Usage Examples

### Creating an Account
```json
POST /accounts
{
    "currency": "USD"
}
```

### Making a Transfer
```json
POST /transfers
{
    "from_account_id": 1,
    "to_account_id": 2,
    "amount": 100,
    "currency": "USD"
}
```

## 🧪 Testing

The project includes comprehensive test coverage for core functionality:

```bash
# Run all tests
go test -v ./...

# Run specific test
go test -v ./db/sqlc
```

## 🔐 Security Features

- SQL injection prevention using prepared statements
- Secure password hashing
- Token-based authentication
- Parametrized queries
- Transaction isolation levels for data consistency

## 📝 License

Add your license information here

## 👥 Contributing

1. Fork the repository
2. Create a feature branch
3. Commit your changes
4. Push to the branch
5. Create a new Pull Request

## ✨ Acknowledgments

- Thanks to all contributors
- Special thanks to the Go community
- Inspired by modern banking system architectures
