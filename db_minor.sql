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
	transaction_id varchar(10) not null primary key,
	amount int ,
    sale_date timestamp default current_timestamp
);
    
create  table cart(
	cart_id int auto_increment primary key not null ,
    transaction_id varchar(10),
    user_id varchar(100) not null,
    foreign key(transaction_id) references sale(transaction_id) ,
    foreign key(user_id) references users(username)
);

create table cart_items(
	cart_id int not null ,
    product_id int not null ,
    added_on datetime DEFAULT CURRENT_TIMESTAMP,
    foreign key(product_id) references products(product_id) ,
	foreign key(cart_id) references cart(cart_id),
    PRIMARY KEY( `cart_id`, `product_id`)
);
    
-- DATA DUMP
insert into role(display_name)  values
	('admin'),
    ('customer');

insert into users(username, password, role_id) values("admin", "$2a$10$K3pP8EONwz1bQJ/AyFImu.T.uNa7VeY/8LFGHltcRjDKg002sob16", 1);
insert into cart values(NULL, NULL, "admin");

insert into products values
	(NULL, "Pen", "a pen"),
    (NULL, "Pencil", "a pencil"),
    (NULL, "Rubber", "a rubber"),
    (NULL, "Burger", "a big burger.");
insert into stock values
	(NULL, 1, 40, "amazon"),
    (NULL, 1, 30, "azon"),
    (NULL, 2, 40, "amazon"),
    (NULL, 4, 10, "bakery");
    
insert into tag values
	(NULL, "stationary", "some description for stationary"),
    (NULL, "edible", "to fill tummy"),
    (NULL, "unhealthy", "don't eat it!");

insert into product_tags values
	(1,1),(2,1), (3,1) ,(4,2),(4,3);
    
-- view for listing products with best price
create view products_with_best_prices as 
select p.*, s.price, s.dealer, count(s1.price), s.stock_id from products p left join stock s on s.stock_id = (
	select stock_id from stock s_
    where s_.product_id = p.product_id
    order by s_.price asc
    limit 1
) 
left join stock s1 on s1.product_id = p.product_id
group by p.product_id;

-- query for listing items in a cart
select p.* from cart_items c inner join products_with_best_prices p on p.product_id = c.product_id and cart_id=1 ;
-- query to show all tags for an item
select t.* from product_tags pt 
join products p on p.product_id = pt.product_id and p.product_id = 3
join tag t on t.tag_id = pt.tag_id;
-- query to show all products with a certain tag
select ps.* from products_with_best_prices ps
join product_tags p on p.product_id = ps.product_id
join tag t on t.tag_id = p.tag_id and t.tag_id = 1;

-- query to get sum of items in a cart
select sum(p.price) from cart_items c inner join products_with_best_prices p on p.product_id = c.product_id and cart_id=1 ;

-- query to get all past orders
select s.* from sale s
join cart c on s.transaction_id = c.transaction_id
join users u on u.username = c.user_id and u.username = "admin" ;

-- query to get cart with TID
select p.*, s.amount from sale s
join cart c on s.transaction_id = c.transaction_id and c.transaction_id = "MRAjWwhTHc"
join cart_items ci on ci.cart_id = c.cart_id
join products p on p.product_id = ci.product_id;

-- trigger for handling cart checkout
delimiter $$
create trigger cart_checkout
before update on cart
for each row
begin
-- declare
declare available_item_count, wanted_item_count, amount int;
-- check if all items are available
select count(distinct(s.product_id)) into available_item_count from cart_items ci
join stock s on s.product_id = ci.product_id and ci.cart_id=NEW.cart_id;

select count(distinct(ci.product_id)) into wanted_item_count from cart_items ci where ci.cart_id=NEW.cart_id;

if available_item_count < wanted_item_count then
	signal sqlstate '45000' set message_text = 'Some items are not available';	
else
	select sum(pp.price) into amount from cart_items ci
	join products_with_best_prices pp on ci.product_id = pp.product_id and cart_id=NEW.cart_id;
    
    -- delete stock
	delete s from stock s, products_with_best_prices p, cart_items ci 
	where s.stock_id = p.stock_id
	and ci.product_id = p.product_id and ci.cart_id=NEW.cart_id;
    
	-- create sale
	insert into sale(transaction_id, amount) values(NEW.transaction_id, amount);
end if;
end$$
delimiter ;