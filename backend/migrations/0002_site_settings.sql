CREATE TABLE IF NOT EXISTS site_settings (
  id BIGSERIAL PRIMARY KEY,
  site_name VARCHAR(255) NOT NULL DEFAULT '',
  site_subtitle VARCHAR(255) NOT NULL DEFAULT '',
  logo VARCHAR(500) NOT NULL DEFAULT '',
  about_text TEXT NOT NULL DEFAULT '',
  theme_color VARCHAR(32) NOT NULL DEFAULT '#f9a8d4',
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

INSERT INTO site_settings (site_name, site_subtitle, logo, about_text, theme_color)
SELECT '恋爱小屋', '记录我们的日常与心动', '', '这里记录我们的生活', '#f9a8d4'
WHERE NOT EXISTS (SELECT 1 FROM site_settings);
