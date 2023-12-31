drop table if exists messages;
drop domain if exists address;
drop domain if exists uint256;
drop domain if exists signature;
create domain address as bytea check(octet_length(value) = 20);
create domain uint256 as bytea check(octet_length(value) <= 32);
create domain signature as bytea check(octet_length(value) = 65);
create table messages (
  signer address not null,
  nonce uint256 not null,
  gas_order uint256 not null,
  on_behalf address not null,
  deadline uint256 not null,
  endpoint address not null,
  gas uint256 not null,
  data bytea not null,
  orig_sign signature not null,
  valid_sign signature not null unique,
  id serial primary key not null
);