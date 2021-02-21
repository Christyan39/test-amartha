package database

import "test/amartha/usecase/model"

func (db *DB) GetShortenByCode(code string) (shorten *model.ShortlnRequest) {
	q := `
	SELECT 
		code,
		url,
		created_at,
		last_seen_at, 
		count
	FROM 
		shorten
	WHERE
		code = ?
	`

	rows, err := db.Master.Query(q, code)
	if err != nil {
		return nil
	}

	defer rows.Close()

	for rows.Next() {
		shorten = &model.ShortlnRequest{}
		err := rows.Scan(
			&shorten.Shortcode,
			&shorten.Url,
			&shorten.CreatedAt,
			&shorten.LastSeen,
			&shorten.Count,
		)
		if err != nil {
			return nil
		}
	}

	return shorten
}

func (db *DB) CreateShortenCode(shorten *model.ShortlnRequest) (err error) {
	q := `
	INSERT INTO shorten (code, url, created_at, last_seen_at, count) VALUES (?, ?, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, 0);
	`

	_, err = db.Master.Exec(q, shorten.Shortcode, shorten.Url)
	if err != nil {
		return err
	}

	return nil
}

func (db *DB) CountVisitingURL(code string) (err error) {
	q := `
	UPDATE shorten SET last_seen_at = CURRENT_TIMESTAMP, count = count+1 WHERE code = ?
	`

	_, err = db.Master.Exec(q, code)
	if err != nil {
		return err
	}

	return nil
}
