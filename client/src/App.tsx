import React, { useState, useEffect } from "react";
import axios from "axios";
import Header from "./Header";
import Details from "./Details";
import Figures from "./Figures";

export default function App() {
  const [isLoading, setIsLoading] = useState(true);
  const [houses, setHouses] = useState<any>([]);
  const [total, setTotal] = useState(0);
  const [lastWeek, setLastWeek] = useState<any>([]);
  const [stat, setStat] = useState<any>([]);

  useEffect(() => {
    axios.get("/api/v1/houses").then(resp => {
      console.log(resp);
      setHouses(resp.data);
      setTotal(resp.data.length);
      setIsLoading(false);
    });
    axios.get("/api/v1/houses/last").then(resp => {
      setLastWeek(resp.data);
    });

    axios.get("/api/v1/stats").then(resp => {
      setStat(resp.data);
    });
  }, []);

  return (
    <div className="bg-houseBlue-dark w-full">
      <Header />
      <Figures
        isLoading={isLoading}
        total={total}
        lastWeek={lastWeek}
        stat={stat}
      />
      <Details data={houses} lastWeek={lastWeek} />
    </div>
  );
}
