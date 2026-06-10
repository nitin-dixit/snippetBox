-- +goose Up
-- +goose StatementBegin
create table if not exists sessions(
token char(43) primary key,
  data bytea not null,
  expiry timestamp(6) not null
);

create index sessions_expiry_idx on sessions (expiry);
-- +goose StatementEnd
-- +goose Down
drop table sessions;
