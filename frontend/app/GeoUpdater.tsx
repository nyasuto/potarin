"use client";
import { useEffect } from "react";

export default function GeoUpdater() {
  useEffect(() => {
    if (navigator.geolocation) {
      navigator.geolocation.getCurrentPosition((pos) => {
        fetch(
          `${process.env.NEXT_PUBLIC_API_URL ?? "http://localhost:8080"}/api/v1/location`,
          {
            method: "POST",
            headers: { "Content-Type": "application/json" },
            body: JSON.stringify({
              lat: pos.coords.latitude,
              lng: pos.coords.longitude,
            }),
          },
        ).catch(() => {
          // ignore errors
        });
      });
    }
  }, []);
  return null;
}
