export interface IURL {
  id: string;
  url: string;
  results: { url: string; urlsFound: string[] }[];
  count: number;
  isComplete: boolean;
}
