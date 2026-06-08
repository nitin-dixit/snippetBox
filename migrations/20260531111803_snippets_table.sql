-- +goose Up
-- +goose StatementBegin
create table if not exists snippets(
  id bigserial primary key,
title varchar(100) not null,
content text not null,
created timestamptz not null default current_timestamp,
expires timestamptz not null

)

-- +goose StatementEnd
-- +goose Down
drop table snippets;
