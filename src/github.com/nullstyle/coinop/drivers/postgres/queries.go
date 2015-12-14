package postgres

// Queries is a helper that provides easy access to SQL used by coinop
var Queries struct {
	KV struct {
		Load   string
		Insert string
		Update string
	}
	Delivery struct {
		Insert         string
		MarkSuccessful string
		MarkFailed     string
		SelectFailed   string
		Start          string
	}
	Webhook struct {
		All            string
		ForDestination string
		Insert         string
		Delete         string
	}
}

func init() {
	Queries.KV.Load = `
	SELECT key, value FROM coinop.kv WHERE key = $1
	`

	Queries.KV.Insert = `
	INSERT INTO coinop.kv (key, value) VALUES($1, $2)
	RETURNING key
	`

	Queries.KV.Update = `
	UPDATE coinop.kv SET value = $2 WHERE key = $1
	`

	Queries.Delivery.Insert = `
		INSERT INTO coinop.deliveries (
			url,
	    payment_paging_token,
	    payment_from,
	    payment_to,
	    payment_memo_type,
	    payment_memo,
	    payment_asset_type,
	    payment_asset_code,
	    payment_asset_issuer,
	    payment_amount,
	    created_at,
	    updated_at
		) VALUES (
			:url,
			:payment_paging_token,
			:payment_from,
			:payment_to,
			:payment_memo_type,
			:payment_memo,
			:payment_asset_type,
			:payment_asset_code,
			:payment_asset_issuer,
			:payment_amount,
			:created_at,
			:updated_at
		)
		RETURNING id;
	`

	Queries.Delivery.MarkFailed = `
		UPDATE coinop.deliveries SET
		last_failed_at = $1
		updated_at = $1
		version = version + 1
		WHERE id = $2
		AND version = $3
		RETURNING version
	`

	Queries.Delivery.MarkSuccessful = `
		UPDATE coinop.deliveries SET
		succeeded_at = $1
		updated_at = $1
		version = version + 1
		WHERE id = $2
		AND version = $3
		RETURNING version
	`

	Queries.Delivery.SelectFailed = `
		SELECT *
		FROM coinop.deliveries
		WHERE last_failed_at IS NOT NULL
		AND succeeded_at IS NULL
		AND (
			started_at IS NULL
		OR age(started_at) > interval '1 minute'
		)
	`

	Queries.Delivery.Start = `
		UPDATE coinop.deliveries SET
		started_at = $1
		updated_at = $1
		version = version + 1
		WHERE id = $2
		AND (
			started_at IS NULL
		OR age(started_at) > interval '1 minute'
		)
		RETURNING version
	`

	Queries.Webhook.All = `
    SELECT *
    FROM coinop.webhooks
    ORDER BY id ASC
  `

	Queries.Webhook.ForDestination = `
    SELECT *
    FROM coinop.webhooks
		WHERE destination_filter = $1
    ORDER BY id ASC
  `

	Queries.Webhook.Insert = `
    INSERT INTO
    coinop.webhooks (url, destination_filter, memo_type_filter, memo_filter,
      created_at)
    VALUES ($1, $2, $3, $4, $5)
    RETURNING id;
  `

	Queries.Webhook.Delete = `
    DELETE FROM coinop.webhooks WHERE id = $1
  `
}
