import React from "react";
import { ConvertFromStr } from "./Utils";

interface Props {
  data: Array<any>;
  lastWeek: Array<any>;
}

function Details({ data, lastWeek }: Props): React.ReactElement {
  return (
    <section className="details">
      <div className="bg-houseBlue-light  ml-10 mt-10 mr-10 pb-5 rounded-lg shadow-lg">
        <div className="text-white text-lg ml-5 pt-4 font-semibold">
          Added/Renewed in the last 7 days
        </div>
        <table className="table-auto w-full ml-3 mt-5 text-white">
          <tbody>
            {lastWeek.map((item, index) => (
              <tr key={index} className="h-10">
                <td>
                  <span className="mr-2 text-houseTurquoise-light">
                    &#9679;
                  </span>{" "}
                  <a target="_blank" rel="noopener noreferrer" href={item.url}>
                    {item.url}
                  </a>
                </td>
                <td>{item.beds === 0 ? "N/A" : `${item.beds} beds`}</td>
                <td>{ConvertFromStr(item.date_renewed)}</td>
                <td>
                  {item.price
                    .replace("Inexcessof", "")
                    .replace("AMV", "")
                    .replace("In excess of ", "")}
                </td>
              </tr>
            ))}
          </tbody>
        </table>
      </div>
      <div className="bg-houseBlue-light  ml-10 mt-10 mr-10 pb-5 rounded-lg shadow-lg">
        <div className="text-white text-lg ml-5 pt-4 font-semibold">
          All houses
        </div>
        <table className="table-auto w-full ml-3 mt-5 text-white">
          <tbody>
            {data.map((item, index) => (
              <tr key={index} className="h-10">
                <td>
                  <span className="mr-2 text-houseTurquoise-dark">&#9679;</span>{" "}
                  <a target="_blank" rel="noopener noreferrer" href={item.url}>
                    {item.url}
                  </a>
                </td>
                <td>{item.beds === 0 ? "N/A" : `${item.beds} beds`}</td>
                <td>{ConvertFromStr(item.date_renewed)}</td>
                <td>
                  {item.price
                    .replace("Inexcessof", "")
                    .replace("AMV", "")
                    .replace("In excess of ", "")}
                </td>
              </tr>
            ))}
          </tbody>
        </table>
      </div>
    </section>
  );
}

export default Details;
