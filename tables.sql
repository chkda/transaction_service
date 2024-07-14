CREATE DATABASE loco;

CREATE TABLE `transactions` (
  `Id` bigint NOT NULL,
  `Types` varchar(20) DEFAULT NULL,
  `Amount` double DEFAULT NULL,
  `ParentId` bigint DEFAULT NULL,
  PRIMARY KEY (`Id`),
  KEY `ParentId` (`ParentId`),
  CONSTRAINT `transactions_ibfk_1` FOREIGN KEY (`ParentId`) REFERENCES `transactions` (`Id`)
)