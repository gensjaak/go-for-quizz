-- Roles
INSERT INTO
  user_account_roles(label, isAdmin)
VALUES
  ('Utilisateur', FALSE),
  ('Administrateur', TRUE);

-- Users
INSERT INTO
  user_accounts(
    first_name,
    last_name,
    user_name,
    password,
    user_role_id
  )
VALUES
  ('Doe', 'John', 'johndoe', 'xxx', 1),
  ('Camus', 'Albert', 'camuslegrand', 'xxx', 2),
  ('Gates', 'Billy', 'bilou', 'xxx', 1);

-- All questions
INSERT INTO
  questions (label, question_type)
VALUES
  ('Les satellites volent dans l’espace.', 'TF'),
  ('La Terre a un climat global.', 'TF'),
  (
    'L’atmosphère est composée de trois couches protectrices.',
    'TF'
  ),
  (
    'Les avions et les ballons météorologiques peuvent voler dans la stratosphère.',
    'TF'
  ),
  (
    'La troposphère, soit la couche la plus près de la surface de la Terre, représente 75 % de la masse entière de l’atmosphère.',
    'TF'
  ),
  (
    'Le réchauffement terrestre se produit en partie parce que certains gaz ne peuvent pas s’échapper de l’atmosphère dans l’espace.',
    'TF'
  ),
  (
    'La Bastille a été prise en quelle année ?',
    'SL'
  ),
  (
    'L''institu de formation HETIC a été créée en quelle année ?',
    'SL'
  ),
  (
    'Quelle est l''année de la fin de la 2ème guerre mondiale ?',
    'SL'
  ),
  (
    'La Déclaration des droits de l''homme et du citoyen est proclamée le',
    'MC'
  ),
  (
    'Quel événement entraîne la chute de la monarchie constitutionnelle le 10 août 1792 ?',
    'MC'
  ),
  (
    'À quelle date fut guillotiné Louis XVI ?',
    'MC'
  ),
  (
    'En 1801, Bonaparte pacifie la situation religieuse grâce au',
    'MC'
  ),
  (
    'Napoléon abdique déﬁnitivement après la défaite',
    'MC'
  );

-- True or False questions
INSERT INTO
  true_false_questions (question_id, answerIs)
VALUES
  (1, FALSE),
  (2, TRUE),
  (3, FALSE),
  (4, TRUE),
  (5, TRUE),
  (6, TRUE);

-- Slider questions
INSERT INTO
  slider_questions (question_id, answerMin, answerMax)
VALUES
  (7, 1780, 1810),
  (8, 1996, 2010),
  (9, 1935, 1955);

-- Multiple choice questions
INSERT INTO
  mutiple_choice_questions(question_id)
VALUES
  (10),
  (11),
  (12),
  (13),
  (14);

INSERT INTO
  multiple_choice_question_answers(multiple_question_id, label, isCorrect)
VALUES
  (10, '20 juin 1789', FALSE),
  (10, '4 août 1789', FALSE),
  (10, '26 août 1789', TRUE),
  (
    11,
    'L''insurrection populaire au palais des Tuileries',
    TRUE
  ),
  (
    11,
    'Le retour du roi après la fuite à Varennes',
    FALSE
  ),
  (
    11,
    'La déclaration de guerre de la France à l''Autriche',
    FALSE
  ),
  (12, '22 septembre 1792', FALSE),
  (12, '21 janvier 17911', TRUE),
  (12, '2 juin 1793', FALSE),
  (13, 'Consulat', FALSE),
  (13, 'Concordat', TRUE),
  (13, 'Culte de l''Être suprême', FALSE),
  (14, 'd''Austerlitz', FALSE),
  (14, 'de Leipzig', FALSE),
  (14, 'de Waterloo', TRUE);

-- Quiz and its questions
INSERT INTO
  quizs (label, author)
VALUES
  (
    'Quizz: Revolution française et Sciences Tech',
    2
  );

INSERT INTO
  quiz_questions (quiz_id, question_id)
VALUES
  (1, 1),
  (1, 2),
  (1, 3),
  (1, 4),
  (1, 5),
  (1, 6),
  (1, 7),
  (1, 8),
  (1, 9),
  (1, 10),
  (1, 11),
  (1, 12),
  (1, 13),
  (1, 14);