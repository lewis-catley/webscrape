import React, { useState, useEffect } from "react";
import { Spinner, Table } from "react-bootstrap";
import { useParams } from "react-router-dom";
import { getUrl } from "../../services";
import { IURL } from "../../types";

interface IParamTypes {
  id: string;
}

export const Result: React.FC = () => {
  const [isLoading, setIsLoading] = useState(false);
  const [url, setUrl] = useState<IURL>();
  const { id } = useParams<IParamTypes>();

  useEffect(() => {
    loadData();
  }, []);

  const loadData = async (): Promise<void> => {
    setIsLoading(true);
    const url = await getUrl(id);
    if (url) {
      setUrl(url);
    }
    setIsLoading(false);
  };

  if (isLoading || url === undefined) {
    return <Spinner animation="border" role="status" />;
  }

  return (
    <div>
      <h2>Results for {url.url}.</h2>
      <Table striped bordered hover>
        <thead>
          <tr>
            <th>URL</th>
            <th>Found Urls</th>
          </tr>
        </thead>
        <tbody>
          {url.results.map(
            ({ url, urlsFound }: { url: string; urlsFound: string[] }) => {
              return (
                <tr key={url}>
                  <td>{url}</td>
                  <td>
                    <ul>
                      {urlsFound.map((v: string) => {
                        return <li>{v}</li>;
                      })}
                    </ul>
                  </td>
                </tr>
              );
            }
          )}
        </tbody>
      </Table>
    </div>
  );
};
