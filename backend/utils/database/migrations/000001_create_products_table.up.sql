create table if not exists products (
    id SERIAL PRIMARY KEY,
    nama_obat varchar(255) not null,
    stok int,
    harga decimal(12, 2)
);

create table if not exists orders (
    id SERIAL PRIMARY KEY,
    product_id int not null references products(id),
    quantity int,
    subtotal decimal(12, 2),
    discount_percent decimal(5, 2),
    total decimal(12, 2),
    created_at timestamp default current_timestamp
);