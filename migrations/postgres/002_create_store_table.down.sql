drop table if exists store;

alter table users
    drop column if exists created_at,
    drop column if exists updated_at,
    drop column if exists deleted_at;

alter table stocks
    drop column if exists created_at,
    drop column if exists updated_at,
    drop column if exists deleted_at;

alter table market_updates
    drop column if exists created_at,
    drop column if exists updated_at,
    drop column if exists deleted_at;

alter table orders
    drop column if exists created_at,
    drop column if exists updated_at,
    drop column if exists deleted_at;

alter table order_book
    drop column if exists created_at,
    drop column if exists updated_at,
    drop column if exists deleted_at;

alter table portfolios
    drop column if exists created_at,
    drop column if exists updated_at,
    drop column if exists deleted_at;

alter table trade_confirmations
    drop column if exists created_at,
    drop column if exists updated_at,
    drop column if exists deleted_at;

alter table market_news
    drop column if exists created_at,
    drop column if exists updated_at,
    drop column if exists deleted_at;

alter table account_information
    drop column if exists created_at,
    drop column if exists updated_at,
    drop column if exists deleted_at;

alter table fund_transfers
    drop column if exists created_at,
    drop column if exists updated_at,
    drop column if exists deleted_at;
