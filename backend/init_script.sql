CREATE TABLE IF NOT EXISTS shout_history (
  id SERIAL PRIMARY KEY,
  voice_text TEXT NOT NULL,
  response1 TEXT NOT NULL,
  response2 TEXT NOT NULL,
  response3 TEXT NOT NULL,
  response4 TEXT NOT NULL,
  response5 TEXT NOT NULL,
  response6 TEXT NOT NULL,
  created_at TIMESTAMP WITH TIME ZONE DEFAULT now()
);
