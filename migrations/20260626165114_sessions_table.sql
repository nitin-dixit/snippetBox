-- +goose Up
-- +goose StatementBegin
create table if not exists sessions (
    token TEXT primary key,
    data BYTEA not null,
    expiry TIMESTAMPTZ not null
);

create index concurrently sessions_expiry_idx on sessions (expiry);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop index concurrently if exists sessions_expiry_idx;
drop table if exists sessions;
-- +goose StatementEnd
