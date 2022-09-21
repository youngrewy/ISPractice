CREATE DATABASE IF NOT EXISTS webapp;
USE webapp;

CREATE TABLE IF NOT EXISTS
users(
    id          	 INTEGER UNIQUE NOT NULL,
    login       	 TEXT NOT NULL,
    money_amount     INTEGER NOT NULL,
    card_number      TEXT NOT NULL,
    status      	 INTEGER NOT NULL
);

CREATE TABLE IF NOT EXISTS
pwds(
    id 		INTEGER UNIQUE NOT NULL,
    pwd		TEXT NOT NULL
);

REPLACE INTO users VALUES
        (1,     'admin',    	 999999999,     '4716173295951919',      1),
        (2,     'youngrewy',     0,             '4024007163179418',      1),
        (3,     'pro100Oleg',    100,           '5581109980137223',      0),
        (4,     'sasha',         35353,         '4485095517670242',      1),
        (5,     'qwerty',        99,            '5493634387046229',      0)
;
REPLACE INTO pwds VALUES
    (1,     'admin'),
    (2,     'wasd335'),
    (3,     'OlegSveta'),
    (4,     'sasha05071975'),
    (5,     'qwerty1337')
;
