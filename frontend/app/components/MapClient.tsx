"use client";

import dynamic from "next/dynamic";
import { MapContainer, TileLayer, Marker, Polyline, Popup } from "react-leaflet";
import "leaflet/dist/leaflet.css";
import { Route } from "potarin-shared/types";

function LeafletMap({ routes }: { routes: Route[] }) {
  const center: [number, number] =
    routes.length > 0
      ? [routes[0].position.lat, routes[0].position.lng]
      : [0, 0];
  const positions = routes.map((r) => [r.position.lat, r.position.lng]) as [
    number,
    number,
  ][];

  return (
    <MapContainer center={center} zoom={13} className="h-96 w-full">
      <TileLayer url="https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png" />
      {positions.length > 1 && <Polyline positions={positions} />}
      {routes.map((r, idx) => (
        <Marker
          key={idx}
          position={[r.position.lat, r.position.lng] as [number, number]}
        >
          <Popup>
            <strong>{r.title}</strong> - {r.description}
          </Popup>
        </Marker>
      ))}
    </MapContainer>
  );
}

export default dynamic(() => Promise.resolve(LeafletMap), { ssr: false });
