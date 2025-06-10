import { Detail } from "potarin-shared/types";

interface Params {
  params: { id: string };
}

async function getDetail(id: string): Promise<Detail> {
  const res = await fetch(
    `${
      process.env.NEXT_PUBLIC_API_URL ?? "http://localhost:8080"
    }/api/v1/details?id=${id}`,
    { cache: "no-store" }
  );
  if (!res.ok) {
    throw new Error("Failed to fetch detail");
  }
  return res.json();
}

export default async function SuggestionDetail({ params }: Params) {
  const { id } = params;
  const detail = await getDetail(id);

  return (
    <div className="p-4">
      <h1 className="text-xl font-bold mb-4">{detail.summary}</h1>
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
