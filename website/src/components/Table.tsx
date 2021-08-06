import React from "react";
import TableBS from "react-bootstrap/Table";
import { useHistory } from "react-router-dom";

export interface ITableData {
  headers: { [key: string]: string };
  data: any[];
  rowRoute: string | undefined;
}

export const Table: React.FC<ITableData> = (props: ITableData) => {
  const history = useHistory();

  const onRowClick = (r: any): void => {
    if (props.rowRoute === undefined) {
      return;
    }
    history.push(`/results/${r["id"]}`);
  };

  return (
    <TableBS striped bordered hover>
      <thead>
        <tr>
          {Object.values(props.headers).map((v: string, i: number) => (
            <th key={i}>{v}</th>
          ))}
        </tr>
      </thead>
      <tbody>
        {props.data.map((v: any, i: number) => (
          <tr key={v["id"] || i} onClick={() => onRowClick(v)}>
            {Object.entries(v).map(([key, value]: [string, any]) => {
              if (props.headers[key]) {
                return <td key={key}>{value}</td>;
              }
              return null;
            })}
          </tr>
        ))}
      </tbody>
    </TableBS>
  );
};
