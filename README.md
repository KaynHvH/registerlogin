# User Authentication Service

A simple user authentication service written in Go, using SQLite for data storage. The service supports user registration, login, account deletion, and password change functionalities.

## Features

- **User Registration**: Allows new users to register with a username and password.
- **User Login**: Authenticates users with their username and password.
- **Delete Account**: Allows users to delete their account by providing their username and password.
- **Change Password**: Allows users to change their password after providing their current password.

## Technologies

- Golang
- SQLite
- bcrypt
- Gorilla Mux

## Setup

1. **Clone the Repository**

   ```bash
   git clone https://github.com/kaynhvh/registerlogin.git
   cd registerlogin
   ```

2. **Install Dependencies**

   Ensure you have Go installed on your system. Install the necessary Go packages:

   ```bash
   go mod tidy
   ```

3. **Initialize the Database**

   The application initializes the SQLite database automatically upon startup. Ensure the directory where the database file will be created has appropriate write permissions.

4. **Run the Application**

   Start the server with the following command:

   ```bash
   make run
   ```

   The server will be accessible at `http://localhost:8080`.

## API Endpoints

### Register User

- **URL**: `/register`
- **Method**: `POST`
- **Request Body**:

  ```json
  {
    "username": "string",
    "password": "string"
  }
  ```

- **Success Response**:
    - **Code**: `201 Created`
    - **Content**:

      ```json
      {
        "message": "User registered successfully"
      }
      ```

- **Error Response**:
    - **Code**: `400 Bad Request` or `409 Conflict` (for username already taken)

### Login User

- **URL**: `/login`
- **Method**: `POST`
- **Request Body**:

  ```json
  {
    "username": "string",
    "password": "string"
  }
  ```

- **Success Response**:
    - **Code**: `200 OK`
    - **Content**:

      ```json
      {
        "message": "Login successful"
      }
      ```

- **Error Response**:
    - **Code**: `401 Unauthorized` (for invalid credentials)

### Delete Account

- **URL**: `/delete`
- **Method**: `DELETE`
- **Request Body**:

  ```json
  {
    "username": "string",
    "password": "string"
  }
  ```

- **Success Response**:
    - **Code**: `204 No Content`
    - **Content**:

      ```json
      {
        "message": "Successfully deleted",
        "username": "string"
      }
      ```

- **Error Response**:
    - **Code**: `400 Bad Request` or `401 Unauthorized` (for invalid password) or `404 Not Found` (for user not found)

### Change Password

- **URL**: `/change-password`
- **Method**: `PUT`
- **Request Body**:

  ```json
  {
    "username": "string",
    "old_password": "string",
    "new_password": "string"
  }
  ```

- **Success Response**:
    - **Code**: `200 OK`
    - **Content**:

      ```json
      {
        "message": "Password changed successfully"
      }
      ```

- **Error Response**:
    - **Code**: `400 Bad Request` or `401 Unauthorized` (for invalid old password)
