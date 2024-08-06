alter table users
    add column if not exists created_at timestamp default now(),
    add column if not exists updated_at timestamp,
    add column if not exists deleted_at integer default 0;

alter table stocks
    add column if not exists created_at timestamp default now(),
    add column if not exists updated_at timestamp,
    add column if not exists deleted_at integer default 0;

alter table market_updates
    add column if not exists created_at timestamp default now(),
    add column if not exists updated_at timestamp,
    add column if not exists deleted_at integer default 0;

alter table orders
    add column if not exists created_at timestamp default now(),
    add column if not exists updated_at timestamp,
    add column if not exists deleted_at integer default 0;

alter table order_book
    add column if not exists created_at timestamp default now(),
    add column if not exists updated_at timestamp,
    add column if not exists deleted_at integer default 0;

alter table portfolios
    add column if not exists created_at timestamp default now(),
    add column if not exists updated_at timestamp,
    add column if not exists deleted_at integer default 0;


alter table trade_confirmations
    add column if not exists created_at timestamp default now(),
    add column if not exists updated_at timestamp,
    add column if not exists deleted_at integer default 0;

alter table market_news
    add column if not exists created_at timestamp default now(),
    add column if not exists updated_at timestamp,
    add column if not exists deleted_at integer default 0;


alter table account_information
    add column if not exists created_at timestamp default now(),
    add column if not exists updated_at timestamp,
    add column if not exists deleted_at integer default 0;

alter table fund_transfers
    add column if not exists created_at timestamp default now(),
    add column if not exists updated_at timestamp,
    add column if not exists deleted_at integer default 0;
