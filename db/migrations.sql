create sequence product_id_seq start with 1 increment by 50;

create table products
(
    id bigint default nextval('product_id_seq') not null,
    name        text not null unique,
    description text,
    price       numeric not null,
    primary key (id)
);
insert into products( name, description, price) values
('Product 1','Product 1', 34.0),
('Product 2','Product 2', 45.40),
('Product 3','Product 3', 44.50),
('Product 4', 'Product 4',44.50)
;