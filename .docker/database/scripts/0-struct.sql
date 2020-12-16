-- Role management tables
DROP TABLE IF EXISTS "user_account_roles";

CREATE TABLE "user_account_roles" (
  id SERIAL PRIMARY KEY,
  label VARCHAR(20) NOT NULL,
  isAdmin BOOLEAN
);

DROP TABLE IF EXISTS "user_accounts";

CREATE TABLE "user_accounts" (
  id SERIAL PRIMARY KEY,
  first_name VARCHAR(30) NOT NULL,
  last_name VARCHAR(50) NOT NULL,
  user_name VARCHAR(20) NOT NULL,
  password TEXT NOT NULL,
  user_role_id INTEGER NOT NULL,
  CONSTRAINT fk_user_role FOREIGN KEY(user_role_id) REFERENCES user_account_roles(id)
);

-- Questions types
DROP TABLE IF EXISTS "questions";

CREATE TABLE "questions" (
  id SERIAL PRIMARY KEY,
  label TEXT NOT NULL,
  question_type CHAR(2) NOT NULL DEFAULT 'TF' CHECK(question_type IN ('MC', 'TF', 'SL')),
  UNIQUE (id, question_type)
);

DROP TABLE IF EXISTS "mutiple_choice_questions";

CREATE TABLE "mutiple_choice_questions" (
  question_id INT NOT NULL,
  question_type CHAR(2) DEFAULT 'MC' CHECK (question_type = 'MC'),
  UNIQUE (question_id),
  FOREIGN KEY (question_id, question_type) REFERENCES questions(id, question_type)
);

DROP TABLE IF EXISTS "true_false_questions";

CREATE TABLE "true_false_questions" (
  question_id INT NOT NULL,
  question_type CHAR(2) DEFAULT 'TF' CHECK (question_type = 'TF'),
  answerIs BOOLEAN,
  FOREIGN KEY (question_id, question_type) REFERENCES questions(id, question_type)
);

DROP TABLE IF EXISTS "slider_questions";

CREATE TABLE "slider_questions" (
  question_id INT NOT NULL,
  question_type CHAR(2) DEFAULT 'SL' CHECK (question_type = 'SL'),
  answerMin INT NOT NULL,
  answerMax INT NOT NULL,
  FOREIGN KEY (question_id, question_type) REFERENCES questions(id, question_type)
);

-- Mutiple choice question answer
DROP TABLE IF EXISTS "multiple_choice_question_answers";

CREATE TABLE "multiple_choice_question_answers" (
  id SERIAL PRIMARY KEY,
  multiple_question_id INT NOT NULL,
  label TEXT NOT NULL,
  isCorrect BOOLEAN,
  CONSTRAINT fk_mutiple_choice_question_answer FOREIGN KEY(multiple_question_id) REFERENCES mutiple_choice_questions(question_id)
);

-- Quiz and its questions
DROP TABLE IF EXISTS "quizs";

CREATE TABLE "quizs" (
  id SERIAL PRIMARY KEY,
  label VARCHAR(50) NOT NULL,
  author INT NOT NULL,
  CONSTRAINT fk_quiz_author FOREIGN KEY (author) REFERENCES user_accounts(id)
);

DROP TABLE IF EXISTS "quiz_questions";

CREATE TABLE "quiz_questions" (
  quiz_id INT NOT NULL,
  question_id INT NOT NULL,
  CONSTRAINT fk_quiz_questions_quiz FOREIGN KEY (quiz_id) REFERENCES quizs(id),
  CONSTRAINT fk_quiz_questions_questions FOREIGN KEY (question_id) REFERENCES questions(id)
);