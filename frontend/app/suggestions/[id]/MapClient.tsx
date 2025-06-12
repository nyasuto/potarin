"use client";
import dynamic from "next/dynamic";
import "leaflet/dist/leaflet.css";
import { Route } from "potarin-shared/types";

const MapContainer = dynamic(
  () => import("react-leaflet").then((m) => m.MapContainer),
  { ssr: false },
) as any;
const TileLayer = dynamic(
  () => import("react-leaflet").then((m) => m.TileLayer),
  { ssr: false },
) as any;
const Marker = dynamic(
  () => import("react-leaflet").then((m) => m.Marker),
  { ssr: false },
) as any;
const Popup = dynamic(
  () => import("react-leaflet").then((m) => m.Popup),
  { ssr: false },
) as any;

export default function MapClient({ routes }: { routes: Route[] }) {
  if (!routes || routes.length === 0) return null;
  const center: [number, number] = [
    routes[0].position.lat,
    routes[0].position.lng,
  ];

  return (
    <MapContainer
      center={center}
      zoom={13}
      style={{ height: "400px", width: "100%" }}
    >
      <TileLayer url="https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png" />
      {routes.map((r, i) => (
        <Marker
          key={i}
          position={[r.position.lat, r.position.lng] as [number, number]}
        >
          <Popup>
            <strong>{r.title}</strong>
            <br />
            {r.description}
          </Popup>
        </Marker>
      ))}
    </MapContainer>
  );
}
