CREATE TABLE IF NOT EXISTS `User` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `email` VARCHAR(255) NOT NULL,
  `password` VARCHAR(500) NOT NULL,
  `confirmationToken` VARCHAR(500) NOT NULL,
  `enabled` TINYINT(1) NOT NULL DEFAULT 1,
  `passwordRequestAt` DATETIME NULL,
  `passwordToken` VARCHAR(45) NULL,
  PRIMARY KEY (`id`),
  UNIQUE INDEX `id_UNIQUE` (`id` ASC),
  UNIQUE INDEX `email_UNIQUE` (`email` ASC),
  UNIQUE INDEX `confirmationToken_UNIQUE` (`confirmationToken` ASC))
ENGINE = InnoDB;
