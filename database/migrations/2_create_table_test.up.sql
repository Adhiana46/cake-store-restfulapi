-- CUMA TEST TABLE UNTUK TESTING MIGRATIONs
CREATE TABLE test (
    id INT(11) NOT NULL AUTO_INCREMENT,
    title VARCHAR(100),
    description	VARCHAR(255),
    rating FLOAT DEFAULT 0,
    image VARCHAR(255),
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (id)
);