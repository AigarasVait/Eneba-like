export const API_BASE_URL = "http://localhost:8080";

export const ENDPOINTS = {
  LIST: "/list",
};

export const buildListSearchUrl = (query: string) =>
  `${ENDPOINTS.LIST}?search=${encodeURIComponent(query)}`;
