CREATE TABLE Users (
	id INT NOT NULL PRIMARY KEY AUTO_INCREMENT,
	username VARCHAR(30) NOT NULL UNIQUE,
	password VARCHAR(30) NOT NULL,
	balance FLOAT NOT NULL DEFAULT 0);

CREATE TABLE Transactions (
	id INT NOT NULL PRIMARY KEY AUTO_INCREMENT,
	debtor int NOT NULL,
	debtee int NOT NULL,
	amount FLOAT NOT NULL,
	description VARCHAR(100),
	ts TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	FOREIGN KEY (Debtor) REFERENCES Users(id),
	FOREIGN KEY (Debtee) REFERENCES Users(id));

DELIMITER $$
CREATE TRIGGER transaction_create AFTER INSERT ON Transactions 
	FOR EACH ROW BEGIN
		UPDATE Users SET balance=balance+NEW.amount WHERE id=NEW.debtee;
		UPDATE Users SET balance=balance-NEW.amount WHERE id=NEW.debtor;
	END$$
DELIMITER ;

DELIMITER $$
CREATE TRIGGER transaction_delete AFTER DELETE ON Transactions
	FOR EACH ROW BEGIN
		UPDATE Users SET balance=balance-OLD.amount WHERE id=OLD.debtee;
		UPDATE Users SET balance=balance+Old.amount WHERE id=OLD.debtor;
	END$$
DELIMITER ;

DELIMITER $$
CREATE TRIGGER transaction_update AFTER UPDATE ON Transactions
	FOR EACH ROW BEGIN
		UPDATE Users SET balance=balance-OLD.amount+NEW.amount WHERE id=OLD.debtee;
		UPDATE Users SET balance=balance+Old.amount-NEW.amount WHERE id=OLD.debtor;
	END$$
DELIMITER ;

