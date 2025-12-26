-- Write your PostgreSQL query statement below
SELECT id
FROM Weather as w1
INNER JOIN Weather w2
ON w1.recordDate = w2.recordDate + INTERVAL '1 day'