CREATE TABLE "account_balances" (
  "account_id" varchar(50) NOT NULL,
  "user_id" varchar(50) DEFAULT NULL,
  "amount" decimal(15,2) DEFAULT NULL,
  "dummy_col_4" varchar(255) DEFAULT NULL,
  PRIMARY KEY ("account_id")
);

CREATE TABLE "account_details" (
  "account_id" varchar(50) NOT NULL,
  "user_id" varchar(50) DEFAULT NULL,
  "color" varchar(10) DEFAULT NULL,
  "is_main_account" boolean DEFAULT NULL,
  "progress" int DEFAULT NULL,
  "dummy_col_5" varchar(255) DEFAULT NULL,
  PRIMARY KEY ("account_id")
);

CREATE TABLE "account_flags" (
  "flag_id" SERIAL PRIMARY KEY,
  "account_id" varchar(50) NOT NULL,
  "user_id" varchar(50) NOT NULL,
  "flag_type" varchar(50) NOT NULL,
  "flag_value" varchar(30) NOT NULL,
  "created_at" timestamp DEFAULT (now()),
  "updated_at" timestamp DEFAULT (now())
);

CREATE TABLE "accounts" (
  "account_id" varchar(50) NOT NULL,
  "user_id" varchar(50) DEFAULT NULL,
  "type" varchar(50) DEFAULT NULL,
  "currency" varchar(10) DEFAULT NULL,
  "account_number" varchar(20) DEFAULT NULL,
  "issuer" varchar(100) DEFAULT NULL,
  "dummy_col_3" varchar(255) DEFAULT NULL,
  PRIMARY KEY ("account_id")
);

CREATE TABLE "banners" (
  "banner_id" varchar(50) NOT NULL,
  "user_id" varchar(50) DEFAULT NULL,
  "title" varchar(255) DEFAULT NULL,
  "description" text,
  "image" varchar(255) DEFAULT NULL,
  "dummy_col_11" varchar(255) DEFAULT NULL,
  PRIMARY KEY ("banner_id")
);

CREATE TABLE "debit_card_design" (
  "card_id" varchar(50) NOT NULL,
  "user_id" varchar(50) DEFAULT NULL,
  "color" varchar(10) DEFAULT NULL,
  "border_color" varchar(10) DEFAULT NULL,
  "dummy_col_9" varchar(255) DEFAULT NULL,
  PRIMARY KEY ("card_id")
);

CREATE TABLE "debit_card_details" (
  "card_id" varchar(50) NOT NULL,
  "user_id" varchar(50) DEFAULT NULL,
  "issuer" varchar(100) DEFAULT NULL,
  "number" varchar(25) DEFAULT NULL,
  "dummy_col_10" varchar(255) DEFAULT NULL,
  PRIMARY KEY ("card_id")
);

CREATE TABLE "debit_card_status" (
  "card_id" varchar(50) NOT NULL,
  "user_id" varchar(50) DEFAULT NULL,
  "status" varchar(20) DEFAULT NULL,
  "dummy_col_8" varchar(255) DEFAULT NULL,
  PRIMARY KEY ("card_id")
);

CREATE TABLE "debit_cards" (
  "card_id" varchar(50) NOT NULL,
  "user_id" varchar(50) DEFAULT NULL,
  "name" varchar(100) DEFAULT NULL,
  "dummy_col_7" varchar(255) DEFAULT NULL,
  PRIMARY KEY ("card_id")
);

CREATE TABLE "transactions" (
  "transaction_id" varchar(50) NOT NULL,
  "user_id" varchar(50) DEFAULT NULL,
  "name" varchar(100) DEFAULT NULL,
  "image" varchar(255) DEFAULT NULL,
  "is_bank" boolean DEFAULT NULL,
  "dummy_col_6" varchar(255) DEFAULT NULL,
  PRIMARY KEY ("transaction_id")
);

CREATE TABLE "user_greetings" (
  "user_id" varchar(50) NOT NULL,
  "greeting" text,
  "dummy_col_2" varchar(255) DEFAULT NULL,
  PRIMARY KEY ("user_id")
);

CREATE TABLE "users" (
  "user_id" varchar(50) NOT NULL,
  "name" varchar(100) DEFAULT NULL,
  "dummy_col_1" varchar(255) DEFAULT NULL,
  PRIMARY KEY ("user_id")
);

INSERT INTO "account_balances" ("account_id", "user_id", "amount", "dummy_col_4") VALUES
('0000016ce1a211ef95a30242ac180002', 'fffff95ee1a111ef95a30242ac180002', 70235.84, 'dummy_value_4'),
('000020ece1a211ef95a30242ac180002', '000018b0e1a211ef95a30242ac180002', 86048.06, 'dummy_value_4'),
('000024eae1a211ef95a30242ac180002', '000018b0e1a211ef95a30242ac180002', 93311.34, 'dummy_value_4');

INSERT INTO "account_details" ("account_id", "user_id", "color", "is_main_account", "progress", "dummy_col_5") VALUES
('0000016ce1a211ef95a30242ac180002', 'fffff95ee1a111ef95a30242ac180002', '#24c875', FALSE, 36, 'dummy_value_5'),
('000020ece1a211ef95a30242ac180002', '000018b0e1a211ef95a30242ac180002', '#24c875', TRUE, 42, 'dummy_value_5'),
('000024eae1a211ef95a30242ac180002', '000018b0e1a211ef95a30242ac180002', '#24c875', FALSE, 82, 'dummy_value_5');

