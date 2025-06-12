"use client";
import { MapContainer, TileLayer, Marker, Popup } from "react-leaflet";
import type { Detail } from "potarin-shared/types";
import "leaflet/dist/leaflet.css";
import L from "leaflet";

const icon = new L.Icon({
  iconUrl: "https://unpkg.com/leaflet@1.9.4/dist/images/marker-icon.png",
  shadowUrl: "https://unpkg.com/leaflet@1.9.4/dist/images/marker-shadow.png",
  iconSize: [25, 41],
  iconAnchor: [12, 41],
});

export default function MapClient({ detail }: { detail: Detail }) {
  const first = detail.routes[0]?.position ?? { lat: 0, lng: 0 };
  return (
    <MapContainer
      center={[first.lat, first.lng]}
      zoom={13}
      style={{ height: "400px", width: "100%" }}
    >
      <TileLayer url="https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png" />
      {detail.routes.map((r, i) => (
        <Marker
          key={i}
          position={[r.position.lat, r.position.lng]}
          icon={icon}
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
