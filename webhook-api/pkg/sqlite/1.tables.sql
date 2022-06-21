CREATE TABLE IF NOT EXISTS "buckets" (
			"id"	TEXT NOT NULL UNIQUE,
			"bucket_name"	TEXT NOT NULL UNIQUE,
			"created_at"	INTEGER NOT NULL,
			PRIMARY KEY("id")
		);
		CREATE TABLE IF NOT EXISTS "records" (
			"id"	TEXT NOT NULL UNIQUE,
			"bucket_id"	TEXT NOT NULL,
			"headers"	TEXT,
			"body"	TEXT,
			"created_at"	INTEGER NOT NULL,
			FOREIGN KEY("bucket_id") REFERENCES "buckets"("id"),
			PRIMARY KEY("id")
		);