# Simple CRUD Server with GoLang

## 🧱 Architecture
This project implements a RESTful API using the Gin web framework for GoLang. The API provides basic Create, Read, Update, and Delete (CRUD) functionality for a resource.

## 🪫 Database
The application uses Postgres as the database backend, interacting with it through the GORM (GoLang's Object-Relational Mapping) library.

## 📦 Main Packages
- **GORM**: Object-Relational Mapping (ORM) library for interacting with the Postgres database.
- **Gin**: Web framework for building the RESTful API.
- **jwt**: Library for generating and parsing JSON Web Tokens for authentication.
- **Bcrypt**: Library for hashing and salting passwords for secure storage.

This project provides a simple, yet robust, foundation for building a backend server in GoLang. It demonstrates how to set up a REST API, interact with a Postgres database using an ORM, and implement basic authentication and authorization functionality.