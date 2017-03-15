-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE `Countries` (
	`Id`	INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
	`Code`	INTEGER NOT NULL UNIQUE,
	`Name`	INTEGER NOT NULL,
	`CreatedAt`	INTEGER NOT NULL,
	`UpdatedAt`	INTEGER NOT NULL,
	`DeletedAt`	INTEGER
);
-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE Countries;
