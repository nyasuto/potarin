"use client";

import { useEffect, useState } from "react";
import type { Suggestion } from "../../shared/types";

export default function Home() {
  const [suggestions, setSuggestions] = useState<Suggestion[]>([]);

  useEffect(() => {
    const base = process.env.NEXT_PUBLIC_API_URL ?? "http://localhost:8080";
    fetch(`${base}/api/v1/suggestions`)
      .then((res) => res.json())
      .then((data: Suggestion[]) => setSuggestions(data))
      .catch((err) => console.error(err));
  }, []);

  return (
    <div className="p-4 space-y-2">
      {suggestions.map((s) => (
        <div key={s.id} className="border p-2 rounded">
          <h2 className="font-bold">{s.title}</h2>
          <p>{s.description}</p>
        </div>
      ))}
    </div>
  );
}
