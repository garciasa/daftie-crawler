import React, { useState, useEffect } from "react";
import Header from "../components/Header";
import Details from "../components/Details";
import Figures from "../components/Figures";
import {
  getAllHouses,
  getLastHouses,
  getStats,
  Statistics,
} from "../api/crawlerApi";

export default function App() {
  const [isLoading, setIsLoading] = useState(true);
  const [houses, setHouses] = useState<any>([]);
  const [total, setTotal] = useState(0);
  const [lastWeek, setLastWeek] = useState<any>([]);
  const [stat, setStat] = useState<any>([]);

  useEffect(() => {
    async function fetchData() {
      try {
        const houses = await getAllHouses();
        setHouses(houses);
        const lastHouses = await getLastHouses();
        setLastWeek(lastHouses);
        const stats = await getStats();
        setStat(stats.crawler);
        setTotal(stats.general.total);
        setIsLoading(false);
      } catch (err) {
        //TODO: put error in some place
        console.log(err);
      }
    }
    setIsLoading(true);
    fetchData();
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
