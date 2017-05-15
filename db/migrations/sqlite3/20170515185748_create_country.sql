
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE `Countries` (
	`Id`	INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
	`Code`	char(2) NOT NULL UNIQUE,
	`Name`	varchar(100) NOT NULL,
	`CreatedAt`	timestamp NOT NULL,
	`UpdatedAt`	timestamp NOT NULL,
	`DeletedAt`	timestamp DEFAULT NULL
);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE Countries;
