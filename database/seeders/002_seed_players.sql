-- Seeder: Sample players data
-- Description: Data contoh untuk pemain sepak bola

-- Players untuk Persija Jakarta (team_id = 1)
INSERT INTO players (team_id, name, height, weight, position, jersey_number) VALUES
(1, 'Andritany Ardhiyasa', 183.00, 75.00, 'Penjaga Gawang'::player_position, 1),
(1, 'Rizky Ridho', 178.00, 70.00, 'Bertahan'::player_position, 5),
(1, 'Marko Simic', 190.00, 85.00, 'Penyerang'::player_position, 9),
(1, 'Riko Simanjuntak', 172.00, 68.00, 'Gelandang'::player_position, 10)
ON CONFLICT DO NOTHING;

-- Players untuk Persib Bandung (team_id = 2)
INSERT INTO players (team_id, name, height, weight, position, jersey_number) VALUES
(2, 'Teja Paku Alam', 185.00, 78.00, 'Penjaga Gawang'::player_position, 1),
(2, 'Nick Kuipers', 188.00, 82.00, 'Bertahan'::player_position, 4),
(2, 'Beckham Putra', 175.00, 70.00, 'Gelandang'::player_position, 7),
(2, 'David da Silva', 180.00, 75.00, 'Penyerang'::player_position, 10)
ON CONFLICT DO NOTHING;

-- Players untuk Arema FC (team_id = 3)
INSERT INTO players (team_id, name, height, weight, position, jersey_number) VALUES
(3, 'Teguh Amiruddin', 182.00, 76.00, 'Penjaga Gawang'::player_position, 21),
(3, 'Arthur Irawan', 177.00, 72.00, 'Bertahan'::player_position, 3),
(3, 'Hanif Sjahbandi', 173.00, 68.00, 'Gelandang'::player_position, 8),
(3, 'Carlos Fortes', 185.00, 80.00, 'Penyerang'::player_position, 9)
ON CONFLICT DO NOTHING;

