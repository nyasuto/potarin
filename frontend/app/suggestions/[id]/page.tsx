import { Detail, Suggestion } from "potarin-shared/types";
import MapClient from "../../components/MapClient";

async function getSuggestions(): Promise<Suggestion[]> {
  const res = await fetch(
    `${
      process.env.NEXT_PUBLIC_API_URL ?? "http://localhost:8080"
    }/api/v1/suggestions`,
    { cache: "no-store" },
  );
  if (!res.ok) {
    throw new Error("Failed to fetch suggestions");
  }
  return res.json();
}

async function getDetail(s: Suggestion): Promise<Detail> {
  const res = await fetch(
    `${process.env.NEXT_PUBLIC_API_URL ?? "http://localhost:8080"}/api/v1/details`,
    {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify(s),
      cache: "no-store",
    },
  );
  if (!res.ok) {
    throw new Error("Failed to fetch detail");
  }
  return res.json();
}

export default async function SuggestionDetail({
  params,
}: {
  params: Promise<{ id: string }>;
}) {
  const { id } = await params;
  const suggestions = await getSuggestions();
  const suggestion = suggestions.find((s) => s.id === id);
  if (!suggestion) {
    throw new Error("Suggestion not found");
  }
  const detail = await getDetail(suggestion);

  return (
    <div className="p-4">
      <h1 className="text-xl font-bold mb-4">{detail.summary}</h1>
      <div className="mb-4">
        <MapClient routes={detail.routes} />
      </div>
      <ul className="space-y-2">
        {detail.routes.map((r, index) => (
          <li key={index} className="border p-2 rounded">
            <strong>{r.title}</strong> - {r.description} ({r.position.lat},{" "}
            {r.position.lng})
          </li>
        ))}
      </ul>
    </div>
  );
}
