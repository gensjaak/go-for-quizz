# 🏆 go-for-quizz

This API is a backend of a quizz application. It provide such routes for any front-end app to create quiz, questions, multiple choices questions, true or false questions and slider questions. It also supports user management and quiz responding pause and reprise. It's build with GoLang, so it's incredibly fast !

## 🌈 Features

- Create an user (who can be a simple user or an administrator)

Note that only administrators can create quizs

- Create mutiple choice questions

When creating mutiple choice questions, you should provide the proposals and precise which ones are corrects

- Create true or false questions

- Create slider questions

Answers to these questions can be accepted if it's between an interval

## ✨ API demo

The API demo will be available soon on Heroku

## 🔥 Setup

The project requires Golang v`1.14.4`. Make sure you have [GoLang](https://golang.org/dl) installed and follow these steps.

### 🚀 Production mode

```sh
cd go-for-quizz

docker-compose up --build
```

### 🎮 Technical Choices

Feel free to discuss with any contributor about the technical choices that were made.

Go version : `+1.14.4`

PostgreSQL : `+13.1`

### 🏄‍♂️ Authors

[Jean-Jacques](https://github.com/gensjaak)
