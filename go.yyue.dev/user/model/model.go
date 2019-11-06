package model

// CREATE TABLE `userinfo` (
// 	`uid` INT(10) NOT NULL AUTO_INCREMENT,
// 	`name` VARCHAR(64) NULL DEFAULT NULL,
// 	`phone` VARCHAR(64) NULL DEFAULT NULL,
// 	`password` VARCHAR(64) NULL DEFAULT NULL,
// 	`created` TIMESTAMP CURRENT_TIMESTAMP,
// 	`updated` TIMESTAMP NO UPDATE CURRENT_TIMESTAMP,
// 	PRIMARY KEY (`uid`)
// );

type User struct {
	UID      int
	Name     string
	Phone    int
	Password string
}
