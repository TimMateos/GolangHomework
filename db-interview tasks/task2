Нужно описать модель библиотеки. Есть 3 сущности: "Автор", "Книга", "Читатель".
Физически книга только одна и может быть только у одного читателя.
Нужно составить таблицы для библиотеки так, чтобы это учесть.
Первый запрос — выбрать названия всех книг, которые на руках.
Второй запрос — выбрать названия всех книг в библиотеке, у которых больше 3 авторов.
Третий запрос — выбрать имена топ-3 читаемых авторов на данный момент.   в ответе таблицы и запросы выведи в sql

CREATE TABLE author (
  author_id INT PRIMARY KEY AUTO_INCREMENT,
  name VARCHAR(100) NOT NULL
);

CREATE TABLE reader (
  reader_id INT PRIMARY KEY AUTO_INCREMENT,
  name VARCHAR(100) NOT NULL
);

CREATE TABLE book (
  book_id INT PRIMARY KEY AUTO_INCREMENT,
  title VARCHAR(50),
  reader_id INT,
  in_stock BOOLEAN DEFAULT TRUE,
  FOREIGN KEY (reader_id) REFERENCES reader (reader_id)
);

CREATE TABLE book_author (
    book_id INT NOT NULL,
    author_id INT NOT NULL,
    PRIMARY KEY (book_id, author_id),
    FOREIGN KEY (book_id) REFERENCES book(book_id),
    FOREIGN KEY (author_id) REFERENCES author(author_id)
);

Запросы:
№1 - SELECT title FROM book
     WHERE reader_id IS NOT NULL;

№2 - SELECT title FROM book b
     LEFT JOIN book_author ba ON b.author_id = ba.book_id
     WHERE reader_id IS NOT NULL 
     GROUP BY b.book_id, b.title
     HAVING COUNT(ba.author_id) > 3;

№3 - SELECT a.name
     FROM author a
     JOIN book_author ba ON a.author_id = ba.author_id
     JOIN book b ON ba.book_id = b.book_id
     WHERE b.reader_id IS NOT NULL
     GROUP BY a.author_id, a.name
     ORDER BY COUNT(b.book_id) DESC
     LIMIT 3;
