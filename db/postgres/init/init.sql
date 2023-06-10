-- すでに存在するスキーマ(データベース)を削除する
-- CASCADEを指定することで、依存関係にあるオブジェクトも削除する
DROP SCHEMA IF EXISTS "portfolio" CASCADE;

-- スキーマを作成する
CREATE SCHEMA "portfolio";

-- スキーマを使用する
SET search_path TO "portfolio";

-- 実行確認用テーブルを作成する
CREATE TABLE "test" (
  "id" VARCHAR(32) PRIMARY KEY NOT NULL,
  "name" VARCHAR(255) NOT NULL
);

-- 実行確認用データを挿入する
INSERT INTO "test" ("id", "name") VALUES ('1', 'テスト1');
INSERT INTO "test" ("id", "name") VALUES ('2', 'テスト2');
INSERT INTO "test" ("id", "name") VALUES ('3', 'テスト3');
