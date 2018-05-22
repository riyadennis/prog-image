-- noinspection SqlDialectInspectionForFile

-- noinspection SqlNoDataSourceInspectionForFile

CREATE TABLE IF NOT EXISTS images(id varchar(100) NOT NULL PRIMARY KEY,source varchar(500),InsertedDatetime DATETIME);
CREATE TABLE IF NOT EXISTS image_types(id varchar(100) NOT NULL,image_type varchar(100) NOT NULL ,InsertedDatetime DATETIME);
