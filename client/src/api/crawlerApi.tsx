import axios from "axios";

export interface House {
  id: string;
  url: string;
  price: string;
  beds: number;
  baths: number;
  provider: string;
  eircode: string;
  date_renewed: Date;
  firs_listed: Date;
  property_id: string;
}

export interface Stat {
  id: number;
  name: string;
  start_date: Date;
  end_date: Date;
}

export async function getAllHouses(): Promise<House[]> {
  try {
    const housesResponse = await axios.get<House[]>("/api/v1/houses");
    return housesResponse.data;
  } catch (err) {
    throw err;
  }
}

export async function getLastHouses(): Promise<House[]> {
  try {
    const housesResponse = await axios.get<House[]>("/api/v1/houses/last");
    return housesResponse.data;
  } catch (err) {
    throw err;
  }
}

export async function getStats(): Promise<Stat[]> {
  try {
    const statsResponse = await axios.get<Stat[]>("/api/v1/stats");
    return statsResponse.data;
  } catch (err) {
    throw err;
  }
}
