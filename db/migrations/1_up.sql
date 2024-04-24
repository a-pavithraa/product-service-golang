create sequence product_id_seq start with 1 increment by 50;

create table products
(
    id bigint default nextval('product_id_seq') not null,
    name        text not null unique,
    description text,
    price       numeric not null,
    primary key (id)
);
