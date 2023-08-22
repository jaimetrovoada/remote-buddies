-- +goose Up
-- +goose StatementBegin
ALTER TABLE "User" ADD COLUMN "interests" varchar ARRAY[5];
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE "User" DROP COLUMN "interests";
-- +goose StatementEnd
