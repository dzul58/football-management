-- Seeder: Sample teams data
-- Description: Data contoh untuk tim sepak bola

INSERT INTO teams (name, logo_url, founded_year, home_address, home_city) VALUES
('Persija Jakarta', 'https://example.com/logos/persija.png', 1928, 'Jl. Gatot Subroto, Jakarta', 'Jakarta'),
('Persib Bandung', 'https://example.com/logos/persib.png', 1933, 'Jl. Sulanjana No. 17, Bandung', 'Bandung'),
('Arema FC', 'https://example.com/logos/arema.png', 1987, 'Jl. Laksamana Martadinata, Malang', 'Malang'),
('PSM Makassar', 'https://example.com/logos/psm.png', 1915, 'Jl. Racing Centre, Makassar', 'Makassar'),
('Bali United', 'https://example.com/logos/bali.png', 2015, 'Jl. Bypass Ngurah Rai, Denpasar', 'Denpasar')
ON CONFLICT (name) DO NOTHING;

