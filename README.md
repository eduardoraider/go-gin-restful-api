# Go RESTful API with Gin and Swagger

This project is your guide to creating a robust Go-based back-end using the popular Gin framework. In this repository, you'll learn how to seamlessly integrate your Go API with a Dockerized database and harness the power of GORM, Go's renowned Object-Relational Mapping (ORM) tool. You'll explore the implementation of HTTP routes, data validation, and writing comprehensive Go tests. Additionally, you'll discover how to display API data within an HTML page and master the art of documenting APIs effectively using Gin Swagger within the Gin Framework.

## Application Overview

This Go application allows you to register students and their identification documents, such as CPF and ID. You will be able to perform all CRUD (Create, Read, Update, Delete) operations and search for students by their CPF number.

## Learning Topics

By studying this application, you can learn the following topics related to Go programming, Gin framework, and REST API development:

1. **Go and Gin: Creating a REST API with Simplicity**: Explore the simplicity of creating your own REST API with Go and the Gin framework.

2. **Installing and Creating HTTP Routes with Gin**: Learn how to install Gin and create HTTP routes for your API.

3. **Integrate Your Go API with a Database Running on Docker**: Discover how to integrate your Go API with a Postgres database deployed within a Docker container.

4. **Modularizing the Code, Models, and Database**: Organize your codebase effectively by modularizing it into models and connecting it to the database.

5. **Learn How to Use GORM, Go's Most Famous ORM**: Master the usage of GORM, a widely-used Object-Relational Mapping (ORM) tool for Go.

6. **Struct, Database, and ORM**: Understand the relationships between Go structs, the database, and the ORM.

7. **Learn How to Create Resource Searches Based on Struct Fields**: Implement search functionality based on fields in your struct.

8. **Deleting, Editing, and Searching for Students**: Perform CRUD operations on student records.

9. **Go: Validations, Tests, and HTML Pages**: Validate data within your Go API, write tests efficiently, and render HTML pages.

10. **Testing the Endpoints**: Dive deeper into testing your API endpoints to ensure expected behavior.

11. **Render HTML Pages with Gin**: Configure and customize HTML page rendering with Gin.

12. **Customize Your 404 Page**: Customize your Gin 404 error page for a better user experience.

## Running the Application

To run this Go REST API application, follow these steps:

1. Ensure you have Go installed on your system. If not, you can download and install it from the official [Go website](https://golang.org/dl/).

2. Make sure you have Docker installed and running on your system.

3. Open a terminal and navigate to the project directory.

4. Run the following Docker Compose command to set up the Postgres database and pgAdmin:

   ```bash
   docker-compose up
   ```

   This command will create a Postgres database running on port 5432, and install pgAdmin4 on port 54321. You can access pgAdmin at [http://localhost:54321](http://localhost:54321).

5. Open another terminal and run the following command to start the Go API:

   ```bash
   go run main.go
   ```

   This will execute the `main.go` script, which contains the Go API application. The application will start a server on port 8000.

6. To run the tests, use the following command:

   ```bash
   go test
   ```

7. Check the frontend in your web browser by navigating to:

   [http://localhost:8000/index](http://localhost:8000/index)

   You can also check greetings at the following address, where you can type your name after the slash:

   [http://localhost:8000/eduardo](http://localhost:8000/eduardo)

   To explore the 404 page, visit:

   [http://localhost:8000](http://localhost:8000)

8. You can perform all CRUD operations using the API Docs provided by Swagger:

   [http://localhost:8000/docs/index.html](http://localhost:8000/docs/index.html)

9. To stop the application, use the `Ctrl + C` keyboard shortcut in both terminals.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE.txt) file for details.

---

#### by Eduardo O Raider
🛠 🥋 **Software Engineer**