-- +goose Up
-- +goose StatementBegin
CREATE UNIQUE INDEX `filename_index` ON zip_logs (`filename`) USING BTREE;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX `filename_index` ON zip_logs
-- +goose StatementEnd
