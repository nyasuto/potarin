'use client'

import { MapContainer, TileLayer, Marker, Polyline, Popup } from 'react-leaflet'
import { LatLngExpression } from 'leaflet'
import { Route } from 'potarin-shared/types'
import 'leaflet/dist/leaflet.css'

interface MapClientProps {
  routes: Route[]
}

export default function MapClient({ routes }: MapClientProps) {
  const center: LatLngExpression = routes.length > 0
    ? [routes[0].position.lat, routes[0].position.lng]
    : [0, 0]

  const polyline: LatLngExpression[] = routes.map(r => [r.position.lat, r.position.lng])

  return (
    <MapContainer center={center} zoom={13} style={{ height: '400px', width: '100%' }}>
      <TileLayer
        attribution='&copy; <a href="http://osm.org/copyright">OpenStreetMap</a> contributors'
        url='https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png'
      />
      {routes.map((r, idx) => (
        <Marker key={idx} position={[r.position.lat, r.position.lng]}>
          <Popup>
            <strong>{r.title}</strong><br />{r.description}
          </Popup>
        </Marker>
      ))}
      {polyline.length > 1 && <Polyline positions={polyline} />}
    </MapContainer>
  )
}
