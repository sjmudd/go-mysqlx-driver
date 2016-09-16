/*
 * This script will create a database called xprotocol_test and allow a user
 * to access it. Minimal data will also be added to the database for
 * testing.
 *
 * Load in via the mysql command line using:
 * $ mysql -vvvt < xprotocol_test.sql
 */

CREATE DATABASE IF NOT EXISTS xprotocol_test;

CREATE USER IF NOT EXISTS xprotocol_user@'127.0.0.1';
ALTER USER xprotocol_user@'127.0.0.1' IDENTIFIED BY 'xprotocol_pass';
GRANT ALL PRIVILEGES ON xprotocol_test.* TO xprotocol_user@'127.0.0.1';

USE xprotocol_test

-- integer type
DROP TABLE IF EXISTS t_int;
CREATE TABLE t_int (id int NOT NULL PRIMARY KEY) ENGINE=InnoDB;
INSERT INTO t_int VALUES (1), (2), (3), (4), (5);

-- varchar type
DROP TABLE IF EXISTS t_varchar;
CREATE TABLE t_varchar (v varchar(255) NOT NULL PRIMARY KEY) ENGINE=InnoDB CHARSET=latin1;
INSERT INTO t_varchar VALUES ('A'), ('B'), ('C'), ('D'), ('E');

-- table with lots of types
DROP TABLE IF EXISTS field_test;
CREATE TABLE `field_test` (
  id int unsigned NOT NULL PRIMARY KEY AUTO_INCREMENT,
  `my_int_u` int unsigned NOT NULL,
  `my_int` int DEFAULT NULL,
  `my_tinyint_u` tinyint unsigned DEFAULT NULL,
  `my_tinyint` tinyint DEFAULT NULL,
  `my_smallint_u` smallint unsigned DEFAULT NULL,
  `my_smallint` smallint DEFAULT NULL,
  `my_mediumint_u` mediumint unsigned DEFAULT NULL,
  `my_mediumint` mediumint DEFAULT NULL,
  `my_bigint_u` bigint unsigned DEFAULT NULL,
  `my_bigint` bigint DEFAULT NULL,
  `my_double` double DEFAULT NULL,
  `my_float` float DEFAULT NULL,
  `my_decimal` decimal(10,2) DEFAULT NULL,
  `my_varchar` varchar(123) DEFAULT NULL,
  `my_char` char(5) DEFAULT NULL,
  `my_varbinary` varbinary(23) DEFAULT NULL,
  `my_binary` binary(3) DEFAULT NULL,
  `my_blob` tinyblob,
  `my_text` tinytext,
  `my_geometry` geometry DEFAULT NULL,
  `my_time` time DEFAULT NULL,
  `my_date` date DEFAULT NULL,
  `my_datetime` datetime DEFAULT NULL,
  `my_datetime_m` datetime(3) DEFAULT NULL,
  `my_year` year(4) DEFAULT NULL,
  `my_timestamp` timestamp NULL DEFAULT NULL,
  `my_timestamp_m` timestamp(6) NULL DEFAULT NULL,
  `my_set` set('1','2','3') DEFAULT NULL,
  `my_enum` enum('r','g','b') DEFAULT NULL,
  `my_bit` bit(17) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

INSERT INTO field_test (my_int_u,my_int,my_tinyint_u,my_tinyint,my_smallint_u,my_smallint)
VALUES
(  0,   0,   0,-128,     0,-32768),
(999,-999, 255, 127, 65535, 32767);

-- -------------------------------------------------------------------------
-- tinyint unsigned

DROP TABLE IF EXISTS xprotocol_test_tinyint_u;
CREATE TABLE `xprotocol_test_tinyint_u` (
  id int unsigned NOT NULL PRIMARY KEY AUTO_INCREMENT,
  `my_tinyint_u` tinyint unsigned
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

INSERT INTO xprotocol_test_tinyint_u
    (my_tinyint_u)
VALUES
--  (NULL), -- not handled yet
    (0),    -- min
    (255);  -- max

-- -------------------------------------------------------------------------
-- tinyint

DROP TABLE IF EXISTS xprotocol_test_tinyint;
CREATE TABLE `xprotocol_test_tinyint` (
  id int unsigned NOT NULL PRIMARY KEY AUTO_INCREMENT,
  `my_tinyint` tinyint
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

INSERT INTO xprotocol_test_tinyint
    (my_tinyint)
VALUES
--  (NULL), -- not handled yet
    (-128), -- min
    (127);  -- max

-- -------------------------------------------------------------------------
-- unsigned int

-- test each type in a single table
DROP TABLE IF EXISTS xprotocol_test_int_u;
CREATE TABLE `xprotocol_test_int_u` (
  id int unsigned NOT NULL PRIMARY KEY AUTO_INCREMENT,
  `my_int_u` int unsigned
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

INSERT INTO xprotocol_test_int_u
    (my_int_u)
VALUES
--  (NULL), -- not handled yet
    (0),    -- min
    (999);  -- some value

-- -------------------------------------------------------------------------
-- int

DROP TABLE IF EXISTS xprotocol_test_int;
CREATE TABLE `xprotocol_test_int` (
  id int unsigned NOT NULL PRIMARY KEY AUTO_INCREMENT,
  `my_int` int
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

INSERT INTO xprotocol_test_int
    (my_int)
VALUES
--  (NULL), -- not handled yet
    (-999), -- negative value
    (999);  -- positive value

-- -------------------------------------------------------------------------
-- float

DROP TABLE IF EXISTS xprotocol_test_float;
CREATE TABLE `xprotocol_test_float` (
  id int unsigned NOT NULL PRIMARY KEY AUTO_INCREMENT,
  `my_float` float
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

INSERT INTO xprotocol_test_float
    (my_float)
VALUES
--  (NULL),  -- not handled yet
    (0),     -- zero
    (-99.9), -- negative value
    (99.9);  -- positive value

-- -------------------------------------------------------------------------
-- double

DROP TABLE IF EXISTS xprotocol_test_double;
CREATE TABLE `xprotocol_test_double` (
  id int unsigned NOT NULL PRIMARY KEY AUTO_INCREMENT,
  `my_double` double
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

INSERT INTO xprotocol_test_double
    (my_double)
VALUES
--  (NULL),  -- not handled yet
    (0),     -- zero
    (-99.9), -- negative value
    (99.9);  -- positive value
