drop database if EXISTS billing;
create database if not exists billing character set UTF8;

CREATE USER IF NOT EXISTS 'billing'@'%' IDENTIFIED BY 'billing';

grant all on billing.* to 'billing'@'%';

CREATE TABLE billing.`ROLE`
(
    `ID`   TINYINT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    `NAME` enum('SUPER_ADMIN','ADMINISTRATOR_FULL','ADMINISTRATOR_BASIC','TRADER','SUPER_AGENT','AGENT', 'PROSUMER') DEFAULT NULL UNIQUE
);

CREATE TABLE billing.`PROVIDER` -- partner/spółka obrotu
(
    `ID`                      INT(11) NOT NULL AUTO_INCREMENT PRIMARY KEY,
    `NAME`                    VARCHAR(255) NOT NULL UNIQUE,
    `STATUS`                  TINYINT      DEFAULT 0,
    `KRS`                     VARCHAR(255) DEFAULT NULL,
    `NIP`                     VARCHAR(255) DEFAULT NULL,
    `REGON`                   VARCHAR(255) DEFAULT NULL,
    `EMAIL`                   VARCHAR(255) DEFAULT NULL,
    `PHONE_NUMBER`            VARCHAR(255) DEFAULT NULL,
    `STREET`                  VARCHAR(255) DEFAULT NULL,
    `BUILDING_NUMBER`         VARCHAR(255) DEFAULT NULL,
    `APARTMENT_NUMBER`        VARCHAR(255) DEFAULT NULL,
    `POSTAL_CODE`             VARCHAR(255) DEFAULT NULL,
    `PROVINCE`                VARCHAR(255) DEFAULT NULL,
    `CITY`                    VARCHAR(255) DEFAULT NULL,
    `COUNTRY`                 VARCHAR(255) DEFAULT NULL,
    `LICENSE_ID`              VARCHAR(255) DEFAULT NULL,
    `LICENSE_EXPIRATION_DATE` DATE,
    `LICENSE_AREA`            VARCHAR(255) DEFAULT NULL
);

CREATE TABLE billing.`WORKER` -- pracownik spółki/partnera AGENT/HANDLOWIEC, hanldowiec moze byc opiekunem agenta
(
    `ID`                     INT(11) NOT NULL AUTO_INCREMENT PRIMARY KEY,
    `PROVIDER_ID`            INT(11) NOT NULL,
    `SUPERVISOR`             INT(11) DEFAULT NULL,
    `FIRST_NAME`             VARCHAR(255) DEFAULT NULL,
    `LAST_NAME`              VARCHAR(255) DEFAULT NULL,
    `EMAIL`                  VARCHAR(255) DEFAULT NULL,
    `PHONE`                  VARCHAR(255) DEFAULT NULL,
    `STREET`                 VARCHAR(255) DEFAULT NULL,
    `BUILDING_NUMBER`        VARCHAR(255) DEFAULT NULL,
    `APARTMENT_NUMBER`       VARCHAR(255) DEFAULT NULL,
    `POSTAL_CODE`            VARCHAR(255) DEFAULT NULL,
    `PROVINCE`               VARCHAR(255) DEFAULT NULL,
    `CITY`                   VARCHAR(255) DEFAULT NULL,
    `COUNTRY`                VARCHAR(255) DEFAULT NULL,
    `WORK_START_DATE`        DATE         DEFAULT NULL, -- data rozpoczecia wspolpracy
    `WORK_END_DATE`          DATE         DEFAULT NULL, -- data zakonczenia wspolpracy
    `BLOCKCHAIN_ACC_ADDRESS` VARCHAR(255) DEFAULT NULL,
    `STATUS`                 TINYINT(4) DEFAULT 0,
    FOREIGN KEY (`PROVIDER_ID`) REFERENCES billing.`PROVIDER` (`ID`),
    FOREIGN KEY (`SUPERVISOR`) REFERENCES billing.`WORKER` (`ID`) ON DELETE CASCADE
);

CREATE TABLE billing.`CUSTOMER_TYPE` -- typ klienta
(
    `NAME` enum('CONSUMER','PROSUMER','PRODUCER') PRIMARY KEY
);

