CREATE DATABASE IF NOT EXISTS smartplus;
USE smartplus;

CREATE TABLE IF NOT EXISTS customer (
    id INT(11) AUTO_INCREMENT COMMENT 'id',
    lastname VARCHAR(32) NOT NULL COMMENT 'also surname or family_name',
    firstname VARCHAR(32) NOT NULL COMMENT 'also given name',
    middlename VARCHAR(64) COMMENT 'optional ',
    email VARCHAR(255) COMMENT 'maximum 254 characters',
    phone CHAR(16) COMMENT 'maximum 16 characters',
    facebook VARCHAR(128) COMMENT 'facebook link',
    school VARCHAR(256) COMMENT 'language: English',
    company VARCHAR(256) COMMENT 'language: English',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT 'created at',
    PRIMARY KEY (`id`),
    UNIQUE INDEX `by_email` (`email`) USING BTREE
)
ENGINE = InnoDB DEFAULT CHARSET = utf8;


CREATE TABLE IF NOT EXISTS user (
    id INT(11) AUTO_INCREMENT COMMENT 'id',
    role VARCHAR(32) NOT NULL COMMENT 'admin, staff, other',
    lastname VARCHAR(32) NOT NULL COMMENT 'also surname or family_name',
    firstname VARCHAR(32) NOT NULL COMMENT 'also given name',
    middlename VARCHAR(64) COMMENT 'optional ',
    email VARCHAR(255) NOT NULL COMMENT 'maximum 254 characters',
    password CHAR(255) COMMENT 'MD5 hash',
    phone CHAR(16) NOT NULL COMMENT 'maximum 16 characters',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT 'created at',
    PRIMARY KEY (`id`),
    UNIQUE INDEX `by_email` (`email`) USING BTREE
)
ENGINE = MyISAM DEFAULT CHARSET = utf8;

CREATE TABLE IF NOT EXISTS course (
    id INT(11) AUTO_INCREMENT COMMENT 'id', 
    code CHAR(16) NOT NULL COMMENT 'code',
    name VARCHAR(64) COMMENT 'long name',
    description TEXT COMMENT 'some description',
    teacher INT COMMENT 'user id',
    fee INT COMMENT 'fee',
    timetable SET('Sunday', 'Monday', 'Tuesday', 'Wednesday', 'Thursday', 'Friday', 'Saturday'),
    start_date DATE COMMENT 'start date',
    end_date DATE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT 'created at',
    PRIMARY KEY (`id`),
    UNIQUE INDEX `by_code` (`code`) USING BTREE
)
ENGINE = InnoDB DEFAULT CHARSET = utf8;


CREATE TABLE IF NOT EXISTS student (
    id INT(11) AUTO_INCREMENT COMMENT 'student id', 
    customer_id INT(11) NOT NULL COMMENT 'customer id',
    course_id INT(11) NOT NULL COMMENT 'course id',
    paid_amount INT(11) DEFAULT 0 COMMENT 'paid for the course',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT 'created at',
    PRIMARY KEY (`id`),
    UNIQUE INDEX `by_customer` (`customer_id`) USING BTREE,
    UNIQUE INDEX `by_course` (`course_id`) USING BTREE
)
ENGINE = InnoDB DEFAULT CHARSET = utf8;