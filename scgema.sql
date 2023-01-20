create database max_invetory;
use max_invetory;

create table USERS(
    id int not null auto_increment,
    username varchar(25) not null,
    name varchar(25) not null,
    password varchar(25) not null,
    primary key (id)
);

create table PRODUCTS(
    id int not null auto_increment,
    name varchar(25) not null,
    description varchar(25) not null,
    price float not null,
    created_by int not null,
    primary key (id),
    foreing key (created_by) references USERS(id)
);

create table ROLES(
    id int not null auto_increment,
    name varchar(25) not null,
    primary key (id)
    
);

create table USER_ROLES(
    id int not null auto_increment
    user_id int not null,
    role_id int not null,
    primary key (id)
    foreign (user_id) references USERS(id)
    foreign (role_id) references ROLES(id)    
);

user_id int not null