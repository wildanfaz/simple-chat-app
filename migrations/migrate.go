package migrations

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func Migrate(ctx context.Context, db *pgx.Conn) {
	for _, v := range queries {
		_, err := db.Exec(ctx, v)
		if err != nil {
			panic(err)
		}
	}
}

var queries = []string{`
CREATE TABLE IF NOT EXISTS chats (
    chat_id BIGSERIAL PRIMARY KEY,
    room_id VARCHAR(100) NOT NULL,
    sender_id BIGINT NOT NULL,
    receiver_id BIGINT NOT NULL,
    message TEXT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ
)
`, `
CREATE INDEX IF NOT EXISTS chats_room_id_idx ON chats (room_id)
`, `
CREATE INDEX IF NOT EXISTS chats_sender_id_idx ON chats (sender_id)
`, `
CREATE INDEX IF NOT EXISTS chats_receiver_id_idx ON chats (receiver_id)
`, `
CREATE INDEX IF NOT EXISTS chats_message_idx ON chats (message)
`, `
DO $$ 
BEGIN 
    IF (SELECT COUNT(*) FROM pg_publication_tables WHERE pubname = 'supabase_realtime' AND tablename = 'chats') < 1 THEN 
        ALTER PUBLICATION supabase_realtime ADD TABLE chats; 
    END IF; 
END $$;
`,
}
