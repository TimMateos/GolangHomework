-- Написать select который сделает OUT

CREATE TABLE test(
                     id SERIAL,
                     name VARCHAR
)
если сделать обычный select, то получим:
                         | id | name  |
                         | 1  | name1 |
                         | 2  | name2 |
                         | 3  | name3 |
                         | 4  | name4 |
                         | 5  | name5 |
                         OUT - то как надо вывести
                         | id | name  |
                         | 1  | name1 |
                         | 3  | name3 |
                         | 2  | name2 |
                         | 4  | name4 |
                         | 5  | name5 |


WITH ranked AS (
    SELECT 
        id, 
        name,
        id % 2 AS is_odd,
        ROW_NUMBER() OVER (PARTITION BY id % 2 ORDER BY id) AS rn
    FROM test
)
SELECT id, name
FROM ranked
ORDER BY 
    (rn - 1) / 2, 
    is_odd DESC,  
    id;

Сначала присваиваем всем четным и нечетным числам номера по группам 
Нечётным числам 1, 3, 5 номера rn = 1, 2, 3., а чётным числам 2, 4 номера rn = 1, 2.
Далее разбиваем их на пачки при сортировке выражением  (rn - 1) / 2
ORDER BY отсортирует сначала группу 0 (нечётные 1, 3, потом чётные 2, 4), а затем группу 1 (нечётное 5).
