USE uccdemy;

INSERT INTO users (id,email,user_name, first_name, last_name, user_type, password_hash, creation_time, last_updated) VALUES
(1, 'pepito@gmail.com', 'pepe16','pepe', 'sanchez', 1, 'password','2024-05-30 00:00:00', '2024-05-30 00:00:00'),
(2, 'luisito@gmail.com', 'luisito24','luis', 'fernandez', 0, 'password2','2024-05-30 00:00:00', '2024-05-30 00:00:00');

INSERT INTO courses (id, name, description, instructor_id, category, requirements, length, image_url,creation_time, last_updated) VALUES
(1, 'Curso de programacion', 'Aprende a programar en Java', 1, 'Programacion', 'Conocimientos basicos de programacion', 10, 'https://i.blogs.es/905760/1366_2000-1-/1366_2000.jpeg', '2024-05-30 00:00:00', '2024-05-30 00:00:00'),
(2, 'Curso de cocina', 'Aprende a cocinar', 2, 'Cocina', 'Conocimientos basicos de cocina', 5, 'https://d3puay5pkxu9s4.cloudfront.net/curso/4271/800_imagen.jpg', '2024-05-30 00:00:00', '2024-05-30 00:00:00');