CREATE TABLE billing.`CUSTOMER_ACCOUNT` -- prosument/klient
(
    `ID`                  INT(11) NOT NULL AUTO_INCREMENT PRIMARY KEY,
    `CUSTOMER_TYPE_NAME`  enum('CONSUMER','PROSUMER','PRODUCER') NOT NULL,
    `PROVIDER_ID`         INT(11) NOT NULL,
    `FIRST_NAME`          VARCHAR(255) DEFAULT NULL,
    `LAST_NAME`           VARCHAR(255) DEFAULT NULL,
    `PESEL`               VARCHAR(255) DEFAULT NULL UNIQUE,
    `NIP`                 VARCHAR(255) DEFAULT NULL UNIQUE,
    `REGON`               VARCHAR(255) DEFAULT NULL UNIQUE,
    `EMAIL`               VARCHAR(255) DEFAULT NULL,
    `PHONE`               VARCHAR(255) DEFAULT NULL,
    `STREET`              VARCHAR(255) DEFAULT NULL,
    `BUILDING_NUMBER`     VARCHAR(255) DEFAULT NULL,
    `APARTMENT_NUMBER`    VARCHAR(255) DEFAULT NULL,
    `POSTAL_CODE`         VARCHAR(255) DEFAULT NULL,
    `PROVINCE`            VARCHAR(255) DEFAULT NULL,
    `CITY`                VARCHAR(255) DEFAULT NULL,
    `COUNTRY`             VARCHAR(255) DEFAULT NULL,
    `STATUS`              TINYINT(4) DEFAULT '0',
    `BANK_ACC_NUMBER`     VARCHAR(255) DEFAULT NULL,
    `WORKER_ID`           INT(11) DEFAULT NULL,
    `REGISTRATION_NUMBER` BIGINT(12) DEFAULT NULL,
    `BUSINESS_TYPE`       enum('B2B','B2C') DEFAULT NULL,
    `KRS`                 VARCHAR(255) DEFAULT NULL,
    `LINE_OF_BUSINESS` VARCHAR(255) DEFAULT NULL,
    FOREIGN KEY (`WORKER_ID`) REFERENCES billing.`WORKER` (`ID`),
    FOREIGN KEY (`CUSTOMER_TYPE_NAME`) REFERENCES billing.`CUSTOMER_TYPE` (`NAME`),
    FOREIGN KEY (`PROVIDER_ID`) REFERENCES billing.`PROVIDER` (`ID`),
    FOREIGN KEY (`WORKER_ID`) REFERENCES billing.`WORKER` (`ID`)
);

CREATE TABLE billing.`USER` -- konto dla partnera/spółki obrotu
(
    `ID`                   INT(11) NOT NULL AUTO_INCREMENT PRIMARY KEY,
    `LOGIN`                VARCHAR(255) NOT NULL UNIQUE,                              -- czy potrzebne? Logowanie z podaniem emaila?
    `PASSWORD`             TEXT         NOT NULL,                                     -- zahashowane haslo (bcrypt)
    `ACTIVATION_CODE`      VARCHAR(45)  DEFAULT NULL UNIQUE,
    `PROVIDER_ID`          INT(11) DEFAULT NULL,
    `WORKER_ID`            INT(11) DEFAULT NULL,
    `CUSTOMER_ACCOUNT_ID`  INT(11) DEFAULT NULL,                                      -- czy customer_account tez bedzie mial konto w systemie?
    `ROLE_ID`              TINYINT(4) DEFAULT NULL,
    `ACTIVE`               TINYINT(4) DEFAULT '1',                                    -- dla customera false
    `MUST_CHANGE_PASSWORD` TINYINT(4) DEFAULT '0',
    `EMAIL`                VARCHAR(255) DEFAULT NULL,
    FOREIGN KEY (`PROVIDER_ID`) REFERENCES billing.`PROVIDER` (`ID`),
    FOREIGN KEY (`WORKER_ID`) REFERENCES billing.`WORKER` (`ID`),
    FOREIGN KEY (`CUSTOMER_ACCOUNT_ID`) REFERENCES billing.`CUSTOMER_ACCOUNT` (`ID`), -- ?
    FOREIGN KEY (`ROLE_ID`) REFERENCES billing.`ROLE` (`ID`)
);

CREATE TABLE billing.`SERVICE_ACCESS_POINT`
(
    `ID`           INT(11) NOT NULL AUTO_INCREMENT PRIMARY KEY,
    `ACCOUNT_ID`   INT(11) NOT NULL,         -- czy potrzebne? Logowanie z podaniem emaila?
    `CITY`         VARCHAR(45) DEFAULT NULL, -- zahashowane haslo (bcrypt)
    `ADDRESS`      VARCHAR(45) DEFAULT NULL,
    `SAP_CODE`     VARCHAR(45) DEFAULT NULL,
    `METER_NUMBER` VARCHAR(45) DEFAULT NULL,
    `PROVIDER_ID`  INT(11) NOT NULL,         -- czy customer_account tez bedzie mial konto w systemie?
    `NAME`         VARCHAR(45) DEFAULT NULL,
    FOREIGN KEY (`PROVIDER_ID`) REFERENCES billing.`PROVIDER` (`ID`),
    FOREIGN KEY (`ACCOUNT_ID`) REFERENCES billing.`CUSTOMER_ACCOUNT` (`ID`)
);

