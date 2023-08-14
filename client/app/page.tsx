"use client";
import { useEffect, useState } from "react";

interface Coords {
  lat: number;
  lon: number;
}

export default function Home() {
  const [location, setLocation] = useState<Coords>();

  async function sendLoc(location: Coords) {
    console.log({ ...location });
    try {
      await fetch("http://localhost:8000/api/users/john/location", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({ ...location }),
      });
    } catch (error) {
      console.log({ error });
    }
  }
  useEffect(() => {
    navigator.geolocation.getCurrentPosition((position) => {
      setLocation({
        lat: position.coords.latitude,
        lon: position.coords.longitude,
      });
    });
  }, []);

  return (
    <div className="flex flex-col gap-4">
      coords: {location?.lat}-lat {location?.lon}-lon
      <button onClick={async () => await sendLoc(location!!)}>
        send location
      </button>
    </div>
  );
}
