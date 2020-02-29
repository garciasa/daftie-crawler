import moment from "moment";

interface House {
  brandlink: string;
  price: string;
  date: moment.Moment;
  newdevelopment: boolean;
  meters: string;
  eircode: string;
}

function compareDate(a: any, b: any): number {
  let result = 1;
  if (a.date > b.date) {
    result = -1;
  }
  return result;
}

export function Convert(data: Array<any>): Array<House> {
  data.map(item => (item.date = moment(item.date, "DD/MM/YYYY")));

  return data;
}

export function OrderByDate(data: Array<House>): Array<House> {
  data.sort(compareDate);
  return data;
}

export function AddedLastWeek(data: Array<House>): Array<House> {
  let lastWeek: Array<House> = [];

  data.forEach(item => {
    if (item.date > moment().subtract(7, "days")) {
      lastWeek.push(item);
    }
  });

  return lastWeek;
}

export function ConvertFromStr(date: string): string {
  const cDate = moment(date).format("DD/MM/YYYY");
  return cDate === "31/12/0000" ? "--" : cDate;
}

export function GetTime(date: string): string {
  return moment(date).format("HH:mm");
}
