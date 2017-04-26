CREATE TABLE `Countries` (
	`Id`	INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
	`Code`	char(2) NOT NULL UNIQUE,
	`Name`	varchar(100) NOT NULL,
	`CreatedAt`	timestamp NOT NULL,
	`UpdatedAt`	timestamp NOT NULL,
	`DeletedAt`	timestamp DEFAULT NULL
);
-- +