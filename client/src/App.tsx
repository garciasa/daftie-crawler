import React, { useState, useEffect } from "react";
import axios from "axios";
import Header from "./Header";
import Details from "./Details";
import Figures from "./Figures";
import { Convert, OrderByDate, AddedLastWeek } from "./Utils";

export default function App() {
  const [isLoading, setIsLoading] = useState(true);
  const [houses, setHouses] = useState<any>([]);
  const [total, setTotal] = useState(0);
  const [lastWeek, setLastWeek] = useState<any>([]);

  useEffect(() => {
    axios.get("http://127.0.0.1:8080/api/v1/houses").then(resp => {
      setHouses(OrderByDate(Convert(resp.data)));
      setTotal(resp.data.length);
      setLastWeek(AddedLastWeek(resp.data));
      setIsLoading(false);
    });
  }, []);

  //const data = OrderByDate(Convert(fakeData));
  //const total = data.length;
  //const lastWeek = AddedLastWeek(data);
  return (
    <div className="bg-houseBlue-dark w-full">
      <Header />
      <Figures isLoading={isLoading} total={total} lastWeek={lastWeek} />
      <Details data={houses} lastWeek={lastWeek} />
    </div>
  );
}
