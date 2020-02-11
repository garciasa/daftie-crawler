import React from "react";
import { GetTime, ConvertFromStr } from "./Utils";

interface Props {
  total: number;
  lastWeek: Array<any>;
  isLoading: boolean;
  stat: Array<any>;
}

function Figures({
  total,
  lastWeek,
  isLoading,
  stat
}: Props): React.ReactElement {
  return (
    <section className="figures flex p-6">
      <div className="bg-houseBlue-light ml-10 mt-5 rounded-lg h-32 w-64 text-center shadow-lg">
        <div className="text-6xl text-houseTurquoise-light">
          {lastWeek.length}
        </div>
        <div className="text-sm font-semibold text-white">
          Added last 7 days
        </div>
      </div>
      <div className="bg-houseBlue-light ml-10 mt-5 rounded-lg h-32 w-64 text-center shadow-lg">
        <div className="text-6xl text-houseTurquoise-dark">
          {isLoading ? "-" : total}
        </div>
        <div className="text-sm font-semibold text-white">Total</div>
      </div>
      <div className="bg-houseBlue-light ml-10 mt-5 rounded-lg h-32 w-64 text-center shadow-lg">
        <div className="text-5xl text-white">
          {stat.length > 0 ? GetTime(stat[1].end_date) : "--:--"}
        </div>
        <div className="text-sm text-white">
          {stat.length > 0 ? ConvertFromStr(stat[1].end_date) : "--/--/--"}
        </div>
        <div className="text-sm font-semibold text-white">Last Parsed</div>
      </div>
    </section>
  );
}

export default Figures;
