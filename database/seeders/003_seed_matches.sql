-- Seeder: Sample matches data
-- Description: Data contoh untuk pertandingan

-- Pertandingan yang sudah selesai
INSERT INTO matches (match_date, match_time, home_team_id, away_team_id, home_score, away_score, status) VALUES
('2024-10-01', '15:30:00', 1, 2, 2, 1, 'Completed'::match_status),
('2024-10-05', '19:00:00', 3, 4, 1, 1, 'Completed'::match_status),
('2024-10-08', '15:30:00', 2, 3, 3, 0, 'Completed'::match_status)
ON CONFLICT DO NOTHING;

-- Pertandingan yang akan datang
INSERT INTO matches (match_date, match_time, home_team_id, away_team_id, status) VALUES
('2024-10-20', '15:30:00', 1, 3, 'Scheduled'::match_status),
('2024-10-22', '19:00:00', 4, 5, 'Scheduled'::match_status),
('2024-10-25', '15:30:00', 2, 4, 'Scheduled'::match_status)
ON CONFLICT DO NOTHING;