INSERT INTO "account_flags" ("flag_id", "account_id", "user_id", "flag_type", "flag_value", "created_at", "updated_at") VALUES
(1, '5a5b0fafe1a111ef80ca0242ac180002', 'fffff95ee1a111ef95a30242ac180002', 'system', 'Flag3', '2025-02-02 20:07:39', '2025-02-02 20:07:39'),
(2, '000020ece1a211ef95a30242ac180002', '000018b0e1a211ef95a30242ac180002', 'system', 'Flag4', '2025-02-02 20:12:17', '2025-02-02 20:12:17'),
(3, '000020ece1a211ef95a30242ac180002', '000018b0e1a211ef95a30242ac180002', 'system', 'Flag5', '2025-02-02 20:12:17', '2025-02-02 20:12:17');

INSERT INTO "accounts" ("account_id", "user_id", "type", "currency", "account_number", "issuer", "dummy_col_3") VALUES
('0000016ce1a211ef95a30242ac180002', 'fffff95ee1a111ef95a30242ac180002', 'saving-account', 'THB', '568-2-4732', 'TestLab', 'dummy_value_3'),
('000020ece1a211ef95a30242ac180002', '000018b0e1a211ef95a30242ac180002', 'saving-account', 'THB', '568-2-62729', 'TestLab', 'dummy_value_3'),
('000024eae1a211ef95a30242ac180002', '000018b0e1a211ef95a30242ac180002', 'saving-account', 'THB', '568-2-94760', 'TestLab', 'dummy_value_3');

INSERT INTO "banners" ("banner_id", "user_id", "title", "description", "image", "dummy_col_11") VALUES
('000018cfe1a211ef95a30242ac180002', '000018b0e1a211ef95a30242ac180002', 'Want some money?', 'You can start applying', 'https://dummyimage.com/54x54/999/fff', 'dummy_value_11'),
('000043d4e1a211ef95a30242ac180002', 'fffff95ee1a111ef95a30242ac180002', 'Want some money?', 'You can start applying', 'https://dummyimage.com/54x54/999/fff', 'dummy_value_11');

INSERT INTO "debit_card_design" ("card_id", "user_id", "color", "border_color", "dummy_col_9") VALUES
('000018cfe1a211ef95a30242ac180002', '000018b0e1a211ef95a30242ac180002', '#00a1e2', '#ffffff', 'dummy_value_9'),
('000043d4e1a211ef95a30242ac180002', 'fffff95ee1a111ef95a30242ac180002', '#00a1e2', '#ffffff', 'dummy_value_9');

INSERT INTO "debit_card_details" ("card_id", "user_id", "issuer", "number", "dummy_col_10") VALUES
('000018cfe1a211ef95a30242ac180002', '000018b0e1a211ef95a30242ac180002', 'TestLab', '6821 5668 7876 2379', 'dummy_value_10'),
('000043d4e1a211ef95a30242ac180002', 'fffff95ee1a111ef95a30242ac180002', 'TestLab', '1895 6835 8492 1957', 'dummy_value_10');

INSERT INTO "debit_card_status" ("card_id", "user_id", "status", "dummy_col_8") VALUES
('000018cfe1a211ef95a30242ac180002', '000018b0e1a211ef95a30242ac180002', 'Active', 'dummy_value_8'),
('000043d4e1a211ef95a30242ac180002', 'fffff95ee1a111ef95a30242ac180002', 'Active', 'dummy_value_8');

INSERT INTO "debit_cards" ("card_id", "user_id", "name", "dummy_col_7") VALUES
('000018cfe1a211ef95a30242ac180002', '000018b0e1a211ef95a30242ac180002', 'My Debit Card', 'dummy_value_7'),
('000043d4e1a211ef95a30242ac180002', 'fffff95ee1a111ef95a30242ac180002', 'My Debit Card', 'dummy_value_7');

INSERT INTO "transactions" ("transaction_id", "user_id", "name", "image", "is_bank", "dummy_col_6") VALUES
('000018c1e1a211ef95a30242ac180002', '000018b0e1a211ef95a30242ac180002', 'Transaction_135017', 'https://dummyimage.com/54x54/999/fff', FALSE, 'dummy_value_6'),
('000043c6e1a211ef95a30242ac180002', 'fffff95ee1a111ef95a30242ac180002', 'Transaction_135018', 'https://dummyimage.com/54x54/999/fff', FALSE, 'dummy_value_6');

INSERT INTO "user_greetings" ("user_id", "greeting", "dummy_col_2") VALUES
('000018b0e1a211ef95a30242ac180002', 'Hello User_000018b0e1a211ef95a30242ac180002', 'dummy_value_2'),
('fffff95ee1a111ef95a30242ac180002', 'Hello Hello User_fffff95ee1a111ef95a30242ac180002', 'dummy_value_2');

INSERT INTO "users" ("user_id", "name", "dummy_col_1") VALUES
('000018b0e1a211ef95a30242ac180002', 'User_000018b0e1a211ef95a30242ac180002', 'dummy_value_1'),
('fffff95ee1a111ef95a30242ac180002', 'User_fffff95ee1a111ef95a30242ac180002', 'dummy_value_1');
