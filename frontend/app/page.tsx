import { Suggestion } from "potarin-shared/types";
import Link from "next/link";
import { ensureUniqueIds } from "../lib/ensureUniqueIds";

async function getSuggestions(): Promise<Suggestion[]> {
  const res = await fetch(
    `${process.env.NEXT_PUBLIC_API_URL ?? "http://localhost:8080"}/api/v1/suggestions`,
    { cache: "no-store" },
  );
  if (!res.ok) {
    throw new Error("Failed to fetch suggestions");
  }
  const data = await res.json();
  return Array.isArray(data) ? data : [];
}

export default async function Home() {
  const suggestions = ensureUniqueIds(await getSuggestions());

  return (
    <div className="max-w-4xl mx-auto p-4">
      <div className="bg-gradient-to-r from-sky-400 to-emerald-300 text-white rounded-lg p-10 mb-8 shadow">
        <h1 className="text-3xl font-bold mb-2">Plan Your Next Adventure</h1>
        <p className="text-lg">AI-powered suggestions for walking & cycling</p>
      </div>
      <h2 className="text-2xl font-semibold mb-4">Course Suggestions</h2>
      <ul className="grid gap-6 md:grid-cols-2">
        {suggestions.map((s) => (
          <li key={s.id} className="bg-white border p-4 rounded shadow">
            <Link
              href={`/suggestions/${s.id}`}
              className="text-blue-600 hover:underline font-medium"
            >
              {s.title}
            </Link>
            <p className="text-sm text-gray-600 mt-2">{s.description}</p>
          </li>
        ))}
      </ul>
    </div>
  );
}
