USE uccdemy;



INSERT INTO users (id,email,user_name, first_name, last_name, user_type, password_hash, creation_time, last_updated) VALUES
(1, 'pepito@gmail.com', 'pepe16','pepe', 'sanchez', 1, 'password','2024-05-30 00:00:00', '2024-05-30 00:00:00'),
(2, 'luisito@gmail.com', 'luisito24','luis', 'fernandez', 0, 'password2','2024-05-30 00:00:00', '2024-05-30 00:00:00');

INSERT INTO courses (id, name, description, instructor_id, category, requirements, length, image_url,creation_time, last_updated) VALUES
(1, 'Curso de programacion', 'Aprende a programar en Java', 1, 'Programacion', 'Conocimientos basicos de programacion', 10, 'https://www.google.com', '2024-05-30 00:00:00', '2024-05-30 00:00:00'),
(2, 'Curso de cocina', 'Aprende a cocinar', 2, 'Cocina', 'Conocimientos basicos de cocina', 5, 'https://www.google.com', '2024-05-30 00:00:00', '2024-05-30 00:00:00');