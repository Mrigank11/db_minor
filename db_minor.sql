drop database db_minor;
create database db_minor;
use db_minor;

create table role(
	role_id int auto_increment primary key not null,
    display_name varchar(100)
);
	
create table users(
	username varchar(100) primary key not null,
	first_name varchar(100) ,
	last_name varchar(100) ,
	mobile bigint ,
	email varchar(100) ,
	avatar varchar(100) ,
	password varchar(100) ,
	role_id int not null ,
	foreign key(role_id) references role(role_id)
);

create table products(
	product_id int auto_increment primary key not null,
    name varchar(100) ,
    description text 
);
    
create table tag(
	tag_id int auto_increment primary key not null,
    name varchar(100) ,
    description text 
);
    
create table product_tags(
	product_id int not null ,
    tag_id int not null ,
    foreign key(product_id) references products(product_id) ,
	foreign key(tag_id) references tag(tag_id) 
);
    
create  table stock(
	stock_id int auto_increment primary key not null,
	product_id int not null ,
    price int ,
    dealer varchar(100) ,
    foreign key(product_id) references products(product_id)
);
    
create  table sale(
	transaction_id int not null primary key,
	amount int ,
    sale_date datetime
);
    
create  table cart(
	cart_id int null ,
    transaction_id int ,
    user_id varchar(100) not null,
    foreign key(transaction_id) references sale(transaction_id) ,
    foreign key(user_id) references users(username)
);
    
insert into role(display_name)  values
	('admin'),
    ('customer');

insert into users(username, password, role_id) values("admin", "$2a$10$K3pP8EONwz1bQJ/AyFImu.T.uNa7VeY/8LFGHltcRjDKg002sob16", 1);

insert into products values
	(NULL, "Pen", "a pen"),
    (NULL, "Pencil", "a pencil"),
    (NULL, "Rubber", "a rubber");
insert into stock values
	(NULL, 1 , 40, "amazon"),
    (NULL, 1 , 30, "azon"),
    (NULL, 2 , 40, "amazon");
    
-- query for listing products with best price
select p.*, s.price, s.dealer, count(s1.price) from products p left join stock s on s.stock_id = (
	select stock_id from stock s_
    where s_.product_id = p.product_id
    order by s_.price asc
    limit 1
) 
left join stock s1 on s1.product_id = p.product_id
group by p.product_id;