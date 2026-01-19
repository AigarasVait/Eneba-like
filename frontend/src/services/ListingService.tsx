import type { ListingCardDTO } from "../types/ListingCardDTO";
import { API_BASE_URL, ENDPOINTS, buildListSearchUrl } from "../constants/api";

export async function getListings(): Promise<ListingCardDTO[]> {
  const res = await fetch(`${API_BASE_URL}${ENDPOINTS.LIST}`);
  
  if (!res.ok) {
    throw new Error(`HTTP error! status: ${res.status}`);
  }
  
  return res.json();
}

export async function searchListings(query: string): Promise<ListingCardDTO[]> {
  const res = await fetch(`${API_BASE_URL}${buildListSearchUrl(query)}`);
  
  if (!res.ok) {
    throw new Error(`HTTP error! status: ${res.status}`);
  }
  
  return res.json();
}