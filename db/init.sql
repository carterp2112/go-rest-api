CREATE TABLE `Post` (
  `ID` int PRIMARY KEY NOT NULL AUTO_INCREMENT,
  `Title` varchar(255) NOT NULL,
  `Content` longtext NOT NULL
);

CREATE TABLE `User` (
  `ID` int UNIQUE PRIMARY KEY NOT NULL AUTO_INCREMENT,
  `First_name` varchar(255) NOT NULL,
  `Last_name` varchar(255) NOT NULL,
  `Password` varchar(255) NOT NULL,
  `Username` varchar(255) NOT NULL,
  `Is_admin` boolean NOT NULL
);
