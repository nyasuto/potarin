import { Suggestion } from "potarin-shared/types";

async function getSuggestions(): Promise<Suggestion[]> {
  const res = await fetch(
    `${process.env.NEXT_PUBLIC_API_URL ?? "http://localhost:8080"}/api/v1/suggestions`,
    { cache: "no-store" },
  );
  if (!res.ok) {
    throw new Error("Failed to fetch suggestions");
  }
  return res.json();
}

export default async function Home() {
  const suggestions = await getSuggestions();

  return (
    <div className="p-4">
      <h1 className="text-xl font-bold mb-4">Course Suggestions</h1>
      <ul className="grid gap-4">
        {suggestions.map((s) => (
          <li key={s.id} className="border p-4 rounded">
            <a
              href={`/suggestions/${s.id}`}
              className="text-blue-600 hover:underline"
            >
              {s.title}
            </a>
            <p className="text-sm text-gray-600">{s.description}</p>
          </li>
        ))}
      </ul>
    </div>
  );
}
