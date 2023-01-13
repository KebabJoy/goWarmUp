-- +goose Up
-- +goose StatementBegin
// success добавляли изначально, думал буду делать асинхронную архивацию файлов.
// Чтобы потом по ID джобы можно было получить файл либо узнать, что еще не запроцессилось
create table zip_logs (id integer NOT NULL AUTO_INCREMENT, success BOOLEAN, filename VARCHAR(255), PRIMARY KEY(id));
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table zip_logs;
-- +goose StatementEnd
