import React from 'react';

interface Props {
  data: Array<any>;
  lastWeek: Array<any>;
}

function Details({data, lastWeek}: Props): React.ReactElement {
  return (
    <section className="details">
        <div className="bg-houseBlue-light ml-10 mr-10 pb-5 rounded-lg shadow-lg">
          <div className="text-white text-lg ml-5 pt-4 font-semibold">
            Added in the last 7 days
          </div>
          <table className="table-auto w-full ml-10 mt-5 text-white">
          <tbody>
              {lastWeek.map((item, index) => 
                <tr key={index} className="h-10">
                  <td className="pr-5">
                    <span className="mr-2 text-houseTurquoise-dark">&#9679;</span> <a target="_blank" rel="noopener noreferrer" href={`https://daft.ie${item.brandlink}`}>{item.brandlink}</a>
                  </td>
                  <td>
                    {item.meters}
                  </td>
                  <td>
                    {item.date.format("DD/MM/YYYY")}
                  </td>
                  <td>
                    {item.price}
                  </td>
                </tr>
              )}
            </tbody>
          </table>
        </div>
        <div className="bg-houseBlue-light  ml-10 mt-10 mr-10 pb-5 rounded-lg shadow-lg">
          <div className="text-white text-lg ml-5 pt-4 font-semibold">
            All houses
          </div>
          <table className="table-auto w-full ml-10 mt-5 text-white">
            <tbody>
              {data.map((item, index) => 
                <tr key={index} className="h-10">
                  <td>
                    <span className="mr-2 text-houseTurquoise-dark">&#9679;</span> <a target="_blank" rel="noopener noreferrer" href={`https://daft.ie${item.brandlink}`}>{item.brandlink}</a>
                  </td>
                  <td>
                    {item.meters}
                  </td>
                  <td>
                    {item.date.format("DD/MM/YYYY")}
                  </td>
                  <td>
                    {item.price}
                  </td>
                </tr>
              )}
            </tbody>
          </table>
        </div>
      </section>
  );
}

export default Details; 