CREATE TABLE user (
  email varchar(128) UNIQUE, 
  password varchar(128),
  first_name varchar(256),
  last_name varchar(256), 
  org_name varchar(256),
  inst varchar(256) default '',
  build_no varchar(128) default '',
  floor_no varchar(256) default '',
  lab_head varchar(128) default '',
  lab_address varchar(128) default '',
  tel varchar(128) default '');