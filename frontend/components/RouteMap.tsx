"use client";

import {
  MapContainer,
  TileLayer,
  Marker,
  Polyline,
  Popup,
} from "react-leaflet";
import { Route } from "potarin-shared/types";
import { useState } from "react";
import type { LatLngExpression } from "leaflet";
import "leaflet/dist/leaflet.css";

interface RouteMapProps {
  routes: Route[];
}

export default function RouteMap({ routes }: RouteMapProps) {
  const [mapKey] = useState(() => Math.random().toString());
  if (routes.length === 0) return null;
  const positions: LatLngExpression[] = routes.map((r) => [
    r.position.lat,
    r.position.lng,
  ]);
  const center: LatLngExpression = positions[0];

  return (
    <MapContainer
      key={mapKey}
      center={center}
      zoom={13}
      className="h-80 w-full my-4"
    >
      <TileLayer
        attribution='&copy; <a href="https://osm.org/copyright">OpenStreetMap</a> contributors'
        url="https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png"
      />
      {routes.map((r, idx) => (
        <Marker position={positions[idx]} key={idx}>
          <Popup>
            <strong>{r.title}</strong>
            <div>{r.description}</div>
          </Popup>
        </Marker>
      ))}
      <Polyline positions={positions} />
    </MapContainer>
  );
}
