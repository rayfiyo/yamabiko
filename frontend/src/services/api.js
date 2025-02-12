// src/services/api.js

export async function shoutVoice({ voice, demoMode }) {
  const response = await fetch("http://localhost:8080/api/shout", {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify({ voice, demoMode }),
  });

  if (!response.ok) {
    const errorData = await response.json();
    throw new Error(errorData.message || "Failed to shout.");
  }
}

export async function fetchHistory() {
  const response = await fetch("http://localhost:8080/api/history");
  if (!response.ok) {
    throw new Error("Failed to fetch history.");
  }
  return response.json(); // JSON配列が返る想定
}
