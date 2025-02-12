// src/services/api.js

export async function shoutVoice({ voice, demoMode }) {
  const response = await fetch("/api/shout", {
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
