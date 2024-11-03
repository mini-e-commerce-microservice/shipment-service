CREATE TABLE shipping_addresses
(
    id                bigserial primary key,
    user_id           bigint       not null,
    address_source    varchar(100) not null,
    address_source_id varchar(100),
    name              varchar(255) not null,
    country           varchar(100) not null,
    country_code      varchar(100) not null,
    province          varchar(100) not null,
    city              varchar(100) not null,
    district          varchar(100) not null,
    address_note      text
)