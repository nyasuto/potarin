"use client";

import dynamic from "next/dynamic";
import { MapContainer, TileLayer, Marker, Polyline } from "react-leaflet";
import type { Route } from "potarin-shared/types";
import "leaflet/dist/leaflet.css";

function LeafletMap({ routes }: { routes: Route[] }) {
  const positions = routes.map((r) => [r.position.lat, r.position.lng] as [number, number]);
  const center = positions[0] ?? [0, 0];
  return (
    <MapContainer center={center} zoom={13} className="h-96 w-full mb-4">
      <TileLayer
        url="https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png"
        attribution="&copy; OpenStreetMap contributors"
      />
      {positions.map((pos, idx) => (
        <Marker key={idx} position={pos} />
      ))}
      {positions.length > 1 && <Polyline positions={positions} />}
    </MapContainer>
  );
}

export default dynamic(() => Promise.resolve(LeafletMap), { ssr: false });
