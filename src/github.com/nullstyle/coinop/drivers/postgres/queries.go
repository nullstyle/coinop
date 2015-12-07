package postgres

// Queries is a helper that provides easy access to SQL used by coinop
var Queries struct {
	Webhook struct {
		All    string
		Insert string
		Delete string
	}
}

func init() {
	Queries.Webhook.All = `
    SELECT *
    FROM webhooks
    ORDER BY id ASC
  `

	Queries.Webhook.Insert = `
    INSERT INTO
    webhooks (url, destination_filter, memo_type_filter, memo_filter,
      created_at)
    VALUES ($1, $2, $3, $4, $5)
    RETURNING id;
  `

	Queries.Webhook.Delete = `
    DELETE FROM webhooks WHERE id = $1
  `
}
