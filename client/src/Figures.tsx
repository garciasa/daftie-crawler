import React  from 'react';

interface Props {
  total: number;
  lastWeek: Array<any>;
}


function Figures({ total, lastWeek }: Props): React.ReactElement {
  return (
    <section className="figures flex p-6">
      <div
        className="bg-houseBlue-light ml-10 mt-5 rounded-lg h-32 w-64 text-center shadow-lg"
      >
        <div className="text-6xl text-houseTurquoise-light">{lastWeek.length}</div>
        <div className="text-sm font-semibold text-white">Added last 7 days</div>
      </div>
      <div
        className="bg-houseBlue-light ml-10 mt-5 rounded-lg h-32 w-64 text-center shadow-lg"
      >
        <div className="text-6xl text-houseTurquoise-dark">{total}</div>
        <div className="text-sm font-semibold text-white">Total</div>
      </div>
      <div
        className="bg-houseBlue-light ml-10 mt-5 rounded-lg h-32 w-64 text-center shadow-lg"
      >
        <div className="text-5xl text-white">18:03</div>
        <div className="text-sm text-white">17/01/20</div>
        <div className="text-sm font-semibold text-white">Last Parsed</div>
      </div>
    </section>
  );
}

export default Figures; 