CREATE TABLE "accounts" (
  "id" bigserial PRIMARY KEY,
  "owner" varchar NOT NULL,
  "balance" bigint NOT NULL,
  "currency" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT NOW()
);

CREATE TABLE "entries" (
  "id" bigserial PRIMARY KEY,
  "acc_id" bigint NOT NULL,
  "amount" bigint NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT NOW()
);

CREATE TABLE "transfers" (
  "id" bigserial PRIMARY KEY,
  "from_acc_id" bigint NOT NULL,
  "to_acc_id" bigint NOT NULL,
  "amount" bigint NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT NOW()
);

-- Create indexes for performance
CREATE INDEX idx_accounts_owner ON "accounts" ("owner");
CREATE INDEX idx_entries_acc_id ON "entries" ("acc_id");
CREATE INDEX idx_transfers_from_acc_id ON "transfers" ("from_acc_id");
CREATE INDEX idx_transfers_to_acc_id ON "transfers" ("to_acc_id");
CREATE INDEX idx_transfers_from_to_acc_id ON "transfers" ("from_acc_id", "to_acc_id");

-- Add a comment to the 'amount' column in 'entries'
COMMENT ON COLUMN "entries"."amount" IS 'Can be positive or negative.';

-- Add foreign key constraints
ALTER TABLE "entries" 
ADD CONSTRAINT fk_entries_acc_id FOREIGN KEY ("acc_id") REFERENCES "accounts" ("id");

ALTER TABLE "transfers" 
ADD CONSTRAINT fk_transfers_from_acc_id FOREIGN KEY ("from_acc_id") REFERENCES "accounts" ("id");

ALTER TABLE "transfers" 
ADD CONSTRAINT fk_transfers_to_acc_id FOREIGN KEY ("to_acc_id") REFERENCES "accounts" ("id");
