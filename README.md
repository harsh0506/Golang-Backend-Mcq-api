# Quiz API

The Quiz API is a RESTful API built with Golang and the Gin framework. It provides endpoints to manage quiz data, including creating quizzes, retrieving quizzes, updating quizzes, and deleting quizzes. The API uses a PostgreSQL database for data storage.

## Technologies Used

- Golang
- Gin
- PostgreSQL
- SQL
- JSON
- RESTful API

## Features

- Create Quiz
- Retrieve Quiz
- Update Quiz
- Delete Quiz

## Getting Started

To get started with the Quiz API, follow these steps:

1. Install Golang and PostgreSQL.
2. Clone the repository: `git clone <repository-url>`.
3. Install dependencies: `go mod download`.
4. Set up the environment variables in the `.env` file.
5. Run the API: `go run main.go`.

The Quiz API exposes the following endpoints:

- `POST /MCQ`: Create a new multiple-choice question.
- `GET /MCQs/:num/:num2`: Retrieve multiple-choice questions based on a specific range.
- `GET /MCQ/:id`: Retrieve a single multiple-choice question by its ID.
- `GET /MCQ-sub/:sub`: Retrieve multiple-choice questions based on a specific subject.
- `GET /MCQ-subjectsArray`: Retrieve all available subjects of multiple-choice questions.
- `PUT /MCQ/:id`: Update a single multiple-choice question by its ID.
- `DELETE /MCQ/:id`: Delete a single multiple-choice question by its ID.

## Things I Learned

- Building RESTful APIs with Golang and Gin.
- Integrating PostgreSQL with Golang.
- CRUD functionality for managing quiz data.
- Handling JSON requests and responses.
- Working with environment variables.

## Future Scope

- Implement rate limiting and authorization.
- Deploy on AWS using Elastic Beanstalk.
- Add pagination and filtering options.
- Write comprehensive unit tests.
- Containerize with Docker.

Feel free to explore and use the Quiz API. Contributions are welcome!

---

**Note**: Make sure to configure the PostgreSQL database connection details and environment variables properly for a successful connection.

Enjoy using the Quiz API and have fun managing quizzes!
