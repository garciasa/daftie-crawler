import React from "react";

interface Props {
  total: number;
  lastWeek: Array<any>;
  isLoading: boolean;
}

function Figures({ total, lastWeek, isLoading }: Props): React.ReactElement {
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
        <div className="text-5xl text-white">--:--</div>
        <div className="text-sm text-white">--/--/--</div>
        <div className="text-sm font-semibold text-white">Last Parsed</div>
      </div>
    </section>
  );
}

export default Figures;
