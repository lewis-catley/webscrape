import React, { useState, useEffect } from "react";
import Spinner from "react-bootstrap/Spinner";
import { AddEntry, Table } from "../../components";
import { postUrl, getUrls } from "../../services/urls";
import { IURL } from "../../types";

export const Home: React.FC = () => {
  const [isLoading, setIsLoading] = useState(false);
  const [urls, setUrls] = useState<IURL[]>([]);

  useEffect(() => {
    loadData();
  }, []);

  const loadData = async (): Promise<void> => {
    setIsLoading(true);
    const urls = await getUrls();
    if (urls) {
      setUrls(urls);
    }
    setIsLoading(false);
  };

  const tableHeaders = {
    url: "URL",
    isComplete: "Complete",
  };

  const addUrl = async (url: string): Promise<void> => {
    await postUrl(url);
    await loadData(); //TODO: definitely a more economical way to refresh the table
  };

  return (
    <div>
      This is a site that will attempt to find all links on a web page, please
      enter a URL and see what happens
      <AddEntry
        label="URL input"
        placeholder="e.g. https://google.com"
        onAdd={addUrl}
      />
      {isLoading ? (
        <Spinner animation="grow" />
      ) : (
        <Table headers={tableHeaders} data={urls} rowRoute={"results"} />
      )}
    </div>
  );
};
