import axios, { AxiosResponse } from "axios";
import { IURL } from "../types";

const urlsEndpoint = process.env.REACT_APP_URLS_API || "";

export const getUrls = async (): Promise<IURL[]> => {
  return axios
    .get<IURL[]>(`${urlsEndpoint}/urls`)
    .then((r: AxiosResponse<IURL[]>) => {
      return r.data;
    });
};

export const getUrl = async (id: string): Promise<IURL> => {
  return axios
    .get<IURL>(`${urlsEndpoint}/urls/${id}`)
    .then((r: AxiosResponse<IURL>) => {
      return r.data;
    });
};

export const postUrl = async (url: string): Promise<string> => {
  return axios
    .post<string>(`${urlsEndpoint}/urls`, { url })
    .then((r: AxiosResponse<string>) => {
      return r.data;
    });
};
