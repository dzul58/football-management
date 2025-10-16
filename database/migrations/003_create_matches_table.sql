-- Migration: Create matches table
-- Description: Tabel untuk menyimpan jadwal dan hasil pertandingan

-- Create ENUM type for match status
CREATE TYPE match_status AS ENUM ('Scheduled', 'Completed', 'Cancelled');

CREATE TABLE IF NOT EXISTS matches (
    id SERIAL PRIMARY KEY,
    match_date DATE NOT NULL,
    match_time TIME NOT NULL,
    home_team_id INTEGER NOT NULL,
    away_team_id INTEGER NOT NULL,
    home_score INTEGER NULL DEFAULT NULL, -- NULL jika belum dimainkan
    away_score INTEGER NULL DEFAULT NULL, -- NULL jika belum dimainkan
    status match_status DEFAULT 'Scheduled',
    deleted_at TIMESTAMP NULL DEFAULT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (home_team_id) REFERENCES teams(id) ON DELETE CASCADE,
    FOREIGN KEY (away_team_id) REFERENCES teams(id) ON DELETE CASCADE,
    CHECK (home_team_id != away_team_id)
);

CREATE INDEX IF NOT EXISTS idx_matches_match_date ON matches(match_date);
CREATE INDEX IF NOT EXISTS idx_matches_status ON matches(status);
CREATE INDEX IF NOT EXISTS idx_matches_home_team ON matches(home_team_id);
CREATE INDEX IF NOT EXISTS idx_matches_away_team ON matches(away_team_id);
CREATE INDEX IF NOT EXISTS idx_matches_deleted_at ON matches(deleted_at);

-- Trigger untuk auto-update updated_at
CREATE TRIGGER update_matches_updated_at BEFORE UPDATE ON matches
FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();

