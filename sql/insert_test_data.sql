INSERT INTO USERS (USERNAME, PASSWORD)
VALUES ("James", "James");

INSERT INTO USERS (USERNAME, PASSWORD)
VALUES ("Matt", "James");

INSERT INTO USERS (USERNAME, PASSWORD)
VALUES ("Alex", "James");

-- Alex is owed some amount of money
INSERT INTO TRANSACTIONS (LOANER_ID, NUM_OWERS)
VALUES ( 2, 1 )

--Matt Owes Alex 100 dollars
INSERT INTO OWERS ( AMOUNT, TX_ID, USER_ID )
VALUES ( 100, 0, 1 )




