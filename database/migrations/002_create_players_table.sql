-- Migration: Create players table
-- Description: Tabel untuk menyimpan informasi pemain sepak bola

-- Create ENUM type for position
CREATE TYPE player_position AS ENUM ('Penyerang', 'Gelandang', 'Bertahan', 'Penjaga Gawang');

CREATE TABLE IF NOT EXISTS players (
    id SERIAL PRIMARY KEY,
    team_id INTEGER NOT NULL,
    name VARCHAR(255) NOT NULL,
    height DECIMAL(5,2) NOT NULL, -- Tinggi badan dalam cm
    weight DECIMAL(5,2) NOT NULL, -- Berat badan dalam kg
    position player_position NOT NULL,
    jersey_number INTEGER NOT NULL,
    deleted_at TIMESTAMP NULL DEFAULT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (team_id) REFERENCES teams(id) ON DELETE CASCADE
);

CREATE UNIQUE INDEX IF NOT EXISTS unique_jersey_per_team ON players(team_id, jersey_number) WHERE deleted_at IS NULL;
CREATE INDEX IF NOT EXISTS idx_players_team_id ON players(team_id);
CREATE INDEX IF NOT EXISTS idx_players_position ON players(position);
CREATE INDEX IF NOT EXISTS idx_players_deleted_at ON players(deleted_at);

-- Trigger untuk auto-update updated_at
CREATE TRIGGER update_players_updated_at BEFORE UPDATE ON players
FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();

