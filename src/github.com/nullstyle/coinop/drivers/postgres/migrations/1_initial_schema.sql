-- +migrate Up
CREATE TABLE coinop.webhooks (
    id bigserial PRIMARY KEY,
    url character varying(1024) NOT NULL,
    destination_filter character varying(255) NOT NULL,
    memo_type_filter character varying(255),
    memo_filter character varying(255),
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone
);
CREATE INDEX ON coinop.webhooks (url);

CREATE TABLE coinop.deliveries (
    id bigserial PRIMARY KEY,
    version bigint NOT NULL DEFAULT 1,
    url character varying(1024) NOT NULL,

    payment_paging_token character varying(255) NOT NULL,
    payment_from character varying(255) NOT NULL,
    payment_to character varying(255) NOT NULL,
    payment_memo_type character varying(255) NOT NULL,
    payment_memo character varying(255) NOT NULL,
    payment_asset_type character varying(255) NOT NULL,
    payment_asset_code character varying(255),
    payment_asset_issuer character varying(255),
    payment_amount bigint NOT NULL,

    started_at timestamp without time zone,
    last_failed_at timestamp without time zone,
    succeeded_at timestamp without time zone,

    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone
);

CREATE TABLE coinop.kv (
  key character varying(255) NOT NULL,
  value jsonb
);
CREATE UNIQUE INDEX ON coinop.kv (key);

-- +migrate Down
DROP TABLE coinop.kv;
DROP TABLE coinop.deliveries;
DROP TABLE coinop.webhooks;
