import React from 'react';
import Header from './Header';
import Details from './Details';
import Figures from './Figures';
import fakeData from './mock-data';
import { Convert, OrderByDate, AddedLastWeek } from './Utils';


function App(): React.ReactElement {
  const data = OrderByDate(Convert(fakeData));
  const total = data.length;
  const lastWeek = AddedLastWeek(data);
  return (
    <div className="bg-houseBlue-dark w-full">
      <Header />
      <Figures total={total} lastWeek={lastWeek} />
      <Details data={data} lastWeek={lastWeek} />
    </div>
  );
}

export default App;